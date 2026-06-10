package service

import (
	"errors"
	"kcardDesgin/backend/internal/domain"
)

type PromptService struct{ Presets []domain.PromptPreset }

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
func (s PromptService) List(purpose string) []domain.PromptPreset {
	out := []domain.PromptPreset{}
	for _, p := range s.Presets {
		if purpose == "" || p.Purpose == purpose {
			out = append(out, p)
		}
	}
	return out
}
