package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReviewHandler 处理复习计划相关的HTTP请求。
type ReviewHandler struct{}

// RegisterReviewRoutes 注册复习计划相关的HTTP路由。
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
// OK 返回200成功响应。
func (ReviewHandler) OK(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"ok": true}) }
// List 返回空列表响应。
func (ReviewHandler) List(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"items": []any{}}) }
// Created 返回201创建成功响应。
func (ReviewHandler) Created(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"id": "created"}) }
// Accepted 返回202任务已接受响应。
func (ReviewHandler) Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"job": gin.H{"id": "job", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}
