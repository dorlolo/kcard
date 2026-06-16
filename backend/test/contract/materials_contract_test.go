package contract

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"kcardDesgin/backend/internal/app"
	httpapi "kcardDesgin/backend/internal/transport/http"
)

func TestCreateMaterialContract(t *testing.T) {
	router := httpapi.NewRouter(&app.Container{})
	body := bytes.NewBufferString(`{"sourceType":"text","title":"Biology","text":"Cells have nuclei","tags":["bio"],"duplicatePolicy":"warn"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/materials", body)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusAccepted {
		t.Fatalf("status=%d body=%s", res.Code, res.Body.String())
	}
}

func TestKnowledgeUpdateContract(t *testing.T) {
	router := httpapi.NewRouter(&app.Container{})
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/knowledge-points/kp-cell", bytes.NewBufferString(`{"approvalStatus":"approved"}`))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Fatalf("status=%d", res.Code)
	}
}
