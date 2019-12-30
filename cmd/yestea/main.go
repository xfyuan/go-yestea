package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xfyuan/go-yestea/cmd/yestea/apis"
	"github.com/xfyuan/go-yestea/cmd/yestea/config"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"log"
	"os"
)

var (
	GO_VERSION string
	BUILD_TIME string
	VSN        string
)

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Version: %s\n", VSN)
		fmt.Printf("Build Time : %s\n", BUILD_TIME)
		fmt.Printf("%s\n", GO_VERSION)
		return
	}
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	fmt.Println(config.Config)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos/:id", apis.GetTodo)
	}

	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	defer config.Config.DB.Close()
	config.Config.DB.AutoMigrate(&models.Todo{})

	log.Println("Successfully connected to database")

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}
