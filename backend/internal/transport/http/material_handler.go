package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kcardDesgin/backend/internal/domain"
	"kcardDesgin/backend/internal/service"
	"kcardDesgin/backend/internal/transport/http/middleware"
)

// MaterialHandler 处理资料导入相关的HTTP请求。
type MaterialHandler struct{ Service service.MaterialService }

type createMaterialRequest struct {
	SourceType      string   `json:"sourceType"`
	Title           string   `json:"title"`
	Text            string   `json:"text"`
	Tags            []string `json:"tags"`
	PromptText      string   `json:"promptText"`
	DuplicatePolicy string   `json:"duplicatePolicy"`
}

// RegisterMaterialRoutes 注册资料相关的HTTP路由。
func RegisterMaterialRoutes(r gin.IRoutes, handler MaterialHandler) {
	r.POST("/materials", handler.Create)
	r.GET("/materials/:materialId", handler.Get)
	r.POST("/materials/:materialId/reanalyze", handler.Reanalyze)
}

// Create 创建新的学习资料并触发分析任务。
func (h MaterialHandler) Create(c *gin.Context) {
	var req createMaterialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}
	workspaceID, _ := c.Get(middleware.WorkspaceIDKey)
	var tags []domain.Tag
	for _, tag := range req.Tags {
		tags = append(tags, domain.Tag{Name: tag})
	}
	result, err := h.Service.Create(c.Request.Context(), service.CreateMaterialInput{WorkspaceID: domain.ID(workspaceID.(string)), SourceType: domain.SourceType(req.SourceType), Title: req.Title, Text: req.Text, Tags: tags, PromptText: req.PromptText, DuplicatePolicy: req.DuplicatePolicy})
	if err != nil {
		Error(c, http.StatusBadRequest, "material_invalid", err.Error())
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"material": result.Material, "job": gin.H{"id": result.JobID, "jobType": "material_analysis", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}

// Get 获取指定资料的详细信息。
func (h MaterialHandler) Get(c *gin.Context) {
	material, err := h.Service.Get(c.Request.Context(), getWorkspaceID(c), domain.ID(c.Param("materialId")))
	if err != nil {
		Error(c, http.StatusNotFound, "material_not_found", err.Error())
		return
	}
	c.JSON(http.StatusOK, material)
}
// Reanalyze 触发资料重新分析任务。
func (h MaterialHandler) Reanalyze(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"job": gin.H{"id": c.Param("materialId") + ":reanalysis", "jobType": "material_analysis", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}
