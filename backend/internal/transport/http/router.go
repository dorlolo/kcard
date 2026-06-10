package httpapi

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"kcardDesgin/backend/internal/app"
	"kcardDesgin/backend/internal/domain"
	"kcardDesgin/backend/internal/service"
	"kcardDesgin/backend/internal/transport/http/middleware"
)

func NewRouter(container *app.Container) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), requestID(), cors(container.Config.AllowedFrontendOrigin), middleware.PlaceholderAuth())
	r.GET("/api/v1/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().UTC()}) })
	api := r.Group("/api/v1")
	RegisterMaterialRoutes(api, MaterialHandler{Service: service.MaterialService{Store: &service.MemoryMaterialStore{}, Jobs: &service.MemoryJobEnqueuer{}}})
	RegisterKnowledgeRoutes(api, KnowledgeHandler{Service: &service.KnowledgeService{Points: map[domain.ID]domain.KnowledgePoint{}}, Graph: &service.KnowledgeGraphService{Points: map[domain.ID]domain.KnowledgePoint{}}})
	RegisterCardRoutes(api, CardHandler{})
	RegisterReviewRoutes(api, ReviewHandler{})
	RegisterPromptRoutes(api, PromptHandler{})
	RegisterDashboardRoutes(api, DashboardHandler{})
	RegisterPortabilityRoutes(api, PortabilityHandler{})
	r.NoRoute(func(c *gin.Context) { Error(c, http.StatusNotFound, "not_found", "route not found") })
	return r
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

func Error(c *gin.Context, status int, code, message string) {
	c.JSON(status, gin.H{"error": gin.H{"code": code, "message": message}})
}
