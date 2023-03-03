package response

type RespUser struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}
