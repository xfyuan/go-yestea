package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
//var Config appConfig
var DB *gorm.DB
var DBErr error

// LoadConfig loads config from files
func LoadConfig(dir string) {
	viper.SetEnvPrefix("yestea")
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir)

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
