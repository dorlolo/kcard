package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PortabilityHandler 处理导入导出相关的HTTP请求。
type PortabilityHandler struct{}

// RegisterPortabilityRoutes 注册导入导出相关的HTTP路由。
func RegisterPortabilityRoutes(r gin.IRoutes, h PortabilityHandler) {
	r.POST("/exports", h.Accepted)
	r.POST("/imports", h.Accepted)
}
// OK 返回200成功响应。
func (PortabilityHandler) OK(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"ok": true}) }
// List 返回空列表响应。
func (PortabilityHandler) List(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"items": []any{}}) }
// Created 返回201创建成功响应。
func (PortabilityHandler) Created(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"id": "created"}) }
// Accepted 返回202任务已接受响应。
func (PortabilityHandler) Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"job": gin.H{"id": "job", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}
