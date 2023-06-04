package req

// GetFrontArts 前台条件查询文章
type GetFrontArts struct {
	PageQuery
	CategoryId int `form:"category_id"`
	TagId      int `form:"tag_id"`
}
