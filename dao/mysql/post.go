package mysql

import (
	"errors"
	"fmt"
	"strings"

	"agricultural_vision/models"
)

// 创建帖子
func CreatePost(p *models.Post) error {
	result := DB.Create(p)
	// 在执行 SQL 语句或与数据库交互过程中是否发生了错误
	if result.Error != nil {
		return result.Error
	}
	// 虽然没有发生错误，但插入操作没有成功插入任何数据
	if result.RowsAffected == 0 {
		return errors.New("创建帖子失败")
	}
	return nil
}

// 根据帖子id查询帖子详情
func GetPostById(pid int64) (*models.Post, error) {
	var post *models.Post
	result := DB.Where("post_id = ?", pid).First(&post)
	return post, result.Error
}

// 查询帖子列表
func GetPostList(pageNum, pageSize int64) ([]*models.Post, error) {
	var posts []*models.Post

	result := DB.
		Order("created_time DESC").
		Offset(int((pageNum - 1) * pageSize)).
		Limit(int(pageSize)).
		Find(&posts)

	return posts, result.Error
}

// 根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) ([]*models.Post, error) {
	var posts []*models.Post

	if len(ids) == 0 {
		return posts, nil
	}

	//order by FIND_IN_SET(post_id, ?) 表示根据 post_id 在另一个给定字符串列表中的位置进行排序。
	//? 是另一个占位符，将被替换为一个包含多个ID的字符串，例如 "1,3,2"。
	result := DB.
		Where("post_id IN ?", ids).
		Order(fmt.Sprintf("FIELD(post_id, %s)", strings.Join(ids, ","))).
		Find(&posts)

	return posts, result.Error
}
