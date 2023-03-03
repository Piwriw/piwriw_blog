package models

// 定义请求参数结构体

type ParamComment struct {
	ArticleId int64  `json:"article_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
	UserID    int64  `json:"user_id" binding:"required"`
	UserName  string `json:"username" binding:"required"`
}
type ParamCategory struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type ParamArticle struct {
	Title       string `json:"title" db:"title" binding:"required"`
	Cid         int    `json:"cid" db:"cid" binding:"required"`
	Desc        string `json:"desc" db:"desc" binding:"required"`
	HtmlContent string `json:"html_content" db:"html_content" binding:"required"`
	Content     string `json:"content" db:"content" binding:"required"`
	Img         string `json:"img" db:"img"`
}
type ParamUser struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
}

type ParamEditUser struct {
	UserName string `json:"username" binding:"required"`
	Role     int    `json:"role" `
}

// ParamLogin :登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
