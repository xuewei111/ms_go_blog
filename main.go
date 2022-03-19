package main

import (
	"encoding/json"
	"fmt"
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
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "go博客"
	indexData.Desc = "入门教程"

	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "go博客"
	indexData.Desc = "入门教程"

	t := template.New("index.html")
	// 获取当前路径
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + `/template/index.html`)

	err := t.Execute(w, indexData)
	fmt.Println(err)
}

func main() {
	// 程序入口, 一个项目只能有一个入口
	// web程序,http协议 ip prot
	server := http.Server{
		Addr: "127.0.0.1:8999",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)

	if err := server.ListenAndServe(); err != nil {
		log.Print(err)
	}
}
