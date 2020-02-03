package main

import (
	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/devit-tel/gogo-blueprint/config"
)

func main() {
	// Load config
	appConfig := config.Get()

	// Init log format
	log := setupLog()

	// Gin setup
	router := gin.New()

	// Set custom log for gin
	router.Use(ginlogrus.Logger(log), gin.Recovery())

	// Jaeger setup
	closer := setupJaeger(appConfig)
	defer func() {
		if err := closer.Close(); err != nil {
			log.Error(err)
		}
	}()

	// Register route to gin
	_ = newApp(appConfig).RegisterRoute(router)

	// Gin start listen
	_ = router.Run()
}
