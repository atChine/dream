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

// GetInfoById 根据 id 获取文章详情
func (*Article) GetInfoById(id int) (res resp.FrontArticleDetailVO) {
	DB.Table("article").Preload("Category").Preload("Tags").Where("id = ? AND is_delete = 0 AND status = 1", id).First(&res)
	return res
}

// GetRecommendList 查询 n 篇推荐文章 (根据标签)
func (*Article) GetRecommendList(id, n int) []resp.RecommendArticleVO {
	list := make([]resp.RecommendArticleVO, 0)
	// sub1: 查出标签id列表
	// SELECT tag_id FROM `article_tag` WHERE `article_id` = ?
	sub1 := DB.Table("article_tag").
		Select("tag_id").
		Where("article_id  = ?", id)
	// sub2: 查出这些标签对应的文章id列表 (去重, 且不包含当前文章)
	// SELECT DISTINCT article_id FROM (sub1) t
	// JOIN article_tag t1 ON t.tag_id = t1.tag_id
	// WHERE `article_id` != ?
	sub2 := DB.Table("(?) t1", sub1).Select("DISTINCT article_id").Joins("JSON article_tag t ON t.tag_id = t1.tag_id").Where("article_id = ?", id)
	// 根据 文章id列表 查出文章信息 (前 6 个)
	DB.Table("(?) t2", sub2).
		Select("id, title, img, created_at").
		Joins("JOIN article a ON t2.article_id = a.id").
		Where("a.is_delete = 0").
		Order("is_top, id DESC").
		Limit(n).
		Find(&list)
	return list
}

// GetNewesList 查询最新的前5个文章
func (*Article) GetNewesList(n int) []resp.RecommendArticleVO {
	list := make([]resp.RecommendArticleVO, 0)
	DB.Table("article").
		Select("id, title, img, created_at").
		Where("is_delete = 0 AND status = 1").
		Order("created_at DESC, id ASC").
		Limit(n).Find(&list)
	return list
}

// GetLast 获取上一篇文章
func (*Article) GetLast(id int) (res resp.ArticlePaginationVO) {
	// 执行子查询获取上一篇文章的 ID
	sub := DB.Table("article").
		Select("max(id)").
		Where("id < ?", id)
	DB.Table("article").
		Select("id, title, img, created_at").
		Where("is_delete = 0 AND status = 1 AND id = (?)", sub).
		Limit(1).
		First(&res)
	return
}

// GetNext 取上一篇文章
func (*Article) GetNext(id int) (res resp.ArticlePaginationVO) {
	// 执行子查询获取上一篇文章的 ID
	sub := DB.Table("article").
		Select("min(id)").
		Where("id > ?", id)
	DB.Table("article").
		Select("id, title, img, created_at").
		Where("is_delete = 0 AND status = 1 AND id = (?)", sub).
		Limit(1).
		First(&res)
	return
}
