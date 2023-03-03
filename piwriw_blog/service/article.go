package service

import (
	"piwriw_blog/dao/mysql"
	"piwriw_blog/models"
	"piwriw_blog/models/response"
)

func AddArticle(article *models.ParamArticle) error {
	return mysql.InsertArticle(article)
}
func GetAllArticles(page, limit int64) (respArt *response.RespPage, err error) {
	articles, err := mysql.GetAllArticles(page, limit)
	if err != nil {
		return
	}
	total, err := mysql.GetAllArticlesTotal()
	if err != nil {
		return
	}
	respArt = &response.RespPage{
		List:  articles,
		Total: total,
	}
	return
}

func GetArticleListByTitle(pagenum, pagesize int64, title string) (resppage *response.RespPage, err error) {
	articles, err := mysql.GetArticleListByTitle(pagenum, pagesize, title)
	if err != nil {
		return
	}
	total, err := mysql.GetArticlesTotal()
	if err != nil {
		return
	}
	sliceI := make([]interface{}, len(articles))
	for i, val := range articles {
		sliceI[i] = val
	}
	resppage = &response.RespPage{
		List:  sliceI,
		Total: total,
	}
	return
}

func GetArticleById(articleId string) (*models.Article, error) {
	article, err := mysql.GetArticleById(articleId)
	if err != nil {
		return nil, err
	}
	err = mysql.UpdateArticleRead(article.ReadCount, articleId)
	if err != nil {
		return nil, err
	}
	return article, nil
}
func UpdateArticle(artid string, article *models.ParamArticle) (int64, error) {
	return mysql.UpdateArticle(artid, article)
}
func DeleteArticleById(artid string) (int64, error) {
	return mysql.DeleteArticleById(artid)
}

func GetArticlesListByCId(cid string, pagenum, pagesize int64) (respArt *response.RespArticle, err error) {
	articles, err := mysql.GetArticlesListByCId(pagenum, pagesize, cid)
	if err != nil {
		return
	}
	total, err := mysql.GetArticlesTotalByCId(cid)
	if err != nil {
		return
	}
	respArt = &response.RespArticle{
		ArticleList: articles,
		Total:       total,
	}
	return
}
