// Package ai 提供 AI 工作流引擎的核心抽象，包括结构化请求/响应模型和客户端接口。
package ai

import (
	"context"
	"encoding/json"
	"kcardDesgin/backend/internal/ai/model"

	"kcardDesgin/backend/internal/ai/schemas"
)

// WorkflowKind 表示工作流类型，用于区分不同的 AI 工作流。
type WorkflowKind string

const (
	// WorkflowClassification 表示知识分类工作流类型。
	WorkflowClassification WorkflowKind = "classification"
	// WorkflowCardGeneration 表示卡片生成工作流类型。
	WorkflowCardGeneration WorkflowKind = "card_generation"
	// WorkflowPlanGeneration 表示复习计划生成工作流类型。
	WorkflowPlanGeneration WorkflowKind = "plan_generation"
	// WorkflowPlanOptimization 表示复习计划优化工作流类型。
	WorkflowPlanOptimization WorkflowKind = "plan_optimization"
)

// WorkflowInput 定义工作流输入，包含工作空间 ID、提示词、来源 ID 列表和负载数据。
type WorkflowInput struct {
	WorkspaceID string
	Prompt      string
	SourceIDs   []string
	Payload     map[string]any
}

// WorkflowOutput 定义工作流输出，包含草稿列表和警告信息。
type WorkflowOutput struct {
	Drafts   []map[string]any
	Warnings []string
}

// Workflows 封装 AI 客户端，提供执行各种工作流的通用能力。
type Workflows struct{ Client model.Client }

// Run 执行指定类型的工作流，返回工作流输出或错误。
func (w Workflows) Run(ctx context.Context, kind WorkflowKind, input WorkflowInput) (WorkflowOutput, error) {
	schemaBytes := schemaForKind(kind)
	payload, _ := json.Marshal(input.Payload)
	resp, err := w.Client.GenerateStructured(ctx, model.StructuredRequest{System: string(kind), Messages: []model.Message{{Role: "user", Content: input.Prompt + "\n" + string(payload)}}, SchemaName: string(kind), Schema: schemaBytes, MaxTokens: 4096})
	if err != nil {
		return WorkflowOutput{}, err
	}
	var raw map[string]any
	if len(resp.JSON) > 0 {
		if err := json.Unmarshal(resp.JSON, &raw); err != nil {
			return WorkflowOutput{}, err
		}
	}
	return WorkflowOutput{Drafts: collectDraftMaps(raw), Warnings: collectWarnings(raw)}, nil
}

func schemaForKind(kind WorkflowKind) []byte {
	switch kind {
	case WorkflowCardGeneration:
		return schemas.DeckGeneration
	case WorkflowPlanGeneration:
		return schemas.ReviewPlan
	case WorkflowPlanOptimization:
		return schemas.PlanOptimization
	default:
		return schemas.KnowledgeClassification
	}
}

func collectDraftMaps(raw map[string]any) []map[string]any {
	var drafts []map[string]any
	for _, key := range []string{"drafts", "knowledgePoints", "decks", "cards", "days", "changes"} {
		items, ok := raw[key].([]any)
		if !ok {
			continue
		}
		for _, item := range items {
			if m, ok := item.(map[string]any); ok {
				drafts = append(drafts, m)
			}
		}
	}
	return drafts
}

func collectWarnings(raw map[string]any) []string {
	items, ok := raw["warnings"].([]any)
	if !ok {
		return nil
	}
	warnings := make([]string, 0, len(items))
	for _, item := range items {
		if value, ok := item.(string); ok {
			warnings = append(warnings, value)
		}
	}
	return warnings
}
