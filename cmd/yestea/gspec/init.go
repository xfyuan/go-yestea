package gspec

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"os"
)

func init() {
	if err := os.Setenv("YESTEA_ENV", "test"); err != nil {
		panic(fmt.Errorf("set test env failed: [%s]", err))
	}

	app.LoadConfig()

	dsn := app.GenerateDSN()

	app.DB, app.DBErr = gorm.Open("postgres", dsn)
	if app.DBErr != nil {
		panic(app.DBErr)
	}

	app.DB.AutoMigrate(&models.Todo{})
}

func ResetDB() *gorm.DB {
	app.DB.DropTableIfExists(&models.Todo{})
	app.DB.AutoMigrate(&models.Todo{})
	return app.DB
}
