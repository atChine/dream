package model

import (
	"gorm.io/gorm"
	"love_blog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignkey:category_id"`
	gorm.Model
	ArticleId           string `gorm:"type:varchar(36);not null" json:"article_id"`
	UserId              uint   `gorm:"type:varchar(36);not null" json:"user_id"`
	ArticleTitle        string `gorm:"type:varchar(100);not null" json:"article_title"`
	ArticleCateId       int    `gorm:"type:int" json:"category_id"`
	ArticleDesc         string `gorm:"type:varchar(200)" json:"article_desc"`
	ArticleContent      string `gorm:"type:longtext" json:"article_content"`
	ArticleCommentCount int    `gorm:"type:int;not null;default:0" json:"article_comment_count"`
	ArticleReadCount    int    `gorm:"type:int;not null;default:0" json:"article_read_count"`
	ArticleLikeCount    int    `gorm:"type:int;not null;default:0" json:"article_like_count"`
}

// AddArt CreateAdd 新增文章
func AddArt(articleAdd *Article) int {
	err := db.Create(&articleAdd)
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //200
}

// DelArtById 通过id删除指定文章
func DelArtById(artId string) int {
	var article Article
	err := db.Where("article_id = ?", artId).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// EdiArtById 根据id修改文章
func EdiArtById(artId string, data *Article) int {
	var art Article
	err := db.Model(&art).Where("article_id = ?", artId).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetArtInfoById 根据id查询文章详情
func GetArtInfoById(artId string) (Article, int) {
	var art Article
	err := db.Where("article_id = ?", artId).Preload("Category").First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	// 更新浏览次数
	db.Model(&art).Where("article_id = ?", artId).UpdateColumn("article_read_count", gorm.Expr("article_read_count + ?", 1))
	return art, errmsg.SUCCSE
}

// GetArtInfoByCate 分页查询分类下的所有文章
func GetArtInfoByCate(cateId string, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("article_cate_id = ?", cateId).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	db.Model(&cateArtList).Where("article_cate_id = ?", cateId).Count(&total)
	return cateArtList, errmsg.SUCCSE, total
}

// GetArtInfo 分页查询所有文章
func GetArtInfo(pageSize int, pageNum int) ([]Article, int, int64) {
	var atrList []Article
	var total int64
	err := db.Select("article.id, article_id, article_title , create_at, updated_at, article_desc, article_comment_count, article_read_count, article_like_count, category.category_name").
		Limit(pageSize).Offset((pageNum - 1) * pageNum).Order("Created_At DESC").Joins("Category").Find(&atrList)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	db.Model(&atrList).Count(&total)
	return atrList, errmsg.SUCCSE, total
}

// GetArtInfoByTitle 按照标题搜索文章
func GetArtInfoByTitle(pageSize int, pageNum int, artTitle string) ([]Article, int, int64) {
	var atrList []Article
	var total int64
	err := db.Select("article.id, article.article_id, article_title , create_at, updated_at, article_desc, article_comment_count, article_read_count, article_like_count, category.category_name").
		Where("title LIKE ?", "%"+artTitle+"%").Limit(pageSize).Offset((pageNum - 1) * pageNum).Order("Created_At DESC").Joins("Category").Find(&atrList).Error
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	db.Model(&atrList).Where("title LIKE ?", "%"+artTitle+"%").Count(&total)
	return atrList, errmsg.SUCCSE, total
}
