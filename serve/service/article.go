package service

import (
	"dream/serve/dao"
	"dream/serve/model"
	"dream/serve/model/req"
	"dream/serve/model/resp"
	"dream/serve/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Article struct{}

/* 前台接口 */

// GetFrontList 获取前台文章列表
func (*Article) GetFrontList(req req.GetFrontArts) []resp.FrontArticleVO {
	list, _ := articleDao.GetFrontList(req)
	return list
}

func (*Article) GetFrontInfo(c *gin.Context, id int) resp.FrontArticleDetailVO {
	// 查询最新文章
	article := articleDao.GetInfoById(id)
	// 查询推荐文章 6篇
	article.RecommendArticles = articleDao.GetRecommendList(id, 6)
	// 查询最新文章 5篇
	article.NewestArticles = articleDao.GetNewesList(5)
	// * 目前请求一次就会增加访问量, 即刷新可以刷访问量
	utils.Redis.ZincrBy(KEY_ARTICLE_VIEW_COUNT, strconv.Itoa(id), 1)
	// 获取上一篇文章, 下一篇文章
	article.LastArticle = articleDao.GetLast(id)
	article.NextArticle = articleDao.GetNext(id)
	// 点赞量, 浏览量
	article.ViewCount = utils.Redis.ZScore(KEY_ARTICLE_VIEW_COUNT, strconv.Itoa(id))
	article.LikeCount = utils.Redis.HGet(KEY_ARTICLE_LIKE_COUNT, strconv.Itoa(id))
	// 评论数量
	article.CommentCount = int(commentDao.GetArticleCommentCount(id))
	return article
}

// Search 前台文章搜索
func (*Article) Search(q req.KeywordQuery) []resp.ArticleSearchVO {
	res := make([]resp.ArticleSearchVO, 0)
	if q.Keyword == "" {
		return res
	}
	articleList := dao.List([]model.Article{}, "*", "", "is_delete = 0 AND status = 1 AND (title LIKE ? OR content LIKE ?)", "%"+q.Keyword+"%", "%"+q.Keyword+"%")
	for _, article := range articleList {
		// 高亮标题中的关键字
		title := strings.ReplaceAll(article.Title, q.Keyword,
			"<span style='color:#7cfc00'>"+q.Keyword+"</span>")

		content := article.Content
		// 关键字在内容中的起始位置
		keywordStartIndex := unicodeIndex(content, q.Keyword)
		if keywordStartIndex != -1 { // 关键字在内容中
			preIndex, afterIndex := 0, 0
			if keywordStartIndex > 25 {
				preIndex = keywordStartIndex - 25
			}
			// 防止中文截取出乱码 (中文在 golang 是 3 个字符, 使用 rune 中文占一个数组下标)
			preText := substring(content, preIndex, keywordStartIndex)
			// string([]rune(content[preIndex:keywordStartIndex]))

			// 关键字在内容中的结束位置
			keywordEndIndex := keywordStartIndex + unicodeLen(q.Keyword)
			afterLength := len(content) - keywordEndIndex
			if afterLength > 175 {
				afterIndex = keywordEndIndex + 175
			} else {
				afterIndex = keywordEndIndex + afterLength
			}
			// afterText := string([]rune(content)[keywordStartIndex:afterIndex])
			afterText := substring(content, keywordStartIndex, afterIndex)
			// 高亮内容中的关键字
			content = strings.ReplaceAll(preText+afterText, q.Keyword,
				"<span style='color:#7cfc00'>"+q.Keyword+"</span>")
		}

		res = append(res, resp.ArticleSearchVO{
			ID:      article.ID,
			Title:   title,
			Content: content,
		})
	}

	return res
}

// GetArchiveList 获取归档列表
func (*Article) GetArchiveList(req req.GetFrontArts) resp.PageResult[[]resp.ArchiveVO] {
	articles, total := articleDao.GetArchiveList(req)
	archives := make([]resp.ArchiveVO, 0)
	for _, article := range articles {
		archives = append(archives, resp.ArchiveVO{
			ID:         article.ID,
			Title:      article.Title,
			Created_at: article.Created_at,
		})
	}
	return resp.PageResult[[]resp.ArchiveVO]{
		Total:    total,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		List:     archives,
	}
}

// 获取带中文的字符串中子字符串的实际位置，非字节位置
func unicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result > 0 {
		prefix := []byte(str)[0:result]
		rs := []rune(string(prefix))
		result = len(rs)
	}
	return result
}

// 获取带中文的字符串实际长度，非字节长度
func unicodeLen(str string) int {
	var r = []rune(str)
	return len(r)
}

// 解决中文获取位置不正确问题
func substring(source string, start int, end int) string {
	var unicodeStr = []rune(source)
	length := len(unicodeStr)
	if start >= end {
		return ""
	}
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start <= 0 && end >= length {
		return source
	}
	var substring = ""
	for i := start; i < end; i++ {
		substring += string(unicodeStr[i])
	}
	return substring
}
