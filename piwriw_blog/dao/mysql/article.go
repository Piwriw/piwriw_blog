package mysql

import (
	"piwriw_blog/models"
	"time"
)

func InsertArticle(article *models.ParamArticle) (err error) {
	sqlStr := "insert into article(title,cid,`desc`,html_content,content,img,create_time) values(?,?,?,?,?,?,?)"
	_, err = db.Exec(sqlStr, article.Title, article.Cid, article.Desc, article.HtmlContent, article.Content, article.Img,
		time.Now())
	if err != nil {
		return
	}
	return
}

func GetAllArticles(page, limit int64) ([]models.PoArticle, error) {
	sqlStr := `SELECT
	a.id,
	a.create_time AS create_time,
	a.update_time AS update_time,
	a.delete_time AS delete_time,
	title,
	a.desc,
	content,
	img,
	comment_count,
	read_count,
	c.NAME AS tag_name 
FROM
	article a
	LEFT JOIN category c ON a.cid = c.id
	WHERE
	a.delete_time is  NULL 
	order by a.create_time  
	LIMIT ?,?`
	var articles []models.PoArticle
	err := db.Select(&articles, sqlStr, (page-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
func GetAllArticlesTotal() (int, error) {
	sqlStr := `SELECT
	count(a.id)  as total
	FROM
	article a
	LEFT JOIN category c ON a.cid = c.id`
	var total int
	err := db.Get(&total, sqlStr)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetArticleById(articleId string) (article *models.Article, err error) {
	article = new(models.Article)
	sqlStr := `select * from article where id=? and
	delete_time is  NULL`
	err = db.Get(article, sqlStr, articleId)
	if err != nil {
		return
	}
	return
}
func UpdateArticle(aid string, article *models.ParamArticle) (affected int64, err error) {
	sqlStr := "update  article set title=? ,cid=?,`desc`=?,content=?,html_content=?,img=?,update_time=? where id=?"
	ret, err := db.Exec(sqlStr, article.Title, article.Cid, article.Desc, article.Content, article.HtmlContent, article.Img, time.Now(), aid)
	affected, err = ret.RowsAffected()
	if err != nil {
		return
	}
	if affected < 0 {
		return affected, ErrorArticleUpdate
	}
	return
}

func DeleteArticleById(aid string) (affected int64, err error) {
	sqlStr := `update  article set delete_time =? where id=?`
	ret, err := db.Exec(sqlStr, time.Now(), aid)
	if err != nil {
		return
	}
	affected, err = ret.RowsAffected()
	if err != nil {
		return
	}
	if affected < 0 {
		return affected, ErrorArticleID
	}
	return
}
func GetArticlesListByCId(pagenum, pagesize int64, cid string) ([]models.PoArticle, error) {
	sqlStr := `SELECT
	a.id,
	a.create_time AS create_time,
	a.update_time AS update_time,
	a.delete_time AS delete_time,
	title,
	a.desc,
	content,
	img,
	comment_count,
	read_count,
	c.NAME AS tag_name 
FROM
	article a
	LEFT JOIN category c ON a.cid = c.id
	WHERE
	a.delete_time is  NULL and cid=?
	order by create_time desc
	LIMIT ?,?
	`
	var articles []models.PoArticle
	err := db.Select(&articles, sqlStr, cid, (pagenum-1)*pagesize, pagesize)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
func GetArticlesTotalByCId(cid string) (int, error) {
	sqlStr := `select count(id) from article where cid=?`
	var count int
	err := db.Get(&count, sqlStr, cid)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func GetArticlesTotal() (int, error) {
	sqlStr := `select count(id) from article `
	var count int
	err := db.Get(&count, sqlStr)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func GetArticleListByTitle(pagenum, pagesize int64, title string) ([]*models.PoArticle, error) {
	sqlStr := `SELECT
	a.id,
	a.create_time AS create_time,
	a.update_time AS update_time,
	a.delete_time AS delete_time,
	title,
	a.desc,
	content,
	img,
	comment_count,
	read_count,
	c.NAME AS tag_name 
FROM
	article a
	LEFT JOIN category c ON a.cid = c.id
	WHERE
	a.delete_time is  NULL and title like ?
	order by create_time   
	LIMIT ?,?  `
	var artilceList []*models.PoArticle
	err := db.Select(&artilceList, sqlStr, "%"+title+"%", (pagenum-1)*pagesize, pagesize)
	if err != nil {
		return nil, err
	}
	return artilceList, nil
}
func UpdateArticleRead(read int, id string) error {
	sqlStr := `update article set read_count=? where id=?`
	_, err := db.Exec(sqlStr, read+1, id)
	if err != nil {
		return err
	}
	return nil
}
