package test

import (
	"testing"

	"github.com/llh4github/go-admin-api/global"
	"github.com/llh4github/go-admin-api/model"
	"github.com/llh4github/go-admin-api/utils"
)

// 测试
func TestPage2(t *testing.T) {
	db := global.MyDB
	r := []model.Role{}
	_db := db.Where("remove_flag = false")
	p, _ := utils.PageDefault(_db, 1, &r)
	t.Logf("TotalPage : %v ,TotalRecord : %v ,Records : %v ,PageSize:%v",
		p.TotalPage, p.TotalRecord, p.Records, p.PageSize)
}
