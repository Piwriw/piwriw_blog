package models

type Profile struct {
	ID            int    `db:"id" json:"id"`
	Name          string `db:"name" json:"name"`
	Desc          string `db:"desc" json:"desc"`
	Qqchat        string `db:"qqchat" json:"qq_chat"`
	Wechat        string `db:"wechat" json:"wechat"`
	Weibo         string `db:"weibo" json:"weibo"`
	Bili          string `db:"bili" json:"bili"`
	Email         string `db:"email" json:"email"`
	Img           string `db:"img" json:"img"`
	Avatar        string `db:"avatar" json:"avatar"`
	GitHub        string `db:"github" json:"github"`
	GitEE         string `db:"gitee" json:"gitee"`
	PublicAccount string `db:"public_account" json:"public_account"`
}
