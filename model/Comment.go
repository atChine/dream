package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId    uint   `gorm:"type:int" json:"user_id"`
	ArticleId uint   `gorm:"type:int" json:"article_id"`
	Content   string `gorm:"type:varchar(500);not null" json:"content"`
	Status    int8   `gorm:"type:tinyint" json:"status"`
	UserName  string `gorm:"type:longtext" json:"user_name"`
	Title     string `gorm:"type:longtext" json:"title"`
}
