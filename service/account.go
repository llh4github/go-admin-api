package service

import (
	"github.com/llh4github/go-admin-api/model"
	"github.com/llh4github/go-admin-api/utils"
	"github.com/llh4github/go-admin-api/vo"
)

// Account 帐户信息服务层
type Account struct {
	User // service.User
}

// NewAccountService 创建帐户信息服务层
func NewAccountService() Account {
	return Account{User{}}
}

// Register 注册一个新帐户
func (a Account) RegisterAccount(acc vo.Account) bool {
	hashed, err := utils.HashPassword(acc.Password)
	if err != nil {
		log.Error(err)
		panic("密码加密错误 !")
	}
	u := model.User{
		Username: acc.Username,
		Password: hashed,
	}
	return a.Add(u)
}
