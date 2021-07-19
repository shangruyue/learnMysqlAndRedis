package service

// 注册请求参数
type UserSignUpInfo struct {
	UserName   string `json:"user_name" binding:"required"`  //binding 用来对参数进行检验 required意思是参数必须 
	Password   string `json:"password"  binding:"required,eqfield=RePassword"`
	RePassword string `json:"repassword"  binding:"required"`
}

//this is a test for tset2
