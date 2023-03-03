package service

import (
	"piwriw_blog/dao/mysql"
	"piwriw_blog/models"
	"piwriw_blog/models/response"
)

func GetAllCommentList(pagenum int64, pagesize int64, aid string) (respcomment *response.RespComment, err error) {
	comments, err := mysql.GetCommentByIdPage(pagenum, pagesize, aid)
	if err != nil {
		return
	}
	total, err := mysql.GetAllCommentTotal()
	if err != nil {
		return
	}
	respcomment = &response.RespComment{CommentList: &comments, Total: total}
	return
}

func GetAllComments(pagenum int64, pagesize int64) ([]models.Comment, error) {
	return mysql.GetAllComments(pagenum, pagesize)
}
func GetCommentById(id string) (*models.Comment, error) {
	return mysql.GetCommentById(id)
}
func CheckComment(id, status string) error {
	return mysql.UpdateComment(id, status)
}
func AddComment(comment *models.ParamComment) error {
	return mysql.InsertComment(comment)
}
