package utils

// 基于github.com/mgw2007/gorm-paginator/paginator修改
// 原版针对gorm老版本1.9实现

import (
	"math"

	"github.com/llh4github/go-admin-api/vo"
	"gorm.io/gorm"
)

// pageScope 分页函数。比直接写limit语句要灵活点
func pageScope(param *vo.PageParam) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := param.Page
		if page == 0 {
			page = 1
		}
		pageSize := param.Limit
		switch {
		case pageSize > 100: // 每页最大100条。看情况修改
			pageSize = 100
		case pageSize <= 0: // 默认每页10条。
			pageSize = 10
		}
		param.Limit = pageSize
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
func countRecords(db *gorm.DB, anyType interface{}, count *int64) {
	db.Model(anyType).Count(count)

}

// PageDefault 分页方法(使用默认设置)
func PageDefault(db *gorm.DB, curPage int, result interface{}) (*vo.Paginator, error) {
	param := vo.PageParam{
		DB:   db,
		Page: curPage,
	}
	return Page(&param, result)
}

// Page 分页方法
// 去掉了协程方法
func Page(p *vo.PageParam, result interface{}) (*vo.Paginator, error) {
	db := p.DB
	var count int64

	countRecords(db, result, &count)
	var paginator vo.Paginator

	res := db.Scopes(pageScope(p)).Find(result)

	if res.Error != nil {
		return nil, res.Error
	}
	paginator.Records = result
	paginator.PageSize = p.Limit
	paginator.TotalRecord = int(count)
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))
	return &paginator, nil

}
