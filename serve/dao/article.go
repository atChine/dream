package dao

import (
	"dream/serve/model/req"
	"dream/serve/model/resp"
)

type Article struct{}

/* 前台相关 */

// GetFrontList 获取前台文章列表
func (*Article) GetFrontList(req req.GetFrontArts) ([]resp.FrontArticleVO, int64) {
	list := make([]resp.FrontArticleVO, 0)
	var total int64

	db := DB.Table("article").
		Select("id, title, content, img, type, is_top, created_at, category_id").
		Where("is_delete = 0 AND status = 1")
	if req.CategoryId != 0 {
		db = db.Where("category_id", req.CategoryId)
	}
	if req.TagId != 0 {
		db = db.Where("id IN (SELECT article_id FROM article_tag WHERE tag_id = ?)", req.TagId)
	}

	db.Count(&total)
	db.Preload("Tags").
		Preload("Category").
		Order("is_top DESC, id DESC").
		Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1)).
		Find(&list)
	return list, total
}
