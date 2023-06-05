package dao

import "dream/serve/model"

type Comment struct {
}

// GetArticleCommentCount 获取评论数量
func (*Comment) GetArticleCommentCount(articleId int) (count int64) {
	DB.Model(&model.Comment{}).
		Where("topic_id = ? AND type = 1 AND is_review = 1", articleId).
		Count(&count)
	return
}
