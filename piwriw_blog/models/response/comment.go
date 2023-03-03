package response

import "piwriw_blog/models"

type RespComment struct {
	CommentList *[]models.Comment `json:"comment_list"`
	Total       int               `json:"total"`
}
