package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/controllers"
	"github.com/xfyuan/go-yestea/cmd/yestea/httputils"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"log"
	"net/http"

	_ "github.com/xfyuan/go-yestea/cmd/yestea/docs"
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
func main() {
	// load application configurations
	app.LoadConfig()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.Use(auth())
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

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			httputils.NewError(c, http.StatusUnauthorized, errors.New("Authorization header is required!"))
			c.Abort()
		}
		if authHeader != viper.Get("apikey") {
			httputils.NewError(c, http.StatusUnauthorized, fmt.Errorf("Not authorized to this operation: api_key=%s", authHeader))
			c.Abort()
		}
		c.Next()
	}
}
