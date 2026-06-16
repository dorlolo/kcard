// Package model 提供 AI 模型配置和创建函数，支持多种模型提供商。
package model

import (
	"context"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	arkModel "github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

// NewChatModel 根据配置创建并返回指定提供商的聊天模型实例，支持 OpenAI 和火山引擎 Ark。
func NewChatModel(cfg ModelConfig) model.ToolCallingChatModel {
	switch cfg.Provider {
	default:
		cm, err := openai.NewChatModel(context.Background(), &openai.ChatModelConfig{
			APIKey:  cfg.APIKey,
			Model:   cfg.ModelID,
			BaseURL: cfg.BaseURL,
			ByAzure: func() bool {
				return os.Getenv("OPENAI_BY_AZURE") == "true"
			}(),
		})
		if err != nil {
			log.Fatalf("openai.NewChatModel failed: %v", err)
		}
		return cm
	case "ark":
		cm, err := ark.NewChatModel(context.Background(), &ark.ChatModelConfig{
			// Add Ark-specific configuration from environment variables
			APIKey:  cfg.APIKey,
			Model:   cfg.ModelID,
			BaseURL: cfg.BaseURL,
			Thinking: &arkModel.Thinking{
				Type: arkModel.ThinkingTypeAuto,
			},
		})
		if err != nil {
			log.Fatalf("ark.NewChatModel failed: %v", err)
		}
		return cm
	}
}
