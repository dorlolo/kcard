package service

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"

	"kcardDesgin/backend/internal/domain"
)

type MaterialStore interface {
	SaveMaterial(ctx context.Context, material domain.SourceMaterial) (domain.SourceMaterial, error)
	FindMaterialDigests(ctx context.Context, workspaceID domain.ID, digest string) ([]domain.SourceMaterial, error)
}

type JobEnqueuer interface {
	EnqueueMaterialAnalysis(ctx context.Context, materialID domain.ID, prompt string) (string, error)
}

type MaterialService struct {
	Store MaterialStore
	Jobs  JobEnqueuer
}

type CreateMaterialInput struct {
	WorkspaceID     domain.ID
	SourceType      domain.SourceType
	Title           string
	Text            string
	URL             string
	Tags            []domain.Tag
	PromptText      string
	DuplicatePolicy string
}
type CreateMaterialResult struct {
	Material         domain.SourceMaterial
	JobID            string
	DuplicateWarning bool
}

func (s MaterialService) Create(ctx context.Context, input CreateMaterialInput) (CreateMaterialResult, error) {
	if input.WorkspaceID == "" {
		return CreateMaterialResult{}, errors.New("workspace id is required")
	}
	if input.SourceType != domain.SourceText {
		return CreateMaterialResult{}, errors.New("only text material intake is implemented in the MVP skeleton")
	}
	material, err := domain.NewTextMaterial(input.WorkspaceID, input.Title, input.Text, input.Tags)
	if err != nil {
		return CreateMaterialResult{}, err
	}
	material.ID = domain.ID(uuid.NewString())
	matches, err := s.Store.FindMaterialDigests(ctx, input.WorkspaceID, material.ContentDigest)
	if err != nil {
		return CreateMaterialResult{}, err
	}
	duplicate := len(matches) > 0
	if duplicate && strings.EqualFold(input.DuplicatePolicy, "warn") {
		material = material.Queue(domain.DuplicatePossible)
	} else {
		material = material.Queue(domain.DuplicateUnique)
	}
	saved, err := s.Store.SaveMaterial(ctx, material)
	if err != nil {
		return CreateMaterialResult{}, err
	}
	jobID := ""
	if s.Jobs != nil {
		jobID, err = s.Jobs.EnqueueMaterialAnalysis(ctx, saved.ID, input.PromptText)
		if err != nil {
			return CreateMaterialResult{}, err
		}
	}
	return CreateMaterialResult{Material: saved, JobID: jobID, DuplicateWarning: duplicate}, nil
}

type MemoryMaterialStore struct{ Materials []domain.SourceMaterial }

func (m *MemoryMaterialStore) SaveMaterial(ctx context.Context, material domain.SourceMaterial) (domain.SourceMaterial, error) {
	m.Materials = append(m.Materials, material)
	return material, nil
}
func (m *MemoryMaterialStore) FindMaterialDigests(ctx context.Context, workspaceID domain.ID, digest string) ([]domain.SourceMaterial, error) {
	var out []domain.SourceMaterial
	for _, item := range m.Materials {
		if item.LearnerWorkspaceID == workspaceID && item.ContentDigest == digest {
			out = append(out, item)
		}
	}
	return out, nil
}

type MemoryJobEnqueuer struct{ Jobs []domain.ID }

func (m *MemoryJobEnqueuer) EnqueueMaterialAnalysis(ctx context.Context, materialID domain.ID, prompt string) (string, error) {
	m.Jobs = append(m.Jobs, materialID)
	return string(materialID) + ":analysis", nil
}
