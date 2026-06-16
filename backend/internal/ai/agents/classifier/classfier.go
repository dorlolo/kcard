package classifier

import (
	"context"
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
		Instruction: `You are an expert ticket booker.
Based on the user's request, use the "BookTicket" tool to book tickets.`,
		Model: model.NewChatModel(cfg),
	})
	return a, err
}
