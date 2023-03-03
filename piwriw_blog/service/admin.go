package service

import (
	"piwriw_blog/dao/mysql"
	"piwriw_blog/models"
	"piwriw_blog/models/response"
	"piwriw_blog/pkg/jwt"
)

func AdminLogin(login *models.ParamLogin) (respUser *response.RespUser, err error) {
	user, err := mysql.LoginAdmin(login)
	if err != nil {
		return
	}
	token, err := jwt.GenToken(user.ID, user.UserName)
	if err != nil {
		return
	}
	respUser = &response.RespUser{
		Id:       user.ID,
		UserName: user.UserName,
		Token:    token,
	}
	return
}
