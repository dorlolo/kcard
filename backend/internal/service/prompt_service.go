// Package service 提供领域服务层的实现，包括资料、卡片、知识图谱等核心业务逻辑。
package service

import (
	"errors"
	"kcardDesgin/backend/internal/domain"
)

// PromptService 处理提示词预设的保存和查询操作。
type PromptService struct{ Presets []domain.PromptPreset }

// Save 保存提示词预设，校验名称和内容不为空。
func (s *PromptService) Save(p domain.PromptPreset) (domain.PromptPreset, error) {
	if p.Name == "" || p.PromptText == "" {
		return domain.PromptPreset{}, errors.New("name and prompt text are required")
	}
	if p.VersionNumber == 0 {
		p.VersionNumber = 1
	}
	s.Presets = append(s.Presets, p)
	return p, nil
}
// List 根据用途筛选并返回提示词预设列表。
func (s PromptService) List(purpose string) []domain.PromptPreset {
	out := []domain.PromptPreset{}
	for _, p := range s.Presets {
		if purpose == "" || p.Purpose == purpose {
			out = append(out, p)
		}
	}
	return out
}
