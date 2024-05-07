package dto

type Pagination struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}

func (m *Pagination) GetPageIndex() int {
	if m.PageNum <= 0 {
		m.PageNum = 1
	}
	return m.PageNum
}

func (m *Pagination) GetPageSize() int {
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	return m.PageSize
}
