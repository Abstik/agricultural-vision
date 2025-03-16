package mysql

import (
	"fmt"
	"strings"

	"agricultural_vision/constants"
	"agricultural_vision/models/entity"
)

// 创建帖子
func CreatePost(p *entity.Post) error {
	result := DB.Create(p)
	// 在执行 SQL 语句或与数据库交互过程中是否发生了错误
	if result.Error != nil {
		return result.Error
	}
	// 虽然没有发生错误，但插入操作没有成功插入任何数据
	if result.RowsAffected == 0 {
		return constants.ErrorNotAffectData
	}
	return nil
}

// 删除帖子
func DeletePost(id int64) error {
	result := DB.Delete(&entity.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return constants.ErrorNotAffectData
	}
	return nil
}

// 根据帖子id查询帖子详情
func GetPostById(pid int64) (*entity.Post, error) {
	var post *entity.Post
	result := DB.Where("post_id = ?", pid).First(&post)
	return post, result.Error
}

// 根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) ([]*entity.Post, error) {
	var posts []*entity.Post

	if len(ids) == 0 {
		return posts, nil
	}

	//order by FIND_IN_SET(post_id, ?) 表示根据 post_id 在另一个给定字符串列表中的位置进行排序。
	//? 是另一个占位符，将被替换为一个包含多个ID的字符串，例如 "1,3,2"。
	result := DB.
		Where("id IN ?", ids).
		Order(fmt.Sprintf("FIELD(id, %s)", strings.Join(ids, ","))).
		Find(&posts)

	return posts, result.Error
}
