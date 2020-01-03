package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/controllers"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"log"
)

func main() {
	// load application configurations
	app.LoadConfig("./config")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos/:id", controllers.GetTodo)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbname"),
	)

	app.DB, app.DBErr = gorm.Open("postgres", dsn)
	if app.DBErr != nil {
		panic(app.DBErr)
	}

	defer app.DB.Close()
	app.DB.AutoMigrate(&models.Todo{})

	log.Println("Successfully connected to database")

	r.Run(fmt.Sprintf(":%v", "1234"))
}
