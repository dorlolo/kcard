package contract

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"kcardDesgin/backend/internal/app"
	httpapi "kcardDesgin/backend/internal/transport/http"
)

func TestKnowledgeGraphContract(t *testing.T) {
	router := httpapi.NewRouter(&app.Container{})
	req := httptest.NewRequest(http.MethodGet, "/api/v1/knowledge-graph?depth=1&relationshipType=related", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", res.Code, res.Body.String())
	}
	if !strings.Contains(res.Body.String(), "nodes") || !strings.Contains(res.Body.String(), "edges") {
		t.Fatalf("graph response missing nodes/edges: %s", res.Body.String())
	}
}

func TestKnowledgeListContract(t *testing.T) {
	router := httpapi.NewRouter(&app.Container{})
	req := httptest.NewRequest(http.MethodGet, "/api/v1/knowledge-points?q=细胞&includeRejected=true", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Fatalf("status=%d", res.Code)
	}
	if !strings.Contains(res.Body.String(), "items") {
		t.Fatalf("list response missing items: %s", res.Body.String())
	}
}

func TestKnowledgeSplitContract(t *testing.T) {
	router := httpapi.NewRouter(&app.Container{})
	body := bytes.NewBufferString(`{"items":[{"content":"细胞膜控制进出"},{"content":"细胞膜维持稳定"}]}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/knowledge-points/kp-membrane/split", body)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusCreated {
		t.Fatalf("status=%d body=%s", res.Code, res.Body.String())
	}
}

func TestKnowledgeMergeContract(t *testing.T) {
	router := httpapi.NewRouter(&app.Container{})
	body := bytes.NewBufferString(`{"knowledgePointIds":["kp-cell","kp-membrane"],"content":"细胞及细胞膜共同构成基础结构理解"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/knowledge-points/merge", body)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusCreated {
		t.Fatalf("status=%d body=%s", res.Code, res.Body.String())
	}
}

func TestKnowledgeRelationshipContract(t *testing.T) {
	router := httpapi.NewRouter(&app.Container{})
	body := bytes.NewBufferString(`{"sourceKnowledgePointId":"kp-cell","targetKnowledgePointId":"kp-nucleus","relationshipType":"supports","label":"帮助理解"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/knowledge-relationships", body)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusCreated {
		t.Fatalf("status=%d body=%s", res.Code, res.Body.String())
	}
}
