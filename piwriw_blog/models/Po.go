package models

type PoArticle struct {
	Id           int    `json:"id" db:"id"`
	Title        string `json:"title" db:"title"`
	Desc         string `json:"desc" db:"desc"`
	HtmlContent  string `json:"html_content" db:"html_content"`
	Content      string `json:"content" db:"content"`
	Img          string `json:"img" db:"img"`
	CommentCount int    `json:"comment_count" db:"comment_count"`
	ReadCount    int    `json:"read_count" db:"read_count"`
	TagName      string `json:"tag_name" db:"tag_name"`
	*Time
}
