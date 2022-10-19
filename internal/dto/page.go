package dto

type Page struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func (p *Page) GetLimit() int {
	return p.PageSize
}

func (p *Page) GetOffset() int {
	return p.PageSize * (p.Page - 1)
}
