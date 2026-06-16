// Package ai 提供 AI 工作流引擎的核心抽象，包括结构化请求/响应模型和客户端接口。
package ai

import (
	"kcardDesgin/backend/internal/ai/agents"
)

// StubClient 是模拟客户端，用于测试场景。
// StubClient 是 agents.StubClient 的类型别名。
type StubClient = agents.StubClient

// EinoClient 是基于 Eino 框架的 AI 客户端实现。
// EinoClient 是 agents.EinoClient 的类型别名。
type EinoClient = agents.EinoClient
