// Package jobs 提供任务队列相关的功能，包括任务定义、队列操作和任务调度。
package jobs

import "context"

// ExportWorker 处理资料导出任务的工作器。
type ExportWorker struct{}

// Handle 执行资料导出任务的处理逻辑。
func (ExportWorker) Handle(ctx context.Context, payload map[string]any) (map[string]any, error) {
	return map[string]any{"status": "accepted"}, nil
}
