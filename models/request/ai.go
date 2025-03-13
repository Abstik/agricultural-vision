package request

// ai对话
type AiRequest struct {
	UserInput string `json:"user_input" binding:"required"` // 前端传来的问题
}
