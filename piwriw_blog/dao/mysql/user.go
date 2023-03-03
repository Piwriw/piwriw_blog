package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"piwriw_blog/models"
	"time"
)

const secret = "Pirwiw.com"

// encryptPassword : 加密密码
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
func Login(user *models.ParamUser) (*models.User, error) {
	oPassword := user.PassWord
	sqlStr := `select id,username,password from user where username=? and delete_time is  null`
	u := new(models.User)
	err := db.Get(u, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return u, ErrorUserNotExist
	}
	if err != nil {
		return u, err
	}
	// password is T
	password := encryptPassword(oPassword)
	if password != u.Password {
		return u, ErrorInvalidPassword
	}
	return u, nil
}
func DeleteUserById(uid string) error {
	sqlStr := `update user set delete_time=? where id=?`
	_, err := db.Exec(sqlStr, time.Now(), uid)
	return err
}
func InsertUser(user *models.ParamUser) error {
	err := CheckUseIsExist(user)
	if err != nil {
		return err
	}
	sqlStr := `insert into user(create_time,username,password,email,role) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, time.Now(), user.UserName, encryptPassword(user.PassWord), user.Email, user.Role)
	if err != nil {
		return err
	}
	return err

}
func GetUserById(uid string) (*models.User, error) {
	sqlStr := `select * from user where id=?`
	user := new(models.User)
	err := db.Get(user, sqlStr, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CheckUseIsExist(user *models.ParamUser) (err error) {
	sqlStr := `select count(id) from user where username=? or email=? `
	var count int
	err = db.Get(&count, sqlStr, user.UserName, user.Email)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func GetAllUsers(pagenum, pagesize int64) ([]models.User, error) {
	sqlStr := `select * from user where delete_time is null limit ?,?`
	var users []models.User
	err := db.Select(&users, sqlStr, (pagenum-1)*pagesize, pagesize)
	if err != nil {
		return users, err
	}
	return users, err
}

func UpdateUserById(uid string, user *models.ParamEditUser) error {
	sqlStr := `update user set username=? ,role=? where id=?`
	_, err := db.Exec(sqlStr, user.UserName, user.Role, uid)
	if err != nil {
		return err
	}
	return nil
}
func UpdateUserPassword(password, id string) error {
	sqlStr := `update user set update_time=?,password=? where id =?`
	_, err := db.Exec(sqlStr, time.Now(), encryptPassword(password), id)
	if err != nil {
		return err
	}
	return nil
}
func GetUserTotal() (int, error) {
	sqlStr := `select count(id) from user`
	var count int
	err := db.Get(&count, sqlStr)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func GetUserList(pagenum, pagesize int64, username string) ([]*models.User, error) {
	sqslStr := `select * from user where username like ? and delete_time is null limit ?,?`
	var users []*models.User
	err := db.Select(&users, sqslStr, "%"+username+"%", (pagenum-1)*pagesize, pagesize)
	if err != nil {
		return users, err
	}
	return users, err
}
