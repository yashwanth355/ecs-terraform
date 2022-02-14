package main

import (
	"go.uber.org/zap"
	"os"
	Config "usermvc/config"
	"usermvc/routes"
	logger2 "usermvc/utility/logger"
)

func main() {
	loggerMgr := logger2.InitLogger()
	zap.ReplaceGlobals(loggerMgr)
	defer loggerMgr.Sync() // flushes buffer, if any
	logger := loggerMgr.Sugar()
	logger.Info("START logging app service")
	Config.LoadConfig()

	if err := routes.SetupRouter().Run(); err != nil {
		logger.Panic("error while running server", err.Error())
		os.Exit(0)
	}
}
