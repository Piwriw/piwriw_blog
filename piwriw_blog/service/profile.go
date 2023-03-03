package service

import (
	"piwriw_blog/dao/mysql"
	"piwriw_blog/models"
)

func GetProfile(id string) (*models.Profile, error) {
	return mysql.GetProfile(id)
}
func UpdateProfile(id string, profile *models.Profile) error {
	return mysql.UpdateProfile(id, profile)
}
