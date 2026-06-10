package contract

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"kcardDesgin/backend/internal/app"
	httpapi "kcardDesgin/backend/internal/transport/http"
)

func TestKnowledgeGraphContract(t *testing.T) {
	router := httpapi.NewRouter(&app.Container{})
	req := httptest.NewRequest(http.MethodGet, "/api/v1/knowledge-graph?depth=1", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", res.Code, res.Body.String())
	}
}

func TestKnowledgeListContract(t *testing.T) {
	router := httpapi.NewRouter(&app.Container{})
	req := httptest.NewRequest(http.MethodGet, "/api/v1/knowledge-points", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Fatalf("status=%d", res.Code)
	}
}
