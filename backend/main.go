package main

import (
	"os"
	"whoknowsyourdata/app"
	"whoknowsyourdata/logger"
)

func main() {
	log := logger.NewLogger(nil)
	env, err := GetEnv(log)
	if err != nil {
		log.Error("%v", err)
		os.Exit(1)
	}

	// Error should not be raised here
	_ = log.SetLevel(env.App.LOG_LEVEL)

	if env.App.TRUSTED_CTX {
		log.Info("App launched in trusted context")
	}

	err = app.ServerBootstrap(log, *env)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
