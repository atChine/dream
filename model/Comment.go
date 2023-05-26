package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CommentId      uint   `gorm:"type:int" json:"comment_id"`
	UserId         uint   `gorm:"type:int" json:"user_id"`
	ArticleId      uint   `gorm:"type:int" json:"article_id"`
	CommentContent string `gorm:"type:varchar(500);not null" json:"comment_content"`
	CommentStatus  int8   `gorm:"type:tinyint" json:"comment_status"`
}
