package httpapi

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"kcardDesgin/backend/internal/app"
	"kcardDesgin/backend/internal/domain"
	"kcardDesgin/backend/internal/jobs"
	"kcardDesgin/backend/internal/repository"
	"kcardDesgin/backend/internal/service"
	"kcardDesgin/backend/internal/transport/http/middleware"
)

// NewRouter 创建并配置Gin引擎，注册所有API路由和中间件。
func NewRouter(container *app.Container) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), requestID(), cors(container.Config.AllowedFrontendOrigin), middleware.PlaceholderAuth())
	r.GET("/api/v1/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().UTC()}) })
	api := r.Group("/api/v1")
	RegisterMaterialRoutes(api, newMaterialHandler(container))
	RegisterKnowledgeRoutes(api, newKnowledgeHandler(container))
	RegisterCardRoutes(api, CardHandler{})
	RegisterReviewRoutes(api, ReviewHandler{})
	RegisterPromptRoutes(api, PromptHandler{})
	RegisterDashboardRoutes(api, DashboardHandler{})
	RegisterPortabilityRoutes(api, PortabilityHandler{})
	r.NoRoute(func(c *gin.Context) { Error(c, http.StatusNotFound, "not_found", "route not found") })
	return r
}

func newMaterialHandler(container *app.Container) MaterialHandler {
	if container != nil && container.DB != nil && container.Redis != nil {
		materialRepo := repository.NewMaterialRepository(container.DB)
		queue := jobs.NewQueue(container.Redis, jobs.DefaultQueueName)
		return MaterialHandler{Service: service.MaterialService{Store: materialRepo, Jobs: jobs.NewMaterialAnalysisEnqueuer(queue)}}
	}
	return MaterialHandler{Service: service.MaterialService{Store: &service.MemoryMaterialStore{}, Jobs: &service.MemoryJobEnqueuer{}}}
}

func newKnowledgeHandler(container *app.Container) KnowledgeHandler {
	if container != nil && container.DB != nil {
		knowledgeRepo := repository.NewKnowledgeRepository(container.DB)
		graphRepo := repository.NewKnowledgeGraphRepository(container.DB)
		return KnowledgeHandler{Service: service.NewPersistentKnowledgeService(knowledgeRepo), Graph: service.NewKnowledgeGraphService(knowledgeRepo, graphRepo)}
	}

	workspaceID := domain.ID("00000000-0000-0000-0000-000000000001")
	points := []domain.KnowledgePoint{
		{ID: "kp-cell", LearnerWorkspaceID: workspaceID, Content: "细胞是生命活动的基本结构和功能单位。", Summary: "细胞是生命基本单位", ApprovalStatus: domain.KnowledgeApproved, CreationSource: domain.CreationAIGenerated, GraphLabel: "细胞", Tags: []domain.Tag{{ID: "tag-biology", Name: "生物"}}},
		{ID: "kp-membrane", LearnerWorkspaceID: workspaceID, Content: "细胞膜控制物质进出，并维持细胞内环境稳定。", Summary: "细胞膜控制物质进出", ApprovalStatus: domain.KnowledgeDraft, CreationSource: domain.CreationAIGenerated, GraphLabel: "细胞膜", Tags: []domain.Tag{{ID: "tag-biology", Name: "生物"}}},
		{ID: "kp-nucleus", LearnerWorkspaceID: workspaceID, Content: "细胞核储存遗传信息，并调控细胞生命活动。", Summary: "细胞核储存遗传信息", ApprovalStatus: domain.KnowledgeNeedsReview, CreationSource: domain.CreationAIGenerated, GraphLabel: "细胞核", Tags: []domain.Tag{{ID: "tag-biology", Name: "生物"}}},
	}
	relationships := []domain.KnowledgeRelationship{
		{ID: "rel-cell-membrane", WorkspaceID: workspaceID, SourceID: "kp-cell", TargetID: "kp-membrane", Type: domain.RelationshipRelated, Label: "组成结构", Weight: 1, SourceType: "system_derived"},
		{ID: "rel-cell-nucleus", WorkspaceID: workspaceID, SourceID: "kp-cell", TargetID: "kp-nucleus", Type: domain.RelationshipRelated, Label: "组成结构", Weight: 1, SourceType: "system_derived"},
		{ID: "rel-membrane-nucleus", WorkspaceID: workspaceID, SourceID: "kp-membrane", TargetID: "kp-nucleus", Type: domain.RelationshipPrerequisite, Label: "理解细胞结构", Weight: .8, SourceType: "system_derived"},
	}
	pointStore := service.NewMemoryKnowledgeStore(points)
	relationshipStore := service.NewMemoryKnowledgeRelationshipStore(relationships)
	return KnowledgeHandler{Service: service.NewPersistentKnowledgeService(pointStore), Graph: service.NewKnowledgeGraphService(pointStore, relationshipStore)}
}

func requestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("X-Request-ID") == "" {
			c.Header("X-Request-ID", time.Now().UTC().Format("20060102150405.000000000"))
		}
		c.Next()
	}
}

func cors(origin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Workspace-ID, X-Request-ID")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// Error 发送JSON格式的错误响应。
func Error(c *gin.Context, status int, code, message string) {
	c.JSON(status, gin.H{"error": gin.H{"code": code, "message": message}})
}
