package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"nchan-auth/app/config"
	"nchan-auth/app/handlers"
)

var cfg *config.Config

func main() {
	// Load config
	var err error
	cfg, err = config.Load(".env")
	if err != nil {
		log.Fatalf("Can't load .env file %v", err)
		return
	}

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware

	router := gin.Default()
	router.Use(injectDependencyServices())

	router.GET("/", request.Check)

	router.GET("/auth", request.Auth)

	// By default it serves on :8080 unless a
	// API_PORT environm+nt variable was defined.
	router.Run(":" + cfg.Port)
}

func injectDependencyServices() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("cfg", cfg)
		c.Next()
	}
}
