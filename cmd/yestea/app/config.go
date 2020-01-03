package app

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
//var Config appConfig
var DB *gorm.DB
var DBErr error

// LoadConfig loads config from files
func LoadConfig() error {
	viper.SetEnvPrefix("yestea")
	viper.BindEnv("env")
	//viper.AutomaticEnv()
	env := viper.Get("ENV")
	if env == "PRODUCTION" {
		viper.SetConfigName("prod")
	} else {
		viper.SetConfigName("dev")
	}

	viper.SetConfigType("yaml")

	viper.AddConfigPath("./config")
	//if err := viper.ReadInConfig(); err != nil {
	//	return fmt.Errorf("failed to read the configuration file: %s", err)
	//}
	return viper.ReadInConfig()
}
