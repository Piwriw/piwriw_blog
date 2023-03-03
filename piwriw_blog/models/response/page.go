package response

type RespPage struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
