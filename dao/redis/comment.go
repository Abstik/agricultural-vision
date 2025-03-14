package redis

import (
	"strconv"

	"github.com/go-redis/redis"

	"agricultural_vision/constants"
	"agricultural_vision/models/request"
)

// 根据排序方式和索引范围，查询评论id列表
func GetCommentIDsInOrder(p *request.CommentListRequest) ([]string, error) {
	//从redis中获取id
	//1.根据用户请求中携带的order参数（排序方式）确定要查询的redis key
	key := getRedisKey(KeyPostCommentTimeZSetPF + strconv.Itoa(int(p.PostID)))
	if p.Order == constants.OrderScore {
		key = getRedisKey(KeyPostCommentScoreZSetPF)
	}

	//2.确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	//3.ZREVRANGE 按分数从大到小查询指定数量的元素
	result, err := client.ZRevRange(key, start, end).Result()

	return result, err
}

// 根据ids列表查询每条评论的赞成票数据
func GetCommentVoteDataByIDs(ids []string) (data []int64, err error) {
	// 使用 pipeline 批量执行 Redis 命令
	pipeline := client.Pipeline()
	var cmds []redis.Cmder

	// 将所有 ZCount 命令添加到 pipeline 中
	for _, id := range ids {
		key := getRedisKey(KeyPostCommentScoreZSetPF + id)
		// 统计该帖子的赞成票数（1代表赞成票）
		cmds = append(cmds, pipeline.ZCount(key, "1", "1"))
	}

	// 执行所有命令
	_, err = pipeline.Exec()
	if err != nil {
		return nil, err
	}

	// 处理返回的结果
	for _, cmd := range cmds {
		// 类型断言为 redis.IntCmd
		voteCmd := cmd.(*redis.IntCmd)
		data = append(data, voteCmd.Val()) // 获取赞成票数
	}

	return data, nil
}

// 根据id列表查询帖子的评论数
func GetCommentNumByID(ids []string) (nums []float64, err error) {
	// 使用 pipeline 批量执行 Redis 命令
	pipeline := client.Pipeline()
	var cmds []redis.Cmder

	// 将所有 ZScore 命令添加到 pipeline 中
	for _, id := range ids {
		key := getRedisKey(KeyPostCommentNumZSet)
		cmds = append(cmds, pipeline.ZScore(key, id)) // 获取每个评论的分数（即评论数）
	}

	// 执行所有命令
	_, err = pipeline.Exec()
	if err != nil {
		return nil, err
	}

	// 处理结果
	for _, cmd := range cmds {
		scoreCmd := cmd.(*redis.FloatCmd)   // 类型断言为 FloatCmd
		nums = append(nums, scoreCmd.Val()) // 获取评论数
	}

	return nums, nil
}
