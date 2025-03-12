package models

import (
	"sync"
	"time"
)

const (
	OrderTime  = "time"
	OrderScore = "score"
)

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

// 向前端响应ai对话结构体
type AiResponse struct {
	Answer string `json:"answer"` // AI 的回答
}

// 查询社区的响应体
type CommunityResponse struct {
	CommunityID   int64  `json:"id" db:"community_id"`
	CommunityName string `json:"name" db:"community_name"`
}

// 查询社区详情的响应体
type CommunityDetailResponse struct {
	ID           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"` //omitempty表示如果此字段为空，则序列化时不展示
	CreateTime   time.Time `json:"created_at" db:"created_at"`
}

// 查询帖子详情的响应体
type PostDetailResponse struct {
	AuthorName               string             `json:"author_name"`
	*Post                    `json:"post"`      // 嵌入帖子结构体
	*CommunityDetailResponse `json:"community"` // 嵌入社区信息
	VoteNum                  int64              `json:"vote_num"` // 赞成票数
}

// 投票数据
type VoteData struct {
	//UserID 从请求中获取当前的用户， 不用在结构体中添加
	PostID    string `json:"post_id" bind:"required"`              //帖子id
	Direction int8   `json:"direction,string" bind:"oneof=-1 0 1"` //赞成票(1)or反对票(-1)or取消投票(0)
	//oneof=-1,0,1表示该字段的值要求只可能是-1或0或1
}

// 批量查询帖子的请求体，get请求参数拼接在url中不是json
type PostListParam struct {
	Page  int64  `json:"page" form:"page"`   //查询第几页的数据
	Size  int64  `json:"size" form:"size"`   //每页数据条数
	Order string `json:"order" form:"order"` //排序方式
}

// 根据社区id查询该社区下帖子详情的请求体
type CommunityPostListParam struct {
	PostListParam       //嵌入获取帖子列表的参数的结构体
	CommunityID   int64 `json:"community_id" form:"community_id"` //社区id
}
