package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaozefeng/apiserver/handler/sd"
	"github.com/xiaozefeng/apiserver/router/middleware"
	"net/http"
	"github.com/xiaozefeng/apiserver/handler/user"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middleware
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 handler
	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "The incorrect API route")
	})

	u := g.Group("/v1/user")
	{
		u.POST("/:username", user.Create)
	}

	// the health check handlers
	group := g.Group("/sd")
	{
		group.GET("/health", sd.HealthCheck)
		group.GET("/disk", sd.DiskCheck)
		group.GET("/cpu", sd.CPUCheck)
		group.GET("/ram", sd.RAMCheck)

	}
	return g
}
