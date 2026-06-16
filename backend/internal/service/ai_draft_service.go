// Package service 提供领域服务层的实现，包括资料、卡片、知识图谱等核心业务逻辑。
package service

import "kcardDesgin/backend/internal/domain"

// AIDraftService 处理 AI 生成的草稿的审批与废弃操作。
type AIDraftService struct{}

// Approve 将 AI 草稿标记为已审批并关联记录 ID。
func (AIDraftService) Approve(draft domain.AIDraft, recordID domain.ID) domain.AIDraft {
	draft.Status = "approved"
	draft.Payload = map[string]any{"approvedRecordId": recordID}
	return draft
}
// Discard 将 AI 草稿标记为已废弃。
func (AIDraftService) Discard(draft domain.AIDraft) domain.AIDraft {
	draft.Status = "discarded"
	return draft
}
