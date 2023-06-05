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

// FrontArticleDetailVO 前端详情vo
type FrontArticleDetailVO struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Img         string `json:"img"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Type        int    `json:"type"`
	OriginalUrl string `json:"original_url"`

	CommentCount int `json:"comment_count"` // 评论数量
	LikeCount    int `json:"like_count"`    // 点赞数量
	ViewCount    int `json:"view_count"`    // 访问数量

	CategoryId int            `json:"category_id"`
	Category   model.Category `gorm:"foreignkey:CategoryId;" json:"category"`
	Tags       []model.Tag    `gorm:"many2many:article_tag;joinForeignKey:article_id" json:"tags"`

	LastArticle ArticlePaginationVO `gorm:"-" json:"last_article"` // 上一篇
	NextArticle ArticlePaginationVO `gorm:"-" json:"next_article"` // 下一篇

	RecommendArticles []RecommendArticleVO `gorm:"-" json:"recommend_articles"` // 推荐文章
	NewestArticles    []RecommendArticleVO `gorm:"-" json:"newest_articles"`    // 最新文章
	// Desc         string   `json:"desc"`
}

// ArticlePaginationVO 文章详情界面: 上一篇文章, 下一篇文章显示, 只需要 标题, 封面
type ArticlePaginationVO struct {
	ID    int    `json:"id"`
	Img   string `json:"img"`
	Title string `json:"title"`
}

// RecommendArticleVO 推荐文章
type RecommendArticleVO struct {
	ID        int       `json:"id"`
	Img       string    `json:"img"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

// ArticleSearchVO 文章搜索结果
type ArticleSearchVO struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
