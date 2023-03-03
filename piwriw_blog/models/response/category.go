package response

import "piwriw_blog/models"

type RespCategory struct {
	CategoryList []models.Category `json:"category_list"`
	Total        int               `json:"total"`
}
