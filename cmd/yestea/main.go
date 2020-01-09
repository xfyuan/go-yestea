package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"github.com/xfyuan/go-yestea/cmd/yestea/routes"
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

	r := routes.Initialize()

	dsn := app.GenerateDSN()

	app.DB, app.DBErr = gorm.Open("postgres", dsn)
	if app.DBErr != nil {
		panic(app.DBErr)
	}

	defer app.DB.Close()
	app.DB.AutoMigrate(&models.Todo{})

	log.Println("Successfully connected to database")

	if err := r.Run(fmt.Sprintf(":%v", "1234")); err != nil {
		panic(fmt.Errorf("gin run failed: [%s]", err))
	}

}
