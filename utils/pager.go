package utils

// GetOffsetLimit 获取分页的offset及limit
func GetOffsetLimit(page, pageSize int) (offset, limit int) {
	if pageSize < 1 {
		pageSize = 10
	}
	return pageSize * (page - 1), pageSize
}
