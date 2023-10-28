package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kararnab/handsongo/pkg/auth/middleware"
	"github.com/kararnab/handsongo/pkg/controllers"
	"github.com/kararnab/handsongo/pkg/initialize"
	legacyHttp "github.com/kararnab/handsongo/pkg/legacy_http"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	initialize.LoadEnv() //Load environment variables

	if os.Getenv("USE_GIN") == "true" {
		serveApplication()
	} else {
		legacyHttp.ServeApplication() // Start the server
	}
}

func serveApplication() {
	router := gin.Default() //gin.Default() uses gin.New() internally to create an engine
	v1 := router.Group("api")
	v1.POST("/register", controllers.Register)
	v1.POST("/login", controllers.Login)
	v1.GET("/health", controllers.HealthCheck)
	v1.Use(middleware.AuthMiddleware())
	v1.GET("/home", controllers.HomeRouter)

	port := legacyHttp.AssignPort()
	log.Printf(fmt.Sprintf("Starting gin server on %s", port))
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal().Err(err)
		return
	}
}
