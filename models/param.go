package models

// 定义请求的参数结构体（DTO）

// 用户注册请求参数
type SignUpParam struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 用户登录请求参数
type LoginParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 发送验证码结构体
type SendVerificationCodeParam struct {
	Email string `json:"email" binding:"required"`
}

// 修改密码结构体
type ChangePasswordParam struct {
	Email    string `json:"email" binding:"required"`
	Code     string `json:"code" binding:"required"` // 邮箱验证码
	Password string `json:"password" binding:"required"`
}

// 修改用户信息结构体
type UpdateUserInfoParam struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
