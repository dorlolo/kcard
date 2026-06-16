// Package service 提供领域服务层的实现，包括资料、卡片、知识图谱等核心业务逻辑。
package service

import (
	"context"
	"errors"
	"sort"
	"strings"
	"time"

	"kcardDesgin/backend/internal/domain"
)

// KnowledgeFilter 是 domain.KnowledgeFilter 的类型别名，表示知识筛选条件。
type KnowledgeFilter = domain.KnowledgeFilter

// KnowledgeStore 定义知识存储的持久化接口。
type KnowledgeStore interface {
	Search(ctx context.Context, filter domain.KnowledgeFilter) ([]domain.KnowledgePoint, error)
	Get(ctx context.Context, workspaceID domain.ID, id domain.ID) (domain.KnowledgePoint, error)
	Create(ctx context.Context, point domain.KnowledgePoint) (domain.KnowledgePoint, error)
	Save(ctx context.Context, point domain.KnowledgePoint) (domain.KnowledgePoint, error)
	UpdateStatus(ctx context.Context, workspaceID domain.ID, id domain.ID, status domain.ApprovalStatus, notes string, now time.Time) (domain.KnowledgePoint, error)
}

// KnowledgeService 处理知识点的搜索、拆分、合并和状态更新操作。
type KnowledgeService struct {
	Store KnowledgeStore
}

// NewKnowledgeService 创建基于内存存储的 KnowledgeService 实例。
func NewKnowledgeService(points []domain.KnowledgePoint) *KnowledgeService {
	return &KnowledgeService{Store: NewMemoryKnowledgeStore(points)}
}

// NewPersistentKnowledgeService 创建基于持久化存储的 KnowledgeService 实例。
func NewPersistentKnowledgeService(store KnowledgeStore) *KnowledgeService {
	return &KnowledgeService{Store: store}
}

// Search 根据筛选条件搜索知识点，忽略错误直接返回结果。
func (s *KnowledgeService) Search(ctx context.Context, filter domain.KnowledgeFilter) []domain.KnowledgePoint {
	points, err := s.Store.Search(ctx, filter)
	if err != nil {
		return nil
	}
	return points
}

// SearchWithError 根据筛选条件搜索知识点，同时返回错误信息。
func (s *KnowledgeService) SearchWithError(ctx context.Context, filter domain.KnowledgeFilter) ([]domain.KnowledgePoint, error) {
	return s.Store.Search(ctx, filter)
}

// UpdateStatus 更新知识点的审批状态和备注信息。
func (s *KnowledgeService) UpdateStatus(ctx context.Context, workspaceID domain.ID, id domain.ID, status domain.ApprovalStatus, notes string, now time.Time) (domain.KnowledgePoint, error) {
	return s.Store.UpdateStatus(ctx, workspaceID, id, status, notes, now)
}

// Split 将指定知识点拆分为多个新知识点并标记原知识点待审核。
func (s *KnowledgeService) Split(ctx context.Context, workspaceID domain.ID, id domain.ID, contents []string) ([]domain.KnowledgePoint, error) {
	source, err := s.Store.Get(ctx, workspaceID, id)
	if err != nil {
		return nil, errors.New("knowledge point not found")
	}
	if len(contents) < 2 {
		return nil, errors.New("split requires at least two points")
	}
	var out []domain.KnowledgePoint
	for idx, content := range contents {
		content = strings.TrimSpace(content)
		if content == "" {
			return nil, errors.New("split content cannot be empty")
		}
		point := source
		point.ID = domain.ID(string(id) + ":split:" + string(rune('a'+idx)))
		point.Content = content
		point.Summary = summarize(content)
		point.GraphLabel = point.Summary
		point.ApprovalStatus = domain.KnowledgeDraft
		point.ApprovedAt = nil
		point.RejectedAt = nil
		created, err := s.Store.Create(ctx, point)
		if err != nil {
			return nil, err
		}
		out = append(out, created)
	}
	source.ApprovalStatus = domain.KnowledgeNeedsReview
	source.Notes = appendNote(source.Notes, "已拆分为多个知识点")
	if _, err := s.Store.Save(ctx, source); err != nil {
		return nil, err
	}
	return out, nil
}

// Merge 将多个知识点合并为一个新知识点并标记原知识点待审核。
func (s *KnowledgeService) Merge(ctx context.Context, workspaceID domain.ID, ids []domain.ID, content string) (domain.KnowledgePoint, error) {
	if len(ids) < 2 {
		return domain.KnowledgePoint{}, errors.New("merge requires at least two points")
	}
	content = strings.TrimSpace(content)
	if content == "" {
		return domain.KnowledgePoint{}, errors.New("merged content cannot be empty")
	}
	first, err := s.Store.Get(ctx, workspaceID, ids[0])
	if err != nil {
		return domain.KnowledgePoint{}, errors.New("knowledge point not found")
	}
	for _, id := range ids[1:] {
		if _, err := s.Store.Get(ctx, workspaceID, id); err != nil {
			return domain.KnowledgePoint{}, errors.New("knowledge point not found")
		}
	}
	merged := first
	merged.ID = domain.ID(string(ids[0]) + ":merged")
	merged.Content = content
	merged.Summary = summarize(content)
	merged.GraphLabel = merged.Summary
	merged.ApprovalStatus = domain.KnowledgeDraft
	merged.ApprovedAt = nil
	merged.RejectedAt = nil
	created, err := s.Store.Create(ctx, merged)
	if err != nil {
		return domain.KnowledgePoint{}, err
	}
	for _, id := range ids {
		point, err := s.Store.Get(ctx, workspaceID, id)
		if err != nil {
			return domain.KnowledgePoint{}, err
		}
		point.ApprovalStatus = domain.KnowledgeNeedsReview
		point.Notes = appendNote(point.Notes, "已合并到 "+string(created.ID))
		if _, err := s.Store.Save(ctx, point); err != nil {
			return domain.KnowledgePoint{}, err
		}
	}
	return created, nil
}

// MemoryKnowledgeStore 实现 KnowledgeStore 接口的内存存储，用于测试和开发。
type MemoryKnowledgeStore struct {
	Points map[domain.ID]domain.KnowledgePoint
}

// NewMemoryKnowledgeStore 创建基于初始知识点列表的内存知识存储实例。
func NewMemoryKnowledgeStore(points []domain.KnowledgePoint) *MemoryKnowledgeStore {
	store := &MemoryKnowledgeStore{Points: map[domain.ID]domain.KnowledgePoint{}}
	for _, point := range points {
		store.Points[point.ID] = point
	}
	return store
}

// Search 根据筛选条件从内存中搜索知识点，支持文本、标签和状态过滤。
func (s *MemoryKnowledgeStore) Search(ctx context.Context, filter domain.KnowledgeFilter) ([]domain.KnowledgePoint, error) {
	query := strings.ToLower(strings.TrimSpace(filter.Query))
	tag := strings.ToLower(strings.TrimSpace(filter.Tag))
	out := make([]domain.KnowledgePoint, 0, len(s.Points))
	for _, point := range s.Points {
		if point.LearnerWorkspaceID != filter.WorkspaceID {
			continue
		}
		if filter.ApprovalStatus != "" && point.ApprovalStatus != filter.ApprovalStatus {
			continue
		}
		if !filter.IncludeRejected && point.ApprovalStatus == domain.KnowledgeRejected {
			continue
		}
		if query != "" && !strings.Contains(strings.ToLower(point.Content+" "+point.Summary+" "+point.Notes), query) {
			continue
		}
		if tag != "" && !pointHasTag(point, tag) {
			continue
		}
		out = append(out, point)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out, nil
}

// Get 根据工作区和 ID 从内存中获取单个知识点。
func (s *MemoryKnowledgeStore) Get(ctx context.Context, workspaceID domain.ID, id domain.ID) (domain.KnowledgePoint, error) {
	point, ok := s.Points[id]
	if !ok || point.LearnerWorkspaceID != workspaceID {
		return domain.KnowledgePoint{}, errors.New("knowledge point not found")
	}
	return point, nil
}

// Create 往内存存储中创建新的知识点记录。
func (s *MemoryKnowledgeStore) Create(ctx context.Context, point domain.KnowledgePoint) (domain.KnowledgePoint, error) {
	if s.Points == nil {
		s.Points = map[domain.ID]domain.KnowledgePoint{}
	}
	s.Points[point.ID] = point
	return point, nil
}

// Save 保存知识点到内存存储，若已存在则覆盖。
func (s *MemoryKnowledgeStore) Save(ctx context.Context, point domain.KnowledgePoint) (domain.KnowledgePoint, error) {
	return s.Create(ctx, point)
}

// UpdateStatus 更新内存中知识点的审批状态、备注和审批时间。
func (s *MemoryKnowledgeStore) UpdateStatus(ctx context.Context, workspaceID domain.ID, id domain.ID, status domain.ApprovalStatus, notes string, now time.Time) (domain.KnowledgePoint, error) {
	point, err := s.Get(ctx, workspaceID, id)
	if err != nil {
		return domain.KnowledgePoint{}, err
	}
	point.ApprovalStatus = status
	if notes != "" {
		point.Notes = notes
	}
	switch status {
	case domain.KnowledgeApproved:
		point = point.Approve(now)
	case domain.KnowledgeRejected:
		point = point.Reject(now)
	case domain.KnowledgeNeedsReview:
		point = point.MarkNeedsReview()
	}
	s.Points[id] = point
	return point, nil
}

func pointHasTag(point domain.KnowledgePoint, tag string) bool {
	for _, item := range point.Tags {
		if strings.ToLower(item.Name) == tag || strings.ToLower(string(item.ID)) == tag {
			return true
		}
	}
	return false
}

func summarize(content string) string {
	content = strings.TrimSpace(content)
	if len([]rune(content)) <= 24 {
		return content
	}
	return string([]rune(content)[:24]) + "…"
}

func appendNote(existing, note string) string {
	if strings.TrimSpace(existing) == "" {
		return note
	}
	return existing + "；" + note
}
