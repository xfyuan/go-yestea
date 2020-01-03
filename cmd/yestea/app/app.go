package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"runtime"
)

// Config is global object that holds all application level variables.
//var Config appConfig
var DB *gorm.DB
var DBErr error

// LoadConfig loads config from files
func LoadConfig() {
	viper.SetEnvPrefix("yestea")
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	viper.AddConfigPath(rootPath() + "config")

	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read common config failed: [%s]", err))
	}

	env := viper.GetString("ENV")
	if env == "" {
		viper.SetConfigName("dev")
	}
	viper.SetConfigName(env)
	if err := viper.MergeInConfig(); err != nil {
		panic(fmt.Errorf("merge environment config failed: [%s]", err))
	}
}

func rootPath() string {
	// find current path
	_, filename, _, _ := runtime.Caller(0)
	currentPath := filepath.Dir(filename)
	log.Printf("[DEBUG] current path is %s", currentPath)
	rootdir := currentPath + "/../../../"
	return rootdir
}

func GenerateDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbname"),
	)
}
