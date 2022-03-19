package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "go博客"
	indexData.Desc = "入门教程"
	t := template.New("index.html")
	// 1.拿到当前的的路径
	path, _ := os.Getwd()
	// 访问博客首页模板的时候,因为有多个模板的嵌套,解析文件的时候,需要将其涉及到的所有模板进行解析
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t, _ = t.ParseFiles(path+"/template/ubdex.html", home, header, footer, personal, post, pagination)

	// 页面上涉及到的所有的数据,必须有定义
	t.Execute(w, indexData)
}

func main() {
	// 程序入口, 一个项目只能有一个入口
	// web程序,http协议 ip prot
	server := http.Server{
		Addr: "127.0.0.1:8999",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Print(err)
	}
}
