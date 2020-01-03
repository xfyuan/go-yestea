package gspec

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"github.com/xfyuan/go-yestea/cmd/yestea/app"
	"github.com/xfyuan/go-yestea/cmd/yestea/models"
	"log"
	"os"
	"runtime"
	"strings"
)

func init() {
	if err := os.Setenv("YESTES_ENV", "test"); err != nil {
		panic(fmt.Errorf("set test env failed: [%s]", err))
	}

	// find runtime current path
	_, filename, _, _ := runtime.Caller(0)
	log.Printf("[DEBUG] [UNIT TEST] current path is %s", filename)

	currentPathSegments := strings.Split(filename, "/")
	basedir := strings.Join(currentPathSegments[:len(currentPathSegments)-4], "/")
	configdir := fmt.Sprintf("%s/config", basedir)

	app.LoadConfig(configdir)

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

	app.DB.AutoMigrate(&models.Todo{})
}

func ResetDB() *gorm.DB {
	app.DB.DropTableIfExists(&models.Todo{})
	app.DB.AutoMigrate(&models.Todo{})
	return app.DB
}
