package model

type Page struct {
	Page     int
	PageSize int
	Total    int
}

func (p Page) Limit() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

func (p Page) Skip() int {
	if p.Page == 0 {
		return 20
	}
	return p.Page * p.PageSize
}
