package ai

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"kcardDesgin/backend/internal/ai/schemas"
)

type KnowledgePointDraft struct {
	Content    string   `json:"content"`
	Summary    string   `json:"summary"`
	Tags       []string `json:"tags"`
	Confidence float64  `json:"confidence"`
}
type ClassificationOutput struct {
	KnowledgePoints []KnowledgePointDraft `json:"knowledgePoints"`
	Warnings        []string              `json:"warnings"`
}

type ClassificationWorkflow struct {
	Client        Client
	DefaultPrompt string
}

func (w ClassificationWorkflow) Classify(ctx context.Context, materialText, prompt string) (ClassificationOutput, error) {
	if strings.TrimSpace(materialText) == "" {
		return ClassificationOutput{}, errors.New("material text is empty")
	}
	if prompt == "" {
		prompt = w.DefaultPrompt
	}
	resp, err := w.Client.GenerateStructured(ctx, StructuredRequest{System: prompt, Messages: []Message{{Role: "user", Content: materialText}}, SchemaName: "knowledge_classification", Schema: schemas.KnowledgeClassification, MaxTokens: 4096})
	if err != nil {
		return ClassificationOutput{}, err
	}
	var out ClassificationOutput
	if len(resp.JSON) > 2 {
		if err := json.Unmarshal(resp.JSON, &out); err != nil {
			return ClassificationOutput{}, err
		}
	}
	if len(out.KnowledgePoints) == 0 {
		out.KnowledgePoints = append(out.KnowledgePoints, KnowledgePointDraft{Content: strings.TrimSpace(materialText), Summary: "Draft extracted from material", Confidence: 0.5})
	}
	return out, nil
}
