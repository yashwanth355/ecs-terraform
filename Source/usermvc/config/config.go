//Config/Database.go
package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

const (
	APP_ENV            = "APP_ENV"
	PRODUCTION         = "production"
	DEVELOPMENT        = "development"
	DBPORT             = "DB_PORT"
	DBNAME             = "DB_NAME"
	DBUSERNAME         = "DB_USERNAME"
	DBPASSWORD         = "DB_PASSWORD"
	DBHOST             = "DB_HOST"
	AWSCOGNITOREGION   = "AWS_COGNITO_REGION"
	COGNITOUSERPOOLID  = "COGNITO_USER_POOL_ID"
	COGNITOAPPCLIENTID = "COGNITO_APP_CLIENT_ID"
)

// DBConfig represents db configuration
var (
	configDir      = "../config"
	configfile     = "development"
	configFileType = "yaml"
)
var (
	Config *AppConfig
)

type dbconfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AppConfig struct {
	App      *appcofig
	DbConfig *dbconfig
	Logger   *logger
}

type logger struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
	LocalTime  bool   `json:"local_time"`
	Compress   bool   `json:"compress"`
}

type appcofig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func LoadConfig() *AppConfig {

	viper.AutomaticEnv()
	env := viper.Get(APP_ENV)
	if env == PRODUCTION {
		configfile = PRODUCTION
	}

	fmt.Println(configfile)
	viper.SetConfigName(configfile)
	viper.AddConfigPath(configDir)
	viper.SetConfigType(configFileType)
	var configuration AppConfig
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	viper.SetDefault("database.dbname", "test_db")
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	configuration.DbConfig = &dbconfig{
		//Host:     viper.GetString(DBHOST),
		//Port:     viper.GetInt(DBPORT),
		//User:     viper.GetString(DBUSERNAME),
		//Name:     viper.GetString(DBNAME),
		//Password: viper.GetString(DBPASSWORD),
		Host:     "ccl-psql-dev.cclxlbtddgmn.ap-south-1.rds.amazonaws.com",
		Port:     5432,
		User:     "postgres",
		Name:     "ccldevdb",
		Password: "Ccl_RDS_DB#2022",
	}

	fmt.Println(configuration.DbConfig)
	return &configuration
}
