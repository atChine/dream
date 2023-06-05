package dao

import (
	"dream/serve/model/req"
	"dream/serve/model/resp"
)

type Category struct{}

// GetList 前后台分类列表
func (*Category) GetList(req req.PageQuery) ([]resp.CategoryVo, int64) {
	var datas = make([]resp.CategoryVo, 0)
	var total int64
	db := DB.Table("category c").
		Select("c.id", "c.name", "COUNT(a.id) AS article_count", "c.created_at", "c.updated_at").
		Joins("LEFT JOIN article a ON c.id = a.category_id AND a.is_delete = 0 AND a.status = 1")
	// 条件查询
	if req.Keyword != "" {
		db = db.Where("name LIKE ?", "%"+req.Keyword+"%")
	}
	// 判断是否需要分页查询
	if req.PageSize > 0 && req.PageNum > 0 {
		db = db.Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1))
	}
	db.Group("c.id").
		Order("c.id DESC").
		Count(&total).
		Find(&datas)
	return datas, total
}
