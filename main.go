package main

import (
	"application/app/global"
	"application/bootstrap"
	"application/router"
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取根路径
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	// 初始化环境
	filePath := bootstrap.EnvInit(root)
	// 初始化数据库
	err = bootstrap.PostgresInit(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	// minio初始化
	err = bootstrap.MinioInit(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 初始化http client
	err = bootstrap.HttpClientInit(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 初始化定时任务
	// task.Init()
	// 初始化gin服务器
	r := global.Router()
	// 初始化路由
	err = bootstrap.ContextPathInit(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	router.Route(r, bootstrap.ContextPath)
	err = r.Run(fmt.Sprintf(":%s", bootstrap.Server.Env.Port))
	if err != nil {
		return
	}
}
