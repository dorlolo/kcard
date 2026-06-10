package httpapi

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"kcardDesgin/backend/internal/domain"
	"kcardDesgin/backend/internal/service"
	"kcardDesgin/backend/internal/transport/http/middleware"
)

type KnowledgeHandler struct {
	Service *service.KnowledgeService
	Graph   *service.KnowledgeGraphService
}

type splitKnowledgeRequest struct {
	Items []struct {
		Content string `json:"content"`
	} `json:"items"`
}

type mergeKnowledgeRequest struct {
	KnowledgePointIDs []string `json:"knowledgePointIds"`
	Content           string   `json:"content"`
}

type createRelationshipRequest struct {
	SourceKnowledgePointID string  `json:"sourceKnowledgePointId"`
	TargetKnowledgePointID string  `json:"targetKnowledgePointId"`
	RelationshipType       string  `json:"relationshipType"`
	Label                  string  `json:"label"`
	Weight                 float64 `json:"weight"`
}

type updateRelationshipRequest struct {
	Archived bool `json:"archived"`
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
	workspaceID := workspaceID(c)
	points := h.Service.Search(c.Request.Context(), service.KnowledgeFilter{
		WorkspaceID:     workspaceID,
		Query:           c.Query("q"),
		ApprovalStatus:  domain.ApprovalStatus(c.Query("approvalStatus")),
		Tag:             c.Query("tag"),
		IncludeRejected: c.Query("includeRejected") == "true",
	})
	c.JSON(http.StatusOK, gin.H{"items": points, "meta": gin.H{"page": 1, "pageSize": len(points), "total": len(points)}})
}

func (h KnowledgeHandler) Split(c *gin.Context) {
	var req splitKnowledgeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}
	contents := make([]string, 0, len(req.Items))
	for _, item := range req.Items {
		contents = append(contents, item.Content)
	}
	points, err := h.Service.Split(c.Request.Context(), workspaceID(c), domain.ID(c.Param("knowledgePointId")), contents)
	if err != nil {
		Error(c, http.StatusBadRequest, "split_failed", err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"items": points})
}

func (h KnowledgeHandler) Merge(c *gin.Context) {
	var req mergeKnowledgeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}
	ids := make([]domain.ID, 0, len(req.KnowledgePointIDs))
	for _, id := range req.KnowledgePointIDs {
		ids = append(ids, domain.ID(id))
	}
	point, err := h.Service.Merge(c.Request.Context(), workspaceID(c), ids, req.Content)
	if err != nil {
		Error(c, http.StatusBadRequest, "merge_failed", err.Error())
		return
	}
	c.JSON(http.StatusCreated, point)
}

func (h KnowledgeHandler) GraphView(c *gin.Context) {
	depth, _ := strconv.Atoi(c.DefaultQuery("depth", "1"))
	relationshipTypes := parseRelationshipTypes(c.QueryArray("relationshipType"))
	graph, err := h.Graph.Graph(c.Request.Context(), service.GraphQuery{
		WorkspaceID:       workspaceID(c),
		FocusID:           domain.ID(c.Query("focusKnowledgePointId")),
		Depth:             depth,
		Query:             c.Query("q"),
		ApprovalStatus:    domain.ApprovalStatus(c.Query("approvalStatus")),
		RelationshipTypes: relationshipTypes,
		IncludeArchived:   c.Query("includeArchived") == "true",
		IncludeRejected:   c.Query("includeRejected") == "true",
		MaxNodes:          250,
		MaxEdges:          1000,
	})
	if err != nil {
		Error(c, http.StatusBadRequest, "graph_failed", err.Error())
		return
	}
	c.JSON(http.StatusOK, graph)
}

func (h KnowledgeHandler) CreateRelationship(c *gin.Context) {
	var req createRelationshipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, http.StatusBadRequest, "bad_request", err.Error())
		return
	}
	rel, err := h.Graph.AddRelationship(c.Request.Context(), service.KnowledgeRelationship{
		WorkspaceID: workspaceID(c),
		SourceID:    domain.ID(req.SourceKnowledgePointID),
		TargetID:    domain.ID(req.TargetKnowledgePointID),
		Type:        service.RelationshipType(req.RelationshipType),
		Label:       req.Label,
		Weight:      req.Weight,
		SourceType:  "learner_created",
	})
	if err != nil {
		Error(c, http.StatusBadRequest, "relationship_failed", err.Error())
		return
	}
	c.JSON(http.StatusCreated, relationshipResponse(rel))
}

func (h KnowledgeHandler) UpdateRelationship(c *gin.Context) {
	var req updateRelationshipRequest
	_ = c.ShouldBindJSON(&req)
	if !req.Archived {
		c.JSON(http.StatusOK, gin.H{"id": c.Param("relationshipId"), "archived": false})
		return
	}
	rel, err := h.Graph.ArchiveRelationship(c.Request.Context(), domain.ID(c.Param("relationshipId")))
	if err != nil {
		Error(c, http.StatusNotFound, "relationship_not_found", err.Error())
		return
	}
	c.JSON(http.StatusOK, relationshipResponse(rel))
}

func workspaceID(c *gin.Context) domain.ID {
	value, _ := c.Get(middleware.WorkspaceIDKey)
	if id, ok := value.(string); ok && id != "" {
		return domain.ID(id)
	}
	return "00000000-0000-0000-0000-000000000001"
}

func parseRelationshipTypes(values []string) []service.RelationshipType {
	var out []service.RelationshipType
	for _, value := range values {
		for _, part := range strings.Split(value, ",") {
			part = strings.TrimSpace(part)
			if part != "" {
				out = append(out, service.RelationshipType(part))
			}
		}
	}
	return out
}

func relationshipResponse(rel service.KnowledgeRelationship) gin.H {
	return gin.H{
		"id":               rel.ID,
		"sourceId":         rel.SourceID,
		"targetId":         rel.TargetID,
		"relationshipType": rel.Type,
		"label":            rel.Label,
		"weight":           rel.Weight,
		"sourceType":       rel.SourceType,
		"archived":         rel.Archived,
		"updatedAt":        time.Now().UTC(),
	}
}
