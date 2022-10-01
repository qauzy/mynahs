package model

type Page struct {
	Page     int `gorm:"column:page" json:"page"`
	PageSize int `gorm:"column:page_size" json:"pageSize"`
}

func (p *Page) GetLimit() int {
	return p.PageSize
}

func (p *Page) GetOffset() int {
	return p.PageSize * (p.Page - 1)
}
