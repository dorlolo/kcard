package service

import (
	"context"
	"testing"

	"kcardDesgin/backend/internal/domain"
)

func TestMaterialCreateQueuesTextMaterial(t *testing.T) {
	store := &MemoryMaterialStore{}
	jobs := &MemoryJobEnqueuer{}
	svc := MaterialService{Store: store, Jobs: jobs}
	result, err := svc.Create(context.Background(), CreateMaterialInput{WorkspaceID: "workspace-1", SourceType: domain.SourceText, Title: "Biology", Text: "Cells contain organelles.", DuplicatePolicy: "warn"})
	if err != nil {
		t.Fatalf("Create returned error: %v", err)
	}
	if result.Material.ProcessingStatus != domain.MaterialQueued {
		t.Fatalf("status=%s", result.Material.ProcessingStatus)
	}
	if result.Material.ContentDigest == "" {
		t.Fatal("content digest was not set")
	}
	if result.JobID == "" {
		t.Fatal("analysis job was not enqueued")
	}
}

func TestMaterialCreateFlagsDuplicate(t *testing.T) {
	store := &MemoryMaterialStore{}
	svc := MaterialService{Store: store, Jobs: &MemoryJobEnqueuer{}}
	input := CreateMaterialInput{WorkspaceID: "workspace-1", SourceType: domain.SourceText, Title: "Repeat", Text: "same text", DuplicatePolicy: "warn"}
	if _, err := svc.Create(context.Background(), input); err != nil {
		t.Fatal(err)
	}
	result, err := svc.Create(context.Background(), input)
	if err != nil {
		t.Fatal(err)
	}
	if !result.DuplicateWarning {
		t.Fatal("expected duplicate warning")
	}
	if result.Material.DuplicateStatus != domain.DuplicatePossible {
		t.Fatalf("duplicate status=%s", result.Material.DuplicateStatus)
	}
}

func TestMaterialCreateRejectsEmptyText(t *testing.T) {
	svc := MaterialService{Store: &MemoryMaterialStore{}}
	_, err := svc.Create(context.Background(), CreateMaterialInput{WorkspaceID: "workspace-1", SourceType: domain.SourceText, Text: "   "})
	if err == nil {
		t.Fatal("expected error")
	}
}
