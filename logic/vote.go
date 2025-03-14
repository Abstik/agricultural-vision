package logic

import (
	"strconv"

	"go.uber.org/zap"

	"agricultural_vision/dao/redis"
	"agricultural_vision/models/request"
)

// 给帖子投票
func VoteForPost(userID int64, p *request.VotePostRequest) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.Int64("postID", p.PostID),
		zap.Int8("direction", p.Direction),
	)
	return redis.VoteForPost(strconv.Itoa(int(userID)), strconv.Itoa(int(p.PostID)), float64(p.Direction))
}
