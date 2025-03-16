package mysql

import (
	"fmt"
	"strings"

	"agricultural_vision/constants"
	"agricultural_vision/models/entity"
	"agricultural_vision/models/request"
)

// 创建评论
func CreateComment(comment *entity.Comment) error {
	result := DB.Create(comment)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return constants.ErrorNotAffectData
	}

	return nil
}

// 删除评论
func DeleteComment(commentID int64) error {
	result := DB.Delete(&entity.Comment{}, commentID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return constants.ErrorNotAffectData
	}

	return nil
}

// 根据评论ID获取父评论ID和帖子ID
func GetParentIDAndPostIDByCommentID(commentID int64) (*int64, *int64, error) {
	// 定义一个结构体来接收查询结果
	var result struct {
		ParentID *int64 `json:"parent_id"`
		PostID   *int64 `json:"post_id"`
	}

	// 查询父评论ID和帖子ID
	err := DB.Model(&entity.Comment{}).Select("parent_id", "post_id").
		Where("id = ?", commentID).
		Scan(&result).Error
	if err != nil {
		return nil, nil, err
	}

	return result.ParentID, result.PostID, nil
}

// 根据评论ID列表获取评论列表（一级评论适用）
func GetCommentListByIDs(ids []string) ([]*entity.Comment, error) {
	var comments []*entity.Comment

	if len(ids) == 0 {
		return comments, nil
	}

	//order by FIND_IN_SET(post_id, ?) 表示根据 post_id 在另一个给定字符串列表中的位置进行排序。
	//? 是另一个占位符，将被替换为一个包含多个ID的字符串，例如 "1,3,2"。
	result := DB.
		Where("id IN ?", ids).
		Order(fmt.Sprintf("FIELD(id, %s)", strings.Join(ids, ","))).
		Find(&comments)

	return comments, result.Error
}

// 分页查询二级评论
func GetSecondLevelCommentList(commentID int64, request *request.ListRequest) ([]*entity.Comment, error) {
	var comments []*entity.Comment

	// 计算偏移量
	offset := (request.Page - 1) * request.Size

	// 查询二级评论（parent_id 为 commentID）
	result := DB.
		Where("parent_id = ?", commentID).
		Order("created_at DESC"). // 默认按时间倒序排序
		Limit(int(request.Size)).
		Offset(int(offset)).
		Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}
