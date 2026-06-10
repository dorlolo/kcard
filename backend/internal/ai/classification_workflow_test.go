package ai

import (
	"context"
	"testing"
)

type fakeClient struct{ body []byte }

func (f fakeClient) GenerateStructured(ctx context.Context, req StructuredRequest) (StructuredResponse, error) {
	return StructuredResponse{JSON: f.body, ModelID: "test", StopReason: "end"}, nil
}

func TestClassificationWorkflowParsesStructuredOutput(t *testing.T) {
	workflow := ClassificationWorkflow{Client: fakeClient{body: []byte(`{"knowledgePoints":[{"content":"Cell membrane controls transport","summary":"Cell membrane","tags":["biology"],"confidence":0.9}]}`)}}
	out, err := workflow.Classify(context.Background(), "cell membrane", "classify")
	if err != nil {
		t.Fatal(err)
	}
	if len(out.KnowledgePoints) != 1 {
		t.Fatalf("points=%d", len(out.KnowledgePoints))
	}
	if out.KnowledgePoints[0].Content == "" {
		t.Fatal("missing content")
	}
}

func TestClassificationWorkflowRejectsEmptyMaterial(t *testing.T) {
	workflow := ClassificationWorkflow{Client: fakeClient{body: []byte(`{}`)}}
	_, err := workflow.Classify(context.Background(), " ", "")
	if err == nil {
		t.Fatal("expected error")
	}
}
