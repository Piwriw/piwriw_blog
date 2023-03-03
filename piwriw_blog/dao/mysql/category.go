package mysql

import (
	"piwriw_blog/models"
	"time"
)

func GetAllCategoryTotal() (int, error) {
	sqlStr := `select count(id) from category WHERE
	delete_time is  NULL`
	var count int
	err := db.Get(&count, sqlStr)
	if err != nil {
		return 0, err
	}
	return count, err
}
func GetAllCategory() ([]models.Category, error) {
	sqlStr := `select * from category WHERE
	delete_time is  NULL`
	var categories []models.Category
	err := db.Select(&categories, sqlStr)
	if err != nil {
		return categories, err
	}
	return categories, err
}

func GetCategoryList(pagenum, pagesize int64) ([]models.Category, error) {
	sqlStr := `select * from category WHERE
	delete_time is  NULL  order by create_time desc limit ?,?`
	var categories []models.Category
	err := db.Select(&categories, sqlStr, (pagenum-1)*pagesize, pagesize)
	if err != nil {
		return categories, err
	}
	return categories, err
}
func GetArticleByCateId(cid string) ([]models.Article, error) {
	sqlStr := `SELECT * from article WHERE cid=? and delete_time is null `
	var aricles []models.Article
	err := db.Select(&aricles, sqlStr, cid)
	if err != nil {
		return aricles, err
	}
	return aricles, nil
}
func InsertCategory(category *models.ParamCategory) error {
	sqlStr := `insert into category(create_time,name) values (?,?)`
	_, err := db.Exec(sqlStr, time.Now(), category.Name)
	if err != nil {
		return err
	}
	return nil
}
func UpdateCategory(category *models.ParamCategory) error {
	sqlStr := `update category set name=? ,update_time=?,where id=?`
	ret, err := db.Exec(sqlStr, category.Name, time.Now(), category.Id)
	if err != nil {
		return err
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 0 {
		return ErrorCategortID
	}
	return nil
}
func GetCategoryById(cid string) (err error, category *models.Category) {
	sqlStr := `select * from category where id=? and delete_time is null`
	category = new(models.Category)
	err = db.Get(category, sqlStr, cid)
	if err != nil {
		return
	}
	return
}
func DeleteCategoryById(cid string) error {
	sqlStr := `update category set delete_time=? where id=?`
	_, err := db.Exec(sqlStr, time.Now(), cid)
	if err != nil {
		return err
	}
	return nil
}
