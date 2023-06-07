package model

import (
	"dream/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

// GetArt 查询文章列表
func GetArt(pageSize, pageNum int) ([]Article, int, int64) {
	var artList []Article
	var err error
	var total int64
	err = db.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("Created_At DESC").
		Joins("Category").
		Find(&artList).Error
	db.Model(&artList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return artList, errmsg.SUCCSE, total
}

// GetArtByTitle 根据title查询文章列表
func GetArtByTitle(title string, pageSize, pageNum int) ([]Article, int, int64) {
	var artList []Article
	var err error
	var total int64
	err = db.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").
		Where("title LIKE ?", "%"+title+"%").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Order("Created_At DESC").
		Joins("Category").
		Find(&artList).Error
	db.Model(&artList).Where("title LIKE ?", "%"+title+"%").Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return artList, errmsg.SUCCSE, total
}
