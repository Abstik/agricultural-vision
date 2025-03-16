package logic

import (
	"agricultural_vision/dao/mysql"
	"strconv"

	"go.uber.org/zap"

	"agricultural_vision/dao/redis"
	"agricultural_vision/models/request"
)

// 给帖子投票
func VoteForPost(userID int64, p *request.VoteRequest) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.Int64("postID", p.PostID),
		zap.Int8("direction", p.Direction),
	)
	return redis.VoteForPost(strconv.Itoa(int(userID)), strconv.Itoa(int(p.PostID)), float64(p.Direction))
}

// 给评论投票
func CommentVote(userID int64, p *request.VoteRequest) error {
	zap.L().Debug("CommentVote",
		zap.Int64("userID", userID),
		zap.Int64("commentID", p.CommentID),
		zap.Int8("direction", p.Direction),
	)

	// 查询父评论id和帖子id
	parentID, postID, err := mysql.GetParentIDAndPostIDByCommentID(p.CommentID)
	if err != nil {
		return err
	}

	if parentID != nil { // 给一级评论投票
		return redis.VoteForComment(strconv.Itoa(int(userID)), strconv.Itoa(int(p.CommentID)), strconv.Itoa(int(*postID)), float64(p.Direction), true)
	} else { // 给二级评论投票
		return redis.VoteForComment(strconv.Itoa(int(userID)), strconv.Itoa(int(p.CommentID)), strconv.Itoa(int(*postID)), float64(p.Direction), false)
	}
}
