package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PromptHandler struct{}

func RegisterPromptRoutes(r gin.IRoutes, h PromptHandler) {
	r.GET("/prompts", h.List)
	r.POST("/prompts", h.Created)
	r.PATCH("/ai/drafts/:draftId", h.OK)
}
func (PromptHandler) OK(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"ok": true}) }
func (PromptHandler) List(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"items": []any{}}) }
func (PromptHandler) Created(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"id": "created"}) }
func (PromptHandler) Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"job": gin.H{"id": "job", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}
