package models

import "sync"

// 定义请求的参数结构体（DTO）

// 用户注册请求参数
type SignUpParam struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Code     string `json:"code" binding:"required"` // 邮箱验证码
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

// 向AI请求体结构体
type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

// 响应体结构体
type Choice struct {
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
}

// 接收AI响应的结构体
type ApiResponse struct {
	Choices []Choice `json:"choices"`
}

// 定义用于保存对话上下文的结构体
type Conversation struct {
	Messages []Message
	Mutex    sync.Mutex // 确保线程安全
}

// 接收前端请求的结构体
type AiRequest struct {
	UserInput string `json:"user_input"` // 前端传来的问题
}

// 向前端响应结构体
type AiResponse struct {
	Answer string `json:"answer"` // AI 的回答
}
