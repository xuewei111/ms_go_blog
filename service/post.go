package service

import (
	"html/template"
	"ms-go-blog/config"
	"ms-go-blog/dao"
	"ms-go-blog/models"
)

func GetPostDetail(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pid)
	if err != nil {
		return nil, err
	}

	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)

	postMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	postMore.Pid = post.Pid
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}

	return postRes, err

}
