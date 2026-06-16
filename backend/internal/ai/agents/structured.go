package agents

import (
	"context"
	"encoding/json"
	"errors"
	"kcardDesgin/backend/internal/ai"
	"strings"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

type EinoClient struct {
	Model   model.BaseChatModel
	ModelID string
}

func (c EinoClient) GenerateStructured(ctx context.Context, req ai.StructuredRequest) (ai.StructuredResponse, error) {
	if c.Model == nil {
		return ai.StructuredResponse{}, errors.New("eino chat model is not configured")
	}
	if len(req.Schema) > 0 && !json.Valid(req.Schema) {
		return ai.StructuredResponse{}, errors.New("structured output schema is not valid JSON")
	}
	messages := make([]*schema.Message, 0, len(req.Messages)+1)
	if strings.TrimSpace(req.System) != "" {
		messages = append(messages, schema.SystemMessage(buildStructuredSystem(req)))
	}
	for _, msg := range req.Messages {
		switch strings.ToLower(msg.Role) {
		case "system":
			messages = append(messages, schema.SystemMessage(msg.Content))
		case "assistant":
			messages = append(messages, schema.AssistantMessage(msg.Content, nil))
		default:
			messages = append(messages, schema.UserMessage(msg.Content))
		}
	}
	opts := []model.Option{}
	if req.MaxTokens > 0 {
		opts = append(opts, model.WithMaxTokens(req.MaxTokens))
	}
	response, err := c.Model.Generate(ctx, messages, opts...)
	if err != nil {
		return ai.StructuredResponse{}, err
	}
	jsonBytes, err := extractJSONObject(response.Content)
	if err != nil {
		return ai.StructuredResponse{}, err
	}
	return ai.StructuredResponse{JSON: jsonBytes, ModelID: c.ModelID, StopReason: "stop"}, nil
}

func buildStructuredSystem(req ai.StructuredRequest) string {
	var builder strings.Builder
	builder.WriteString(req.System)
	builder.WriteString("\n\n请严格输出 JSON 对象，不要输出 Markdown、解释或代码块。")
	if req.SchemaName != "" {
		builder.WriteString("\n输出用途：")
		builder.WriteString(req.SchemaName)
	}
	if len(req.Schema) > 0 {
		builder.WriteString("\nJSON Schema：")
		builder.Write(req.Schema)
	}
	return builder.String()
}

func extractJSONObject(text string) ([]byte, error) {
	text = strings.TrimSpace(text)
	text = strings.TrimPrefix(text, "```json")
	text = strings.TrimPrefix(text, "```")
	text = strings.TrimSuffix(text, "```")
	text = strings.TrimSpace(text)
	start := strings.Index(text, "{")
	end := strings.LastIndex(text, "}")
	if start < 0 || end < start {
		return nil, errors.New("model response does not contain a JSON object")
	}
	candidate := []byte(text[start : end+1])
	if !json.Valid(candidate) {
		return nil, errors.New("model response JSON is invalid")
	}
	return candidate, nil
}
