package models

// 定义请求的参数结构体（DTO）

// 用户注册请求参数
type SignUpParam struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 用户登录请求参数
type LoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 发送验证码结构体
type SendVerificationCodeParam struct {
	Email string `json:"email"`
}

// 修改密码结构体
type ChangePasswordParam struct {
	Email    string `json:"email"`
	Code     string `json:"code"` // 邮箱验证码
	Password string `json:"password"`
}

// 修改用户信息结构体
type UpdateUserInfoParam struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
