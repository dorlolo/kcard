package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CardHandler struct{}

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
func (CardHandler) OK(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"ok": true}) }
func (CardHandler) List(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"items": []any{}}) }
func (CardHandler) Created(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"id": "created"}) }
func (CardHandler) Accepted(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"job": gin.H{"id": "job", "status": "queued", "progressPercent": 0, "currentStep": "queued"}})
}
