package app

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
//var Config appConfig
var Viper = viper.New()
var DB *gorm.DB
var DBErr error

//type appConfig struct {
// the shared DB ORM object
//DB *gorm.DB
// the error thrown be GORM when using DB ORM object
//DBErr error
// the server port. Defaults to 8080
//ServerPort int `mapstructure:"server_port"`
// the data source name (DSN) for connecting to the database. required.
//DSN string `mapstructure:"dsn"`
//Database dbConfog
// the API key needed to authorize to API. required.
//ApiKey string `mapstructure:"api_key"`
//}

//type dbConfog struct {
//	Host     string `mapstructure:"host"`
//Port     string `mapstructure:"port"`
//DBName   string `mapstructure:"dbname"`
//Username string `mapstructure:"username"`
//Password string `mapstructure:"password"`
//}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	Viper.SetEnvPrefix("yestea")
	Viper.BindEnv("env")
	//Viper.AutomaticEnv()
	env := Viper.Get("ENV")
	if env == "PRODUCTION" {
		Viper.SetConfigName("prod")
	} else {
		Viper.SetConfigName("dev")
	}

	Viper.SetConfigType("yaml")

	//Config.DSN = Viper.Get("DSN").(string)
	//Config.ApiKey = Viper.Get("API_KEY").(string)
	//Viper.SetDefault("server_port", 1234)

	for _, path := range configPaths {
		Viper.AddConfigPath(path)
	}
	//if err := Viper.ReadInConfig(); err != nil {
	//	return fmt.Errorf("failed to read the configuration file: %s", err)
	//}
	return Viper.ReadInConfig()
}
