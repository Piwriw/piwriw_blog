package response

import "piwriw_blog/models"

type RespArticle struct {
	ArticleList []models.PoArticle `json:"article_list"`
	Total       int                `json:"total"`
}
type Time struct {
	CreateTime *string `json:"create_time"  db:"create_time"`
	UpdateTime *string `json:"update_time"  db:"update_time"`
	DeleteTime *string `json:"delete_time"  db:"delete_time"`
}
