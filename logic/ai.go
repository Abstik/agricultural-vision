package logic

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"sync"

	"agricultural_vision/constants"
	"agricultural_vision/models/request"
	"agricultural_vision/models/response"
	"agricultural_vision/settings"
)

var userConversations = make(map[string]*response.Conversation) // 使用 map 保存每个用户的对话历史
var mutex = sync.Mutex{}                                        // 保护 map 的并发访问

func AiTalk(aiRequest *request.AiRequest, userID int64) (aiResponse *response.AiResponse, err error) {
	aiResponse = new(response.AiResponse)
	id := strconv.FormatInt(userID, 10)

	// 获取或创建该用户的对话历史
	mutex.Lock() // 锁住整个 map，确保线程安全
	conversation, exists := userConversations[id]
	if !exists {
		// 用户没有对话历史，创建一个新的
		conversation = &response.Conversation{
			Messages: []response.Message{
				{Content: settings.Conf.AiConfig.SystemContent, Role: "system"},
			},
		}
		userConversations[id] = conversation
	}
	mutex.Unlock()

	// 将用户输入添加到对话历史中
	conversation.Mutex.Lock()
	conversation.Messages = append(conversation.Messages, response.Message{Content: aiRequest.UserInput, Role: "user"})
	conversation.Mutex.Unlock()

	// 向 DeepSeek AI 发送请求
	apiKey := settings.Conf.AiConfig.ApiKey
	apiURL := settings.Conf.AiConfig.ApiUrl

	// 构建请求体
	body := map[string]interface{}{
		"messages":   conversation.Messages,
		"model":      "deepseek-chat",
		"max_tokens": 100,
		"stream":     false,
	}

	// 序列化请求体
	jsonData, err := json.Marshal(body)
	if err != nil {
		return
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// 读取响应体
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// 解析 AI 响应
	var apiResponse response.ApiResponse
	err = json.Unmarshal(bodyBytes, &apiResponse)
	if err != nil {
		return
	}

	// 获取 AI 的回答
	if len(apiResponse.Choices) > 0 {
		aiAnswer := apiResponse.Choices[0].Message.Content

		// 将 AI 的回答添加到对话历史中
		conversation.Mutex.Lock()
		conversation.Messages = append(conversation.Messages, response.Message{Content: aiAnswer, Role: "assistant"})
		conversation.Mutex.Unlock()

		// 返回 AI 的回答给前端
		aiResponse.Answer = aiAnswer
		return
	} else {
		return nil, constants.ErrorAiNotAnswer
	}
}
