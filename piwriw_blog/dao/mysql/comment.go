package mysql

import (
	"piwriw_blog/models"
	"time"
)

func GetAllCommentTotal() (int, error) {
	sqlStr := `select count(id) from comment where delete_time is null and status  `
	var count int
	err := db.Get(&count, sqlStr)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func GetCommentByIdPage(pagenum int64, pagesize int64, aid string) (comments []models.Comment, err error) {
	sqlStr := `select * from comment where delete_time is null and status=1 and  article_id=? LIMIT ?,?`
	err = db.Select(&comments, sqlStr, aid, (pagenum-1)*pagesize, pagesize)
	if err != nil {
		return
	}
	return
}
func GetAllComments(pagenum int64, pagesize int64) (comments []models.Comment, err error) {
	sqlStr := `select * from comment where delete_time is null LIMIT ?,?`
	err = db.Select(&comments, sqlStr, (pagenum-1)*pagesize, pagesize)
	if err != nil {
		return
	}
	return
}
func GetCommentById(id string) (comment *models.Comment, err error) {
	sqlStr := `select * from comment where delete_time is null and id=? and status=1`
	comment = new(models.Comment)
	err = db.Get(comment, sqlStr, id)
	if err != nil {
		return
	}
	return
}
func UpdateComment(id, status string) error {
	sqlStr := `update comment set status=? where id=?`
	_, err := db.Exec(sqlStr, status, id)
	if err != nil {
		return err
	}
	return nil
}
func InsertComment(comment *models.ParamComment) error {
	sqlStr := `insert into comment(create_time,article_id,content,user_id,username,status) values (?,?,?,?,?,?)`
	_, err := db.Exec(sqlStr, time.Now(), comment.ArticleId, comment.Content, comment.UserID, comment.UserName, 0)
	if err != nil {
		return err
	}
	return nil
}
