package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PromptHandler 处理提示词相关的HTTP请求。
type PromptHandler struct{}

// RegisterPromptRoutes 注册提示词相关的HTTP路由。
func RegisterPromptRoutes(r gin.IRoutes, h PromptHandler) {
	r.GET("/prompts", h.List)
	r.POST("/prompts", h.Created)
	r.PATCH("/ai/drafts/:draftId", h.OK)
}
// OK 返回200成功响应。
func (PromptHandler) OK(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"ok": true}) }
// List 返回空列表响应。
func (PromptHandler) List(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"items": []any{}}) }
// Created 返回201创建成功响应。
func (PromptHandler) Created(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"id": "created"}) }
// Accepted 返回202任务已接受响应。
func (PromptHandler) Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"job": gin.H{"id": "job", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}
