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
	postIDStr := strconv.FormatInt(postID, 10)

	//开启事务
	pipeline := client.TxPipeline()

	//在redis中更新帖子创建时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postIDStr,
	})

	//在redis中更新帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()), // 默认分数不是0，而是当前的时间戳（这样总分数可以结合投票数和时间）
		Member: postIDStr,
	})

	//在redis中更新帖子和社区关系
	pipeline.SAdd(getRedisKey(KeyCommunitySetPF+strconv.Itoa(int(communityID))), postID)

	// 更新帖子评论数
	pipeline.ZIncrBy(getRedisKey(KeyPostCommentNumZSet), 0, postIDStr)

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
func GetPostIDsInOrder(p *request.PostListRequest) ([]string, error) {
	//从redis中获取id
	//1.根据用户请求中携带的order参数（排序方式）确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
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

// 根据id查询此帖子的赞成票数
func GetPostVoteDataByID(id string) int64 {
	key := getRedisKey(KeyPostVotedZSetPF + id)
	data := client.ZCount(key, "1", "1").Val()
	return data
}

// 根据社区id查询该社区下的id列表
func GetCommunityPostIDsInOrder(p *request.CommunityPostListRequest) (ids []string, err error) {
	//根据指定的排序方式，确定要操作的redis中的key
	//orderKey指定排序方式的键名，按时间排序则是KeyPostTimeZSet，按分数排序则是KeyPostScoreZSet
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == constants.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	//从KeyCommunitySetPF中查询该社区下的帖子id列表，根据id列表去KeyPostTimeZSet或KeyPostScoreZSet中去查询时间或分数
	//也就是查询交集，将查询到的内容（帖子postID和对应的时间或分数）保存到新的自定义的key中

	//社区的key
	communityKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))

	//自定义新key，用来存储两表交集的，值为postID和对应的时间或分数，表示此社区分类下的帖子和时间/分数
	key := orderKey + strconv.Itoa(int(p.CommunityID))

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
