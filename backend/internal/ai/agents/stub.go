package agents

import (
	"context"
	"kcardDesgin/backend/internal/ai/model"
)

type StubClient struct{ ModelID string }

func (c StubClient) GenerateStructured(ctx context.Context, req model.StructuredRequest) (model.StructuredResponse, error) {
	return model.StructuredResponse{JSON: []byte(`{}`), ModelID: c.ModelID, StopReason: "stubbed"}, nil
}
