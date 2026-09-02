package main

import (
	"os"
	"whoknowsyourdata/app"
	"whoknowsyourdata/logger"
)

func main() {
	logger := logger.NewLogger()
	env, err := GetEnv(logger)
	if err != nil {
		logger.Error("%v", err)
		os.Exit(1)
	}

	if env.App.TRUSTED_CTX {
		logger.Info("App launched in trusted context")
	}

	app.ServerBootstrap(logger, *env)
}
