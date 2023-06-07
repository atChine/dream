package utils

// HandlePageSizeAndPageNum 处理页面大小和页数
func HandlePageSizeAndPageNum(pageSize, pageNum int) (int, int) {
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	return pageSize, pageNum
}
