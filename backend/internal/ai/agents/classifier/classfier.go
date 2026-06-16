package classifier

import (
	"context"
	"kcardDesgin/backend/internal/ai/agents"
	"kcardDesgin/backend/internal/ai/model"

	"github.com/cloudwego/eino/adk"
)

const (
	AgentName        = "doc_classify_agent"
	AgentDescription = "A content organization and classification agent"
)

func NewAgent(ctx context.Context, cfg model.ModelConfig, stub bool) (adk.Agent, error) {
	a, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        AgentName,
		Description: AgentDescription,
		Instruction: `你是一个知识点分类大师，你负责将传入的知识点拆分，分成多个类`,
		Model:       model.NewChatModel(cfg),
	})
	return a, err
}

func NewAgentWithClient(ctx context.Context, cfg model.ModelConfig, stub bool) (model.Client, error) {
	if stub {
		return agents.StubClient{ModelID: cfg.ModelID}, nil
	}
	return agents.EinoClient{Model: model.NewChatModel(cfg), ModelID: cfg.ModelID}, nil
}
