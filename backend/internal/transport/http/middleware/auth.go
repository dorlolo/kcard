// Package middleware 提供HTTP中间件。
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WorkspaceIDKey 常量定义工作空间ID在上下文中的键名。
const WorkspaceIDKey = "workspaceID"

// PlaceholderAuth 返回一个占位认证中间件，从请求头提取工作空间ID并注入上下文。
func PlaceholderAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		workspaceID := c.GetHeader("X-Workspace-ID")
		if workspaceID == "" {
			workspaceID = "00000000-0000-0000-0000-000000000001"
		}
		c.Set(WorkspaceIDKey, workspaceID)
		if c.GetHeader("X-Deny-Access") == "true" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": "unauthorized", "message": "workspace access denied"}})
			return
		}
		c.Next()
	}
}
