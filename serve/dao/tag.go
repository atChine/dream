package dao

import (
	"dream/serve/model/req"
	"dream/serve/model/resp"
)

type Tag struct {
}

// GetFrontList 前台标签列表
func (*Tag) GetFrontList(req req.PageQuery) ([]resp.TagVO, int64) {
	var list = make([]resp.TagVO, 0)
	var total int64
	db := DB.Table("tag t").Select("t.id", "t.name", "t.created_at", "t.updated_at", "COUNT(at.article_id) AS article_count").
		Joins("LEFT JOIN article_tag at ON t.id = at.tag_id")
	if req.Keyword != "" {
		db = db.Where("name LIKE ?", "%"+req.Keyword+"%")
	}
	// 判断是否需要分页查询
	if req.PageSize > 0 && req.PageNum > 0 {
		db = db.Limit(req.PageSize).Offset(req.PageSize * (req.PageNum - 1))
	}
	db.Group("t.id").Order("t.id DESC").Count(&total).Find(&list)
	return list, total
}
