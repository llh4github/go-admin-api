package vo

import "gorm.io/gorm"

// PageParam 分页参数
type PageParam struct {
	DB      *gorm.DB `json:"-"`
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
	OrderBy []string `json:"orderBy"`
}

// Paginator 分页结果
type Paginator struct {
	TotalRecord int         `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"records"`
	PageSize    int         `json:"page_size"`
}
