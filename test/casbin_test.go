package test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/llh4github/go-admin-api/global"
	"github.com/llh4github/go-admin-api/wirex"
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前项目目录失败！%v \n", err)
		log.Fatal("获取当前项目目录失败！")
	}
	global.AppPath = strings.TrimSuffix(dir, "test")
	fmt.Println("global.AppPath : ", global.AppPath)
	wirex.InitBase()
}

// 测试
func TestAcfun(t *testing.T) {
	t.Log(global.Enforcer)

}
