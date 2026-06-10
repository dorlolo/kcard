package ai

import "context"

type Message struct {
	Role    string
	Content string
}
type StructuredRequest struct {
	System     string
	Messages   []Message
	SchemaName string
	Schema     []byte
	MaxTokens  int
}
type StructuredResponse struct {
	JSON       []byte
	ModelID    string
	StopReason string
}

type Client interface {
	GenerateStructured(ctx context.Context, req StructuredRequest) (StructuredResponse, error)
}

type AnthropicClient struct {
	APIKey  string
	ModelID string
}

func NewAnthropicClient(apiKey, modelID string) AnthropicClient {
	return AnthropicClient{APIKey: apiKey, ModelID: modelID}
}

func (c AnthropicClient) GenerateStructured(ctx context.Context, req StructuredRequest) (StructuredResponse, error) {
	return StructuredResponse{JSON: []byte(`{}`), ModelID: c.ModelID, StopReason: "stubbed"}, nil
}
