package dto

type PageResponse[T any] struct {
	Data         []T          `json:"data"`
	PageMetadata PageMetadata `json:"paging"`
}

type PageMetadata struct {
	Page      uint `json:"page" form:"page"`
	Size      uint `json:"size" form:"size"`
	TotalItem uint `json:"total_item"`
	TotalPage uint `json:"total_page"`
}

func (p *PageMetadata) Populate() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Size < 1 {
		p.Size = 1
	}
}