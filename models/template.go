package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		w.Write([]byte("error"))
	}
}

func InitTemplate(templateDir string) (HtmlTemplate, error) {
	blogs, err := readTemplate([]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templateDir)

	var htmlTemplate HtmlTemplate
	if err != nil {
		return htmlTemplate, err
	}

	htmlTemplate.Index = blogs[0]
	htmlTemplate.Category = blogs[1]
	htmlTemplate.Custom = blogs[2]
	htmlTemplate.Detail = blogs[3]
	htmlTemplate.Login = blogs[4]
	htmlTemplate.Pigeonhole = blogs[5]
	htmlTemplate.Writing = blogs[6]

	return htmlTemplate, nil

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

func DateDay(layout string) string {
	return time.Now().Format(layout)
}

func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog

	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		// 访问博客首页模板的时候,因为有多个模板的嵌套,解析文件的时候,需要将其涉及到的所有模板进行解析
		home := templateDir + "home.html"
		header := templateDir + "layout\\header.html"
		footer := templateDir + "layout\\footer.html"
		personal := templateDir + "layout\\personal.html"
		post := templateDir + "layout\\post-list.html"
		pagination := templateDir + "layout\\pagination.html"

		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		t, err := t.ParseFiles(templateDir+viewName, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Println("解析模板出错", err)
			return nil, err
		}

		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}

	return tbs, nil

}
