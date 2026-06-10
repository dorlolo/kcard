package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReviewHandler struct{}

func RegisterReviewRoutes(r gin.IRoutes, h ReviewHandler) {
	r.POST("/review/sessions", h.Created)
	r.POST("/review/sessions/:sessionId/answers", h.Created)
	r.PATCH("/review/sessions/:sessionId", h.OK)
	r.GET("/plans", h.List)
	r.POST("/plans", h.Created)
	r.POST("/plans/generate", h.Accepted)
	r.PATCH("/plans/:planId", h.OK)
	r.POST("/plans/:planId/optimize", h.Accepted)
	r.GET("/plans/:planId/revisions", h.List)
	r.POST("/plans/:planId/revisions/:revisionId/restore", h.OK)
	r.GET("/statistics", h.OK)
}
func (ReviewHandler) OK(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"ok": true}) }
func (ReviewHandler) List(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"items": []any{}}) }
func (ReviewHandler) Created(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"id": "created"}) }
func (ReviewHandler) Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"job": gin.H{"id": "job", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}
