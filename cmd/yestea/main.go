package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/controllers"
	_ "github.com/xfyuan/go-yestea/cmd/yestea/docs"
	"github.com/xfyuan/go-yestea/cmd/yestea/middlewares"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"log"
)

// @title Yestea Swagger API
// @version 1.0
// @description Swagger API for Golang Project Yestea.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email yuan.xiaofeng@gmail.com

// @license.name MIT
// @license.url https://github.com/xfyuan/go-yestea/blob/master/LICENSE

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// load application configurations
	app.LoadConfig()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.Use(middlewares.Auth())
		v1.GET("/todos/:id", controllers.GetTodo)
	}

	dsn := app.GenerateDSN()

	app.DB, app.DBErr = gorm.Open("postgres", dsn)
	if app.DBErr != nil {
		panic(app.DBErr)
	}

	defer app.DB.Close()
	app.DB.AutoMigrate(&models.Todo{})

	log.Println("Successfully connected to database")

	r.Run(fmt.Sprintf(":%v", "1234"))
}
