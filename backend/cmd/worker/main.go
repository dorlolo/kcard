package main

import (
	"context"
	"encoding/json"
	"errors"
	"kcardDesgin/backend/internal/ai/agents/classifier"
	"kcardDesgin/backend/internal/ai/model"
	"log/slog"
	"os"
	"time"

	"github.com/redis/go-redis/v9"

	"kcardDesgin/backend/internal/ai"
	"kcardDesgin/backend/internal/app"
	"kcardDesgin/backend/internal/config"
	"kcardDesgin/backend/internal/jobs"
	"kcardDesgin/backend/internal/repository"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("configuration error", "error", err)
		os.Exit(1)
	}
	container, err := app.New(cfg)
	if err != nil {
		slog.Error("worker bootstrap failed", "error", err)
		os.Exit(1)
	}
	defer container.Close(context.Background())

	queue := jobs.NewQueue(container.Redis, jobs.DefaultQueueName)
	materialRepo := repository.NewMaterialRepository(container.DB)
	knowledgeRepo := repository.NewKnowledgeRepository(container.DB)
	aiClient, err := classifier.NewAgentWithClient(context.Background(), model.ModelConfig{Provider: cfg.AIProvider, APIKey: cfg.ArkAPIKey, ModelID: cfg.ArkModel, BaseURL: cfg.ArkBaseURL}, cfg.ArkAPIKey == "")
	if err != nil {
		slog.Error("ai client configuration failed", "error", err)
		os.Exit(1)
	}
	worker := jobs.MaterialAnalysisWorker{Workflow: ai.ClassificationWorkflow{Client: aiClient, DefaultPrompt: "请提取适合制作复习卡片的原子化知识点。"}, Materials: materialRepo, Knowledge: knowledgeRepo}

	slog.Info("worker ready", "queue", jobs.DefaultQueueName)
	for {
		if err := handleNext(context.Background(), queue, worker); err != nil {
			if errors.Is(err, redis.Nil) {
				continue
			}
			slog.Error("job handling failed", "error", err)
		}
	}
}

func handleNext(ctx context.Context, queue jobs.Queue, materialWorker jobs.MaterialAnalysisWorker) error {
	envelope, err := queue.Pop(ctx, 5*time.Second)
	if err != nil {
		return err
	}
	switch envelope.Job.Type {
	case jobs.JobTypeMaterialAnalysis:
		var payload jobs.MaterialAnalysisPayload
		if err := json.Unmarshal(envelope.Payload, &payload); err != nil {
			return err
		}
		_, err := materialWorker.Handle(ctx, payload)
		return err
	case "":
		return nil
	default:
		slog.Warn("unknown job type", "type", envelope.Job.Type)
		return nil
	}
}
