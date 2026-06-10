package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kcardDesgin/backend/internal/domain"
	"kcardDesgin/backend/internal/service"
	"kcardDesgin/backend/internal/transport/http/middleware"
)

type KnowledgeHandler struct {
	Service *service.KnowledgeService
	Graph   *service.KnowledgeGraphService
}

func RegisterKnowledgeRoutes(r gin.IRoutes, h KnowledgeHandler) {
	r.GET("/knowledge-points", h.List)
	r.POST("/knowledge-points/:knowledgePointId/split", h.Split)
	r.POST("/knowledge-points/merge", h.Merge)
	r.GET("/knowledge-graph", h.GraphView)
	r.POST("/knowledge-relationships", h.CreateRelationship)
	r.PATCH("/knowledge-relationships/:relationshipId", h.UpdateRelationship)
}

func (h KnowledgeHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"items": []any{}, "meta": gin.H{"page": 1, "pageSize": 25, "total": 0}})
}
func (h KnowledgeHandler) Split(c *gin.Context) { c.JSON(http.StatusCreated, gin.H{"items": []any{}}) }
func (h KnowledgeHandler) Merge(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"id": "merged", "approvalStatus": "draft"})
}
func (h KnowledgeHandler) GraphView(c *gin.Context) {
	workspaceID, _ := c.Get(middleware.WorkspaceIDKey)
	graph, _ := h.Graph.Graph(c.Request.Context(), domain.ID(workspaceID.(string)), domain.ID(c.Query("focusKnowledgePointId")), 1, c.Query("includeArchived") == "true")
	c.JSON(http.StatusOK, graph)
}
func (h KnowledgeHandler) CreateRelationship(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"id": "relationship", "relationshipType": "related"})
}
func (h KnowledgeHandler) UpdateRelationship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"id": c.Param("relationshipId"), "archived": false})
}
