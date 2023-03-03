package mysql

import (
	"database/sql"
	"piwriw_blog/models"
)

func LoginAdmin(login *models.ParamLogin) (user *models.User, err error) {
	err = CheckAdminIsExist(login)
	if err != nil {
		return nil, err
	}
	sqlStr := `select *  from user where username=? and password=? and role=1 and delete_time is  null `
	user = new(models.User)
	err = db.Get(user, sqlStr, login.Username, encryptPassword(login.Password))
	if err == sql.ErrNoRows {
		return nil, ErrorInvalidPassword
	}
	if err != nil {
		return nil, err
	}
	return user, err
}
func CheckAdminIsExist(user *models.ParamLogin) (err error) {
	sqlStr := `select count(id) from user where username=? and role=1`
	var count int
	err = db.Get(&count, sqlStr, user.Username)
	if err != nil {
		return err
	}
	if count < 0 {
		return ErrorUserNotExist
	}
	return
}
