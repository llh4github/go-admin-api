package test

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/llh4github/go-admin-api/global"
	"github.com/llh4github/go-admin-api/wirex"
)

var e *casbin.Enforcer

func init() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前项目目录失败！%v \n", err)
		log.Fatal("获取当前项目目录失败！")
	}
	global.AppPath = strings.TrimSuffix(dir, "test")
	wirex.InitBase()

	e = global.Enforcer
}
