// Package ark 提供火山引擎 Ark 模型的原生 HTTP 客户端实现，兼容 Eino 模型接口。
package ark

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

// ChatModel 表示火山引擎 Ark 聊天模型，封装 API 密钥、模型名称、基础 URL 和 HTTP 客户端。
type ChatModel struct {
	apiKey string
	model  string
	base   string
	http   *http.Client
}

// NewChatModel 根据模型配置创建并返回一个新的火山引擎 Ark 聊天模型实例。
func NewChatModel(cfg aiModel.ModelConfig) *ChatModel {
	baseURL := strings.TrimRight(cfg.BaseURL, "/")
	if baseURL == "" {
		baseURL = "https://ark.cn-beijing.volces.com/api/v3"
	}
	timeout := cfg.Timeout
	if timeout <= 0 {
		timeout = 60 * time.Second
	}
	return &ChatModel{apiKey: cfg.APIKey, model: cfg.ModelID, base: baseURL, http: &http.Client{Timeout: timeout}}
}

type request struct {
	Model          string    `json:"model"`
	Messages       []message `json:"messages"`
	MaxTokens      int       `json:"max_tokens,omitempty"`
	ResponseFormat *struct {
		Type string `json:"type"`
	} `json:"response_format,omitempty"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type response struct {
	ID      string `json:"id"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    string `json:"code"`
	} `json:"error,omitempty"`
}

// Generate 执行聊天生成，向 Ark API 发送请求并返回响应消息。
func (m *ChatModel) Generate(ctx context.Context, input []*schema.Message, opts ...model.Option) (*schema.Message, error) {
	if strings.TrimSpace(m.apiKey) == "" {
		return nil, errors.New("ARK_API_KEY is required")
	}
	modelID := m.model
	common := model.GetCommonOptions(nil, opts...)
	if common.Model != nil && *common.Model != "" {
		modelID = *common.Model
	}
	if strings.TrimSpace(modelID) == "" {
		return nil, errors.New("ARK_MODEL is required")
	}
	requestBody := request{Model: modelID, Messages: toMessages(input), ResponseFormat: &struct {
		Type string `json:"type"`
	}{Type: "json_object"}}
	if common.MaxTokens != nil {
		requestBody.MaxTokens = *common.MaxTokens
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, m.base+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+m.apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := m.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var decoded response
	if err := json.Unmarshal(data, &decoded); err != nil {
		return nil, fmt.Errorf("decode ark response: %w", err)
	}
	if resp.StatusCode >= 400 {
		if decoded.Error != nil && decoded.Error.Message != "" {
			return nil, fmt.Errorf("ark api error %d: %s", resp.StatusCode, decoded.Error.Message)
		}
		return nil, fmt.Errorf("ark api error %d: %s", resp.StatusCode, string(data))
	}
	if len(decoded.Choices) == 0 {
		return nil, errors.New("ark response has no choices")
	}
	choice := decoded.Choices[0]
	return &schema.Message{Role: schema.Assistant, Content: choice.Message.Content, ResponseMeta: &schema.ResponseMeta{FinishReason: choice.FinishReason}}, nil
}

// Stream 执行流式生成，当前 Ark 实现返回未实现的错误。
func (m *ChatModel) Stream(ctx context.Context, input []*schema.Message, opts ...model.Option) (*schema.StreamReader[*schema.Message], error) {
	return nil, errors.New("ark streaming is not implemented")
}

func toMessages(messages []*schema.Message) []message {
	out := make([]message, 0, len(messages))
	for _, msg := range messages {
		if msg == nil {
			continue
		}
		role := string(msg.Role)
		if role == "" {
			role = "user"
		}
		out = append(out, message{Role: role, Content: msg.Content})
	}
	return out
}
