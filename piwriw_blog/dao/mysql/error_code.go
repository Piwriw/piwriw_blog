package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
	ErrorArticleUpdate   = errors.New("文章更新失败")
	ErrorArticleID       = errors.New("无效的Article ID")
	ErrorCategortID      = errors.New("无效的Category ID")
)
