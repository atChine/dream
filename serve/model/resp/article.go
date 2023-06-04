package resp

import (
	"dream/serve/model"
	"time"
)

// 前台首页文章列表 VO

// FrontArticleVO 多对多 返回前台列表vo
type FrontArticleVO struct {
	ID         int             `json:"id"`
	CreatedAt  time.Time       `json:"created_at"`
	Title      string          `json:"title"`
	Desc       string          `json:"desc"`
	Content    string          `json:"content"`
	Img        string          `json:"img"`
	IsTop      int             `json:"is_top"`
	Type       int             `json:"type"`
	CategoryId int             `json:"category_id"`
	Category   *model.Category `gorm:"foreignkey:CategoryId;" json:"category"`
	Tags       []*model.Tag    `gorm:"many2many:article_tag;joinForeignKey:article_id" json:"tags"`
}
