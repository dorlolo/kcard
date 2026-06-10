package service

import (
	"kcardDesgin/backend/internal/domain"
	"testing"
)

func TestPromptRequiresText(t *testing.T) {
	_, err := (&PromptService{}).Save(domain.PromptPreset{Name: "n"})
	if err == nil {
		t.Fatal("expected error")
	}
}
