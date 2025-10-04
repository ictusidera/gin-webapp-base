package router

import (
	"github.com/gin-gonic/gin"

	"github.com/bandersnach/sample-app/config"
	"github.com/bandersnach/sample-app/internal/handler"
)

// New constructs the HTTP router with all public routes mounted.
func New(cfg config.Config) *gin.Engine {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/healthz", handler.Health)

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		v1.GET("/status", handler.Health)
	}

	return r
}
