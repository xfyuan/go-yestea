package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/controllers"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"log"
)

func main() {
	// load application configurations
	app.LoadConfig()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
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
