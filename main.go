package main

import (
	"html/template"
	"log"
	"ms-go-blog/config"
	"ms-go-blog/models"
	"net/http"
	"os"
	"time"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index]
}

func Date(layout string) string {
	return time.Now().Format(layout)
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

	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)

	if err != nil {
		log.Println(err)
	}
	// 页面上涉及到的所有的数据,必须有定义
	categorys := []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}

	posts := []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}

	hr := &models.HomeData{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}

	// 页面上涉及到的所有的数据,必须有定义
	t.Execute(w, hr)
}

func main() {
	// 程序入口, 一个项目只能有一个入口
	// web程序,http协议 ip prot
	server := http.Server{
		Addr: "127.0.0.1:8999",
	}

	http.HandleFunc("/", index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

	if err := server.ListenAndServe(); err != nil {
		log.Print(err)
	}
}
