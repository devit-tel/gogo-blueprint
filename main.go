package main

import (
	"github.com/devit-tel/gogo-blueprint/config"
	"github.com/devit-tel/jaegerstart"
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	ginlogrus "github.com/toorop/gin-logrus"
)

func main() {
	// Load config
	appConfig := config.Get()

	// Init log format & stdout to logstash (udp)
	log := setupLog()
	setupHookToLogstash(log, appConfig)

	// Gin setup
	router := gin.New()

	// Set custom log for gin
	router.Use(ginlogrus.Logger(log), gin.Recovery())

	// Jaeger setup
	if appConfig.JaegerEndpoint != "" {
		tracer, closer := setupJaeger(appConfig)
		defer closer.Close()
		router.Use(ginhttp.Middleware(tracer, jaegerstart.LogRequestOption()))
	}

	// Register route to gin
	_ = newApp(appConfig).RegisterRoute(router)

	// Gin start listen
	_ = router.Run()
}
