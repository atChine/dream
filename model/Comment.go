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

// GetCommentById 获取单个评论信息

func GetCommentById(comId string) (Comment, int) {
	var comment Comment
	err := db.Where("comment_id = ?", comId).First(&comment).Error
	if err != nil {
		return comment, errmsg.ERROR
	}
	return comment, errmsg.SUCCSE
}

// GetArtCommentList 展示页面获取评论列表
func GetArtCommentList(artId string, pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	db.Find(&Comment{}).Where("article_id = ?", artId).Where("comment_status = ?", 1).Count(&total)
	err = db.Model(&Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).
		Order("Created_At DESC").
		Select("comment.id, comment_id, article.article_title, user_id, article_id, user.user_name, "+
			"comment_content, comment_status,comment.created_at,comment.deleted_at").
		Joins("LEFT JOIN article ON comment.article_id = article.id").
		Joins("LEFT JOIN user ON comment.user_id = user.id").Where("article_id = ?", artId).
		Where("comment_status = ?", 1).Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCSE
}

// DeleteCommentById 删除评论
func DeleteCommentById(comId string) int {
	var comment Comment
	err = db.Where("comment_id = ?", comId).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetCommentList 后台获取所有评论列表
func GetCommentList(pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	db.Find(&Comment{}).Count(&total)
	err = db.Model(&commentList).Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("Created_At DESC").
		Select("comment.id, comment_id, article.article_title, user_id, article_id, user.user_name, " +
			"comment_content, comment_status,comment.created_at,comment.deleted_at").
		Joins("LEFT JOIN article ON comment.article_id = article.id").
		Joins("LEFT JOIN user ON comment.user_id = user.id").Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCSE
}
