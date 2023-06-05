package service

import (
	"dream/serve/model/req"
	"dream/serve/model/resp"
)

type Category struct {
}

// GetFrontList 前台分类列表
func (*Category) GetFrontList() resp.ListResult[[]resp.CategoryVo] {
	list, total := categoryDao.GetList(req.PageQuery{})
	return resp.ListResult[[]resp.CategoryVo]{
		Total: total,
		List:  list,
	}
}
