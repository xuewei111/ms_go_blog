package main

import (
	"log"
	"ms-go-blog/common"
	"ms-go-blog/router"
	"net/http"
)

func init() {
	// 模板加载
	common.LoadTemplate()
}

func main() {

	// 程序入口, 一个项目只能有一个入口
	// web程序,http协议 ip prot
	server := http.Server{
		Addr: "127.0.0.1:8999",
	}

	// 路由
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Print(err)
	}
}
