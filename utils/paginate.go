package utils

func PaginateCalculate(page, limit, total int) (offset, limitOut, totalPage int) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset = (page - 1) * limit

	if total == 0 {
		totalPage = 1
	} else {
		totalPage = total / limit
		if total%limit != 0 {
			totalPage++
		}
	}

	return offset, limit, totalPage
}
