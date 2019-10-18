package main

import (
	"github.com/devit-tel/gogo-blueprint/config"
	"github.com/devit-tel/jaegerstart"
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
)

func main() {
	// Load config
	appConfig := config.Get()

	// Gin setup
	router := gin.Default()

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
