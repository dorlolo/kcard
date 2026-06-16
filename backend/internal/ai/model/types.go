// Package ai 提供 AI 工作流引擎的核心抽象，包括结构化请求/响应模型和客户端接口。
package model

import (
	"context"
)

// Message 表示一条消息，包含角色和内容。
type Message struct {
	Role    string
	Content string
}

// StructuredRequest 定义结构化生成请求，包含系统提示、消息列表、Schema 名称及内容、最大 Token 数。
type StructuredRequest struct {
	System     string
	Messages   []Message
	SchemaName string
	Schema     []byte
	MaxTokens  int
}

// StructuredResponse 定义结构化生成响应，包含返回的 JSON 数据、模型 ID 和停止原因。
type StructuredResponse struct {
	JSON       []byte
	ModelID    string
	StopReason string
}

// Client 定义 AI 客户端接口，提供结构化生成能力。
type Client interface {
	GenerateStructured(ctx context.Context, req StructuredRequest) (StructuredResponse, error)
}
