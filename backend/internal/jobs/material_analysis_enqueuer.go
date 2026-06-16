// Package jobs 提供任务队列相关的功能，包括任务定义、队列操作和任务调度。
package jobs

import (
	"context"

	"kcardDesgin/backend/internal/service"
)

// DefaultQueueName 常量表示默认的队列名称。
// JobTypeMaterialAnalysis 常量表示资料分析任务的类型标识。
const (
	DefaultQueueName        = "kcard:jobs"
	JobTypeMaterialAnalysis = "material_analysis"
)

// MaterialAnalysisEnqueuer 负责将资料分析任务加入队列。
type MaterialAnalysisEnqueuer struct {
	Queue Queue
}

// NewMaterialAnalysisEnqueuer 创建一个新的资料分析任务入队器。
func NewMaterialAnalysisEnqueuer(queue Queue) MaterialAnalysisEnqueuer {
	return MaterialAnalysisEnqueuer{Queue: queue}
}

// EnqueueMaterialAnalysis 将资料分析任务加入队列并返回任务 ID。
func (e MaterialAnalysisEnqueuer) EnqueueMaterialAnalysis(ctx context.Context, input service.MaterialAnalysisRequest) (string, error) {
	payload := MaterialAnalysisPayload{WorkspaceID: string(input.WorkspaceID), MaterialID: string(input.MaterialID), MaterialVersionID: string(input.MaterialVersionID), Text: input.Text, ContentLocation: input.ContentLocation, Prompt: input.Prompt}
	job, err := e.Queue.Enqueue(ctx, JobTypeMaterialAnalysis, "material_analysis:"+payload.WorkspaceID+":"+payload.MaterialID+":"+payload.MaterialVersionID, payload)
	if err != nil {
		return "", err
	}
	return job.ID, nil
}
