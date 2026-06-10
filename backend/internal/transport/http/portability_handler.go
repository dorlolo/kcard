package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PortabilityHandler struct{}

func RegisterPortabilityRoutes(r gin.IRoutes, h PortabilityHandler) {
	r.POST("/exports", h.Accepted)
	r.POST("/imports", h.Accepted)
}
func (PortabilityHandler) OK(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"ok": true}) }
func (PortabilityHandler) List(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"items": []any{}}) }
func (PortabilityHandler) Created(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"id": "created"}) }
func (PortabilityHandler) Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"job": gin.H{"id": "job", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}
