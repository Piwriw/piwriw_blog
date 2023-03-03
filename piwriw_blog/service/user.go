package service

import (
	"piwriw_blog/dao/mysql"
	"piwriw_blog/models"
	"piwriw_blog/models/response"
	"piwriw_blog/pkg/jwt"
)

func Register(user *models.ParamUser) error {
	return mysql.InsertUser(user)
}
func GetAllUsers(pagenum, pagesize int64) ([]models.User, error) {
	return mysql.GetAllUsers(pagenum, pagesize)
}
func Login(user *models.ParamUser) (respUser *response.RespUser, err error) {
	dbuser, err := mysql.Login(user)
	if err != nil {
		return
	}
	token, err := jwt.GenToken(dbuser.ID, dbuser.UserName)
	if err != nil {
		return
	}
	respUser = &response.RespUser{
		Id:       dbuser.ID,
		UserName: dbuser.UserName,
		Token:    token,
	}
	return
}
func DeleteUserById(uid string) error {
	return mysql.DeleteUserById(uid)
}
func GetUserById(uid string) (*models.User, error) {
	return mysql.GetUserById(uid)
}
func UpdateUserById(uid string, user *models.ParamEditUser) error {
	return mysql.UpdateUserById(uid, user)
}
func UpdateUserPassword(password, id string) error {
	return mysql.UpdateUserPassword(password, id)
}
func GetUserList(pagenum, pagesize int64, username string) (resppage *response.RespPage, err error) {
	total, err := mysql.GetUserTotal()
	if err != nil {
		return
	}
	userList, err := mysql.GetUserList(pagenum, pagesize, username)
	sliceI := make([]interface{}, len(userList))
	for i, val := range userList {
		sliceI[i] = val
	}
	resppage = &response.RespPage{
		List:  sliceI,
		Total: total,
	}
	return
}
