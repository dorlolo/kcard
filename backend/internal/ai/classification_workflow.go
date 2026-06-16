// Package ai 提供 AI 工作流引擎的核心抽象，包括结构化请求/响应模型和客户端接口。
package ai

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"kcardDesgin/backend/internal/ai/schemas"
)

// KnowledgePointDraft 表示知识点草稿，包含内容、摘要、标签和置信度。
type KnowledgePointDraft struct {
	Content    string   `json:"content"`
	Summary    string   `json:"summary"`
	Tags       []string `json:"tags"`
	Confidence float64  `json:"confidence"`
}
// ClassificationOutput 定义分类工作流的输出，包含知识点列表和警告信息。
type ClassificationOutput struct {
	KnowledgePoints []KnowledgePointDraft `json:"knowledgePoints"`
	Warnings        []string              `json:"warnings"`
}

// ClassificationWorkflow 封装知识分类工作流，包含 AI 客户端和默认提示词。
type ClassificationWorkflow struct {
	Client        Client
	DefaultPrompt string
}

// Classify 执行知识分类，将材料文本按提示词进行分类，返回分类输出或错误。
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
