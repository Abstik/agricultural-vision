package redis

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"

	"agricultural_vision/constants"
	"agricultural_vision/models/request"
)

// 新建帖子
func CreatePost(postID int64, communityID int64) error {
	//开启事务
	pipeline := client.TxPipeline()

	//在redis中更新帖子创建时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//在redis中更新帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()), // 默认分数不是0，而是当前的时间戳（这样总分数可以结合投票数和时间）
		Member: postID,
	})

	//在redis中更新帖子和社区关系
	pipeline.SAdd(getRedisKey(KeyCommunitySetPF+strconv.Itoa(int(communityID))), postID)

	//初始化帖子评论数
	pipeline.ZAdd(getRedisKey(KeyPostCommentNumZSet), redis.Z{
		Score:  0,
		Member: postID,
	})

	_, err := pipeline.Exec()
	return err
}

// 删除帖子
func DeletePost(postID int64, communityID int64) error {
	pipeline := client.TxPipeline()

	// 先查询出此帖子所有评论的 id
	commentIDs, _ := client.ZRange(getRedisKey(KeyCommentTimeZSetPF+strconv.FormatInt(postID, 10)), 0, -1).Result()

	// 1. 从时间排序集合删除
	pipeline.ZRem(getRedisKey(KeyPostTimeZSet), postID) // 从 zset 中移除指定成员

	// 2. 从热度排序集合删除
	pipeline.ZRem(getRedisKey(KeyPostScoreZSet), postID)

	// 3. 删除帖子点赞记录
	pipeline.Del(getRedisKey(KeyPostVotedZSetPF + strconv.FormatInt(postID, 10))) // 删除整个 key

	// 4. 从社区帖子集合删除
	pipeline.SRem(getRedisKey(KeyCommunitySetPF+strconv.FormatInt(communityID, 10)), postID) // 从 set 中移除指定成员

	// 5. 删除帖子评论数记录
	pipeline.ZRem(getRedisKey(KeyPostCommentNumZSet), postID)

	// 6. 删除该帖子下的所有评论时间记录
	pipeline.Del(getRedisKey(KeyCommentTimeZSetPF + strconv.FormatInt(postID, 10)))

	// 7. 删除该帖子下的所有评论投票记录
	pipeline.Del(getRedisKey(KeyCommentScoreZSetPF + strconv.FormatInt(postID, 10)))

	// 8. 处理该帖子下所有评论的删除
	if len(commentIDs) > 0 {
		// 8.1 删除该帖子下所有评论的点赞记录，拼凑出需要删除的key的列表，一次性删除多个 key，提高性能
		keysToDelete := make([]string, len(commentIDs))
		for i, commentID := range commentIDs {
			keysToDelete[i] = getRedisKey(KeyCommentVotedZSetPF + commentID)
		}
		pipeline.Del(keysToDelete...) // 一次删除多个 key，提高性能

		// 8.2 删除所有评论的子评论数记录，根据commentId列表删除
		interfaceIDs := make([]interface{}, len(commentIDs))
		for i, id := range commentIDs {
			interfaceIDs[i] = id
		}
		pipeline.ZRem(getRedisKey(KeyCommentNumZSet), interfaceIDs...)
	}

	// 执行事务
	_, err := pipeline.Exec()
	return err
}

// 根据键名和索引，查询id列表
func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1

	// ZRevRange 按分数从大到小查询指定数量的元素
	result, err := client.ZRevRange(key, start, end).Result()
	return result, err
}

// 根据排序方式和索引，查询id列表
func GetPostIDsInOrder(p *request.ListRequest) ([]string, error) {
	//从redis中获取id
	//1.根据用户请求中携带的order参数（排序方式）确定要查询的redis key
	key := getRedisKey(KeyCommentTimeZSetPF)
	if p.Order == constants.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	//2.确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	//3.ZREVRANGE 按分数从大到小查询指定数量的元素
	result, err := client.ZRevRange(key, start, end).Result()

	return result, err
}

// 根据ids列表查询每篇帖子的投赞成票的数据
func GetPostVoteDataByIDs(ids []string) (data []int64, err error) {
	// 使用 pipeline 批量执行 Redis 命令
	pipeline := client.Pipeline()
	var cmds []redis.Cmder

	// 将所有 ZCount 命令添加到 pipeline 中
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		// 统计该帖子的赞成票数
		cmds = append(cmds, pipeline.ZCount(key, "1", "1"))
	}

	// 执行所有命令
	_, err = pipeline.Exec()
	if err != nil {
		return nil, err
	}

	// 处理结果
	for _, cmd := range cmds {
		// 类型断言为 redis.IntCmd
		voteCmd := cmd.(*redis.IntCmd)
		data = append(data, voteCmd.Val()) // 获取票数并保存
	}

	return data, nil
}

// 根据社区id查询该社区下的id列表
func GetCommunityPostIDsInOrder(p *request.ListRequest, communityID int64) (ids []string, err error) {
	//根据指定的排序方式，确定要操作的redis中的key
	//orderKey指定排序方式的键名，按时间排序则是KeyPostTimeZSet，按分数排序则是KeyPostScoreZSet
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == constants.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	//从KeyCommunitySetPF中查询该社区下的帖子id列表，根据id列表去KeyPostTimeZSet或KeyPostScoreZSet中去查询时间或分数
	//也就是查询交集，将查询到的内容（帖子postID和对应的时间或分数）保存到新的自定义的key中

	//社区的key
	communityKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))

	//自定义新key，用来存储两表交集的，值为postID和对应的时间或分数，表示此社区分类下的帖子和时间/分数
	key := orderKey + strconv.Itoa(int(communityID))

	// 使用 pipeline 批量执行 Redis 命令
	pipeline := client.Pipeline()

	//通过 ZInterStore 对有序集合communityKey和orderKey进行交集运算，结果存储到key中
	pipeline.ZInterStore(key, redis.ZStore{
		Aggregate: "MAX", //表示交集的分数取较大的值，如果将一个普通的set（无序集合）与一个zset（有序集合）一起参与ZInterStore操作，Redis会自动将set视为一个所有成员分数为1的特殊zset
	}, communityKey, orderKey)

	pipeline.Expire(key, 60*time.Second) // 设置超时时间

	_, err = pipeline.Exec()
	if err != nil {
		return
	}

	//查询指定索引范围的id列表
	return getIDsFormKey(key, p.Page, p.Size)
}
