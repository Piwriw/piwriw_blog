package service

import (
	"piwriw_blog/dao/mysql"
	"piwriw_blog/models"
	"piwriw_blog/models/response"
)

func GetCategoryList(pagenum, pagesize int64) (respCategory *response.RespCategory, err error) {
	total, err := mysql.GetAllCategoryTotal()
	if err != nil {
		return
	}
	category, err := mysql.GetCategoryList(pagenum, pagesize)
	if err != nil {
		return
	}
	respCategory = &response.RespCategory{
		CategoryList: category,
		Total:        total,
	}
	return
}

func GetAllCategory() ([]models.Category, error) {
	return mysql.GetAllCategory()
}
func GetArticleByCateId(cid string) ([]models.Article, error) {
	return mysql.GetArticleByCateId(cid)
}
func AddCategory(category *models.ParamCategory) error {
	return mysql.InsertCategory(category)
}
func UpdateCategory(category *models.ParamCategory) error {
	return mysql.UpdateCategory(category)
}
func GetCategoryById(cid string) (error, *models.Category) {
	return mysql.GetCategoryById(cid)
}
func DeleteCategoryById(cid string) error {
	return mysql.DeleteCategoryById(cid)
}
