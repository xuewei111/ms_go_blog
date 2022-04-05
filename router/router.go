package router

import (
	"ms-go-blog/api"
	"ms-go-blog/views"
	"net/http"
)

func Router() {
	//1.页面 views  2。 api数据 (json) 3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	//http://localhost:8080/c/1  1参数 分类的id
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("api/v1/post", api.API.SaveAndUpdatePost)
	http.Handle("/resource/", http.StripPrefix("/resource", http.FileServer(http.Dir("public/resource/"))))
}
