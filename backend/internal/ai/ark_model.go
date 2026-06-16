// Package ai 提供 AI 工作流引擎的核心抽象，包括结构化请求/响应模型和客户端接口。
package ai

import arkmodel "kcardDesgin/backend/internal/ai/agents/ark"

// ArkConfig 是火山引擎 Ark 模型配置的类型别名。
// ArkConfig 是 agents/ark.Config 的类型别名。
type ArkConfig = arkmodel.Config

// ArkChatModel 是火山引擎 Ark 聊天模型的类型别名。
// ArkChatModel 是 agents/ark.ChatModel 的类型别名。
type ArkChatModel = arkmodel.ChatModel

// NewArkChatModel 创建并返回一个新的火山引擎 Ark 聊天模型实例。
func NewArkChatModel(cfg ArkConfig) *ArkChatModel {
	return arkmodel.NewChatModel(cfg)
}
