package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/xfyuan/go-yestea/cmd/yestea/controllers"
	_ "github.com/xfyuan/go-yestea/cmd/yestea/docs"
	"github.com/xfyuan/go-yestea/cmd/yestea/middlewares"
)

func Initialize() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes := registerRoutes(r)
	return routes
}

func registerRoutes(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api/v1").Use(middlewares.Auth())
	{
		v1.GET("/todos/:id", controllers.GetTodo)
	}
	return r
}
