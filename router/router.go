package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/xiaozefeng/apiserver/docs"
	"github.com/xiaozefeng/apiserver/handler/sd"
	"github.com/xiaozefeng/apiserver/handler/user"
	"github.com/xiaozefeng/apiserver/router/middleware"
	"net/http"
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

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	g.POST("/login", user.Login)

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
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
