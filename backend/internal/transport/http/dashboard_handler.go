package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DashboardHandler struct{}

func RegisterDashboardRoutes(r gin.IRoutes, h DashboardHandler) { r.GET("/dashboard", h.OK) }
func (DashboardHandler) OK(c *gin.Context)                      { c.JSON(http.StatusOK, gin.H{"ok": true}) }
func (DashboardHandler) List(c *gin.Context)                    { c.JSON(http.StatusOK, gin.H{"items": []any{}}) }
func (DashboardHandler) Created(c *gin.Context)                 { c.JSON(http.StatusCreated, gin.H{"id": "created"}) }
func (DashboardHandler) Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"job": gin.H{"id": "job", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}
