package views

import (
	"ms-go-blog/common"
	"ms-go-blog/config"
	"ms-go-blog/models"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
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
	index.WriteData(w, hr)
}
