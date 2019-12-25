package common

import "math"

type Meta struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	Limit     int64 `json:"limit"`
	TotalPage int64 `json:"total_page"`
	Offset    int64 `json:"offset"`
}

func MakeMeta(total, page, limit int64) (m *Meta) {
	m = &Meta{}
	if total < 1 {
		return
	}

	m.Total = total
	if page < 1 {
		page = 1
	}
	m.Page = page
	if limit < 1 {
		limit = 20
	}
	if limit > 10 {
		limit = 10
	}
	m.Limit = limit
	m.Offset = limit * (page - 1)
	m.TotalPage = int64(math.Ceil(float64(total) / float64(limit)))

	return
}
