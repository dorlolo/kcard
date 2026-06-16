// Package jobs 提供任务队列相关的功能，包括任务定义、队列操作和任务调度。
package jobs

import (
	"context"
	"errors"
	"strings"

	"kcardDesgin/backend/internal/ai"
	"kcardDesgin/backend/internal/domain"
	"kcardDesgin/backend/internal/service"
)

// MaterialAnalysisWorker 处理资料分析任务的工作器，执行 AI 分类并将结果持久化。
type MaterialAnalysisWorker struct {
	Workflow  ai.ClassificationWorkflow
	Materials service.MaterialStore
	Knowledge service.KnowledgeStore
}

// MaterialAnalysisPayload 表示资料分析任务的负载参数，包含工作空间、资料和提示信息。
type MaterialAnalysisPayload struct {
	WorkspaceID       string `json:"workspaceId"`
	MaterialID        string `json:"materialId"`
	MaterialVersionID string `json:"materialVersionId"`
	Text              string `json:"text"`
	ContentLocation   string `json:"contentLocation"`
	Prompt            string `json:"prompt"`
}

// Handle 执行资料分析任务的处理逻辑，包括参数校验、AI 分类和知识点的持久化。
func (w MaterialAnalysisWorker) Handle(ctx context.Context, payload MaterialAnalysisPayload) (ai.ClassificationOutput, error) {
	if strings.TrimSpace(payload.WorkspaceID) == "" || strings.TrimSpace(payload.MaterialID) == "" || strings.TrimSpace(payload.MaterialVersionID) == "" {
		return ai.ClassificationOutput{}, errors.New("workspaceID, materialID, and materialVersionID are required")
	}
	if strings.TrimSpace(payload.Text) == "" {
		return ai.ClassificationOutput{}, errors.New("material analysis text is empty")
	}
	workspaceID := domain.ID(payload.WorkspaceID)
	materialID := domain.ID(payload.MaterialID)
	versionID := domain.ID(payload.MaterialVersionID)
	if w.Materials != nil {
		_ = w.Materials.UpdateMaterialStatus(ctx, workspaceID, materialID, domain.MaterialProcessing, "")
	}
	output, err := w.Workflow.Classify(ctx, payload.Text, payload.Prompt)
	if err != nil {
		if w.Materials != nil {
			_ = w.Materials.UpdateMaterialStatus(ctx, workspaceID, materialID, domain.MaterialFailed, err.Error())
		}
		return ai.ClassificationOutput{}, err
	}
	for _, draft := range output.KnowledgePoints {
		content := strings.TrimSpace(draft.Content)
		if content == "" {
			continue
		}
		summary := strings.TrimSpace(draft.Summary)
		point := domain.KnowledgePoint{LearnerWorkspaceID: workspaceID, SourceMaterialID: materialID, MaterialVersionID: versionID, Content: content, Summary: summary, Notes: aiTagNote(draft.Tags, draft.Confidence), ApprovalStatus: domain.KnowledgeNeedsReview, CreationSource: domain.CreationAIGenerated, GraphLabel: graphLabel(summary, content)}
		if w.Knowledge != nil {
			if _, err := w.Knowledge.Create(ctx, point); err != nil {
				if w.Materials != nil {
					_ = w.Materials.UpdateMaterialStatus(ctx, workspaceID, materialID, domain.MaterialFailed, err.Error())
				}
				return ai.ClassificationOutput{}, err
			}
		}
	}
	if w.Materials != nil {
		_ = w.Materials.UpdateMaterialStatus(ctx, workspaceID, materialID, domain.MaterialNeedsReview, "")
	}
	return output, nil
}

func graphLabel(summary string, content string) string {
	if summary != "" {
		return summary
	}
	runes := []rune(content)
	if len(runes) <= 24 {
		return content
	}
	return string(runes[:24]) + "…"
}

func aiTagNote(tags []string, confidence float64) string {
	parts := []string{}
	if len(tags) > 0 {
		parts = append(parts, "AI tags: "+strings.Join(tags, ", "))
	}
	if confidence > 0 {
		parts = append(parts, "confidence recorded")
	}
	return strings.Join(parts, "; ")
}
