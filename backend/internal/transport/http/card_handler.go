// Package httpapi 提供HTTP API路由和请求处理函数。
package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CardHandler 处理卡片和牌组相关的HTTP请求。
type CardHandler struct{}

// RegisterCardRoutes 注册卡片和牌组相关的HTTP路由。
func RegisterCardRoutes(r gin.IRoutes, h CardHandler) {
	r.POST("/decks/generate", h.Accepted)
	r.GET("/decks", h.List)
	r.POST("/decks", h.Created)
	r.PATCH("/decks/:deckId", h.OK)
	r.POST("/decks/merge", h.Created)
	r.POST("/decks/:deckId/restore", h.Created)
	r.GET("/cards", h.List)
	r.POST("/cards", h.Created)
	r.PATCH("/cards/:cardId", h.OK)
}
// OK 返回200成功响应。
func (CardHandler) OK(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"ok": true}) }
// List 返回空列表响应。
func (CardHandler) List(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"items": []any{}}) }
// Created 返回201创建成功响应。
func (CardHandler) Created(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"id": "created"}) }
// Accepted 返回202任务已接受响应。
func (CardHandler) Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"job": gin.H{"id": "job", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}
