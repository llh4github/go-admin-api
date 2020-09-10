package model

// CasbinRule casbin规则表
// 因为有适配器中间件，尽量不对这张表的数据进行操作
// 这里写出来以备额外需求
type CasbinRule struct {
	PType string `json:"p_type" gorm:"size:100;"`
	V0    string `json:"v0" gorm:"size:100;"`
	V1    string `json:"v1" gorm:"size:100;"`
	V2    string `json:"v2" gorm:"size:100;"`
	V3    string `json:"v3" gorm:"size:100;"`
	V4    string `json:"v4" gorm:"size:100;"`
	V5    string `json:"v5" gorm:"size:100;"`
}

// TableName 指定表名 casbin_rule
func (CasbinRule) TableName() string {
	return "casbin_rule"
}
