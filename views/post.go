package views

import (
	"errors"
	"ms-go-blog/common"
	"ms-go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail

	// 获取路径参数
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	// 7.html
	pIdStr = strings.TrimSuffix(pIdStr, ".html")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此请求路径!"))
		return
	}

	postRes, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteData(w, errors.New("查询出错"))
	}

	detail.WriteData(w, postRes)

}
