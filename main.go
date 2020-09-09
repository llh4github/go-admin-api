package main

import (
	"fmt"
	"os"

	"github.com/llh4github/go-admin-api/global"
)

func main() {

	global.Router.Run(":8080")
}

// appLocal 当前项目所处文件系统的位置
// 绝对路径
type appLocal string

func appPath() appLocal {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前项目目录失败！%v \n", err)
		panic("获取当前项目目录失败！")
	}
	global.AppPath = dir
	return appLocal(dir)
}
