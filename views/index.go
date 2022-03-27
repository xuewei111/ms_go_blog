package views

import (
	"errors"
	"log"
	"ms-go-blog/common"
	"ms-go-blog/service"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	// 页面上涉及到的所有的数据,必须有定义
	// 数据库查询
	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("Index 获取数据出错:", err)
		index.WirteError(w, errors.New("系统错误,请联系管理员"))
	}

	// 页面上涉及到的所有的数据,必须有定义
	index.WriteData(w, hr)
}
