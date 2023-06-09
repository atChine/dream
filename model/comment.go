package model

import (
	"dream/utils/errmsg"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	ArticleId uint   `json:"article_id"`
	Title     string `json:"article_title"`
	Username  string `json:"username"`
	Content   string `gorm:"type:varchar(500);not null;" json:"content"`
	Status    int8   `gorm:"type:tinyint;default:2" json:"status"`
}

func GetCommentListFront(id, pageSize, pageNum int) ([]Comment, int64, int) {
	var total int64
	var commentList []Comment
	db.Find(&Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	err = db.Model(&Comment{}).
		Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").
		Select("comment.id, article.title, user_id, article_id, user.username, comment.content, comment.status, comment.created_at, comment.deleted_at").
		Joins("LEFT JOIN article ON comment.article_id = article.id").
		Joins("LEFT JOIN user ON comment.user_id = user.id").Where("article_id = ?", id).
		Where("status = ?", 1).Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCSE
}
