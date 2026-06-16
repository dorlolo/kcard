package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DashboardHandler 处理仪表盘相关的HTTP请求。
type DashboardHandler struct{}

// RegisterDashboardRoutes 注册仪表盘相关的HTTP路由。
func RegisterDashboardRoutes(r gin.IRoutes, h DashboardHandler) { r.GET("/dashboard", h.OK) }
// OK 返回200成功响应。
func (DashboardHandler) OK(c *gin.Context)                      { c.JSON(http.StatusOK, gin.H{"ok": true}) }
// List 返回空列表响应。
func (DashboardHandler) List(c *gin.Context)                    { c.JSON(http.StatusOK, gin.H{"items": []any{}}) }
// Created 返回201创建成功响应。
func (DashboardHandler) Created(c *gin.Context)                 { c.JSON(http.StatusCreated, gin.H{"id": "created"}) }
// Accepted 返回202任务已接受响应。
func (DashboardHandler) Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"job": gin.H{"id": "job", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}
