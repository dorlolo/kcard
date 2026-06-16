package agents

import (
	"context"
	"kcardDesgin/backend/internal/ai"
)

type StubClient struct{ ModelID string }

func (c StubClient) GenerateStructured(ctx context.Context, req ai.StructuredRequest) (ai.StructuredResponse, error) {
	return ai.StructuredResponse{JSON: []byte(`{}`), ModelID: c.ModelID, StopReason: "stubbed"}, nil
}
