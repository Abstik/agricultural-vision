package logic

import (
	"bytes"
	"io"
	"net/http"

	"github.com/goccy/go-json"

	"agricultural_vision/models"
)

func AiTalk(aiRequest *models.AiRequest) (aiResponse *models.AiResponse, err error) {
	// 构建请求体
	messages := []models.Message{
		{"你是一个农业小助手", "system"},
		{aiRequest.UserInput, "user"},
	}

	// 向 DeepSeek AI 发送请求
	apiKey := "sk-0a03ab3a2d18455e97d0a40f4fc20671"
	apiURL := "https://api.deepseek.com/v1/chat/completions"

	// 构建请求体
	body := map[string]interface{}{
		"messages":   messages,
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
	var apiResponse models.ApiResponse
	err = json.Unmarshal(bodyBytes, &apiResponse)
	if err != nil {
		return
	}

	// 获取 AI 的回答
	if len(apiResponse.Choices) > 0 {
		aiAnswer := apiResponse.Choices[0].Message.Content
		// 返回 AI 的回答给前端
		aiResponse.Answer = aiAnswer
		return
	} else {
		return nil, models.ErrorAiNotAnswer
	}
}
