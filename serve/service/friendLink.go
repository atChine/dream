package service

import (
	"dream/serve/dao"
	"dream/serve/model"
)

type FriendLink struct {
}

// GetFrontList 获取前台友链链接
func (*FriendLink) GetFrontList() (data []model.FriendLink) {
	return dao.List([]model.FriendLink{}, "*", "", "")
}
