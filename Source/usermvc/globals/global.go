package globals

import (
	"usermvc/config"
)

var (
	appConfig *Config.AppConfig
)

func GetConfig() *Config.AppConfig {
	if appConfig == nil {
		return Config.LoadConfig()
	}
	return appConfig
}
