package model

import (
	"gorm.io/gorm"
	"love_blog/utils/errmsg"
)

type Comment struct {
	gorm.Model
	CommentId      uint   `gorm:"type:varchar(36);not null" json:"comment_id"`
	UserId         uint   `gorm:"type:varchar(36);not null" json:"user_id"`
	ArticleId      uint   `gorm:"type:varchar(36);not null" json:"article_id"`
	CommentContent string `gorm:"type:varchar(500);not null" json:"comment_content"`
	CommentStatus  int8   `gorm:"type:tinyint" json:"comment_status"`
}

// AddComment 新增评论
func AddComment(data *Comment) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetCommentCount 获取文章评论数量
func GetCommentCount(artId string) int64 {
	var comment Comment
	var count int64
	db.Find(&comment).Where("article_id = ?").Where("comment_status = ?", 1).Count(&count)
	return count
}
