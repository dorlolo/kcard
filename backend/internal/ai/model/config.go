// Package model 提供 AI 模型配置和创建函数，支持多种模型提供商。
package model

import "time"

// ModelConfig 定义模型配置，包括提供商、API 密钥、模型 ID、基础 URL 和超时时间。
type ModelConfig struct {
	Provider string
	APIKey   string
	ModelID  string
	BaseURL  string
	Timeout  time.Duration
}
