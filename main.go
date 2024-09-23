package main

import (
	"go-api/config"
	"go-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.ErrorF("config initialization: %v", err)
		return

	}

	router.Initialize()
}
