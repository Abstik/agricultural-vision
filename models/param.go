package models

// 定义请求的参数结构体（DTO）

// 发送验证码结构体
type SendVerificationCodeParam struct {
	Email string `json:"email"`
}

// SignUpParam 用户注册请求参数
type SignUpParam struct {
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	VerificationCode string `json:"verificationCode"`
}

// LoginParam 用户登录请求参数
type LoginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
