package service

import (
	"dream/serve/model/req"
	"dream/serve/model/resp"
)

type Tag struct {
}

// GetFrontList 前台标签列表
func (*Tag) GetFrontList() resp.ListResult[[]resp.TagVO] {
	list, total := tagDao.GetFrontList(req.PageQuery{})
	return resp.ListResult[[]resp.TagVO]{
		Total: total,
		List:  list,
	}
}
