package models

type Comment struct {
	Id           int64   `json:"id" db:"id"`
	UserId       uint    `json:"user_id" db:"user_id"`
	ArticleId    uint    `json:"article_id" db:"article_id"`
	ArticleTitle *string `json:"article_title" db:"article_title"`
	Username     string  `json:"username" db:"username"`
	Content      string  `json:"content" db:"content"`
	Status       int8    `json:"status" db:"status"`
	Title        *string `json:"title" db:"title"`
	*Time
}
