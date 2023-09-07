package controller

// UserRegister
// 用户注册时的结构体
type UserR struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

// 登录时的结构体
type UserL struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
