package model

import (
	"reflect"
)

const (
	PUBLIC = iota + 1 // 公开
	SECRET            // 私密
	DRAFT             // 草稿
)

// Article 文章
type Article struct {
	Universal
	CategoryId  int    `gorm:"type:bigint;not null;comment:分类 ID" json:"category_id"`
	UserId      int    `gorm:"type:int;not null;comment:用户 ID" json:"user_id"`
	Title       string `gorm:"type:varchar(100);not null;comment:文章标题" json:"title"`
	Desc        string `gorm:"type:varchar(200);comment:文章描述" json:"desc"`
	Content     string `gorm:"type:longtext;comment:文章内容" json:"content"`
	Img         string `gorm:"type:varchar(100);comment:封面图片地址" json:"img"`
	Type        int8   `gorm:"type:tinyint;comment:类型(1-原创 2-转载 3-翻译)" json:"type"`
	Status      int8   `gorm:"type:tinyint;comment:状态(1-公开 2-私密)" json:"status"`
	IsTop       *int8  `gorm:"type:tinyint;not null;default:0;comment:是否置顶(0-否 1-是)" json:"is_top"`
	IsDelete    *int8  `gorm:"type:tinyint;not null;default:0;comment:是否放到回收站(0-否 1-是)" json:"is_delete"`
	OriginalUrl string `gorm:"type:varchar(100);comment:源链接" json:"original_url"`
}

// IsEmpty 判断当前结构体对象是否为空
func (a *Article) IsEmpty() bool {
	return reflect.DeepEqual(a, &Article{})
}

// ArticleTag 文章-标签 关联表
type ArticleTag struct {
	ArticleId int
	TagId     int
}
