package service


// 注册请求参数
type UserSignUpInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	RePassword string `json:"repassword"`
}