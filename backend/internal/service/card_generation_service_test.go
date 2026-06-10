package service

import "testing"

func TestCardGenerationRequiresApprovedPoints(t *testing.T) {
	_, err := CardGenerationService{}.GenerateDrafts(nil)
	if err == nil {
		t.Fatal("expected error")
	}
}
