package service

import (
	"dream/serve/model/req"
	"dream/serve/model/resp"
)

type Article struct{}

/* 前台接口 */

// GetFrontList 获取前台文章列表
func (*Article) GetFrontList(req req.GetFrontArts) []resp.FrontArticleVO {
	list, _ := articleDao.GetFrontList(req)
	return list
}
