package service

import (
	"errors"
	"ms-go-blog/dao"
	"ms-go-blog/models"
	"ms-go-blog/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "qwef")
	user, err := dao.GetUser(userName, passwd)
	if err != nil {
		return nil, errors.New("账号密码不正确")
	}

	uid := user.Uid
	// 生成token，jwt技术进行生成 令牌 3段(头部,载荷,验签)
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token生成未成功！")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil

}
