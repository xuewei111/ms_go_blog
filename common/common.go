package common

import (
	"ms-go-blog/config"
	"ms-go-blog/models"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		// 耗时
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		group.Done()
	}()
	group.Wait()
}
