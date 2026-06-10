package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const WorkspaceIDKey = "workspaceID"

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
