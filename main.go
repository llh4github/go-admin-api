package main

import (
	"fmt"
	"log"
	"os"

	"github.com/llh4github/go-admin-api/global"
	"github.com/llh4github/go-admin-api/wirex"
)

func main() {
	appPath()
	wirex.InitBase()
	global.Router.Run(":8080")
}

// 获取当前项目目录绝对路径
func appPath() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前项目目录失败！%v \n", err)
		log.Fatal("获取当前项目目录失败！")
	}
	global.AppPath = dir
}
