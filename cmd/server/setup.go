package server

import (
	"github.com/Melon-Network-Inc/gateway-service/pkg/config"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// New creates a new server.
func New() *gin.Engine {
	conf := config.New()
	if conf == nil {
		panic("Failed to get config.")
	}

	return setupRouter()
}

func setupRouter() *gin.Engine {
	corsHandler := newCorsHandler()

	router := gin.Default()
	router.Use(corsHandler)

	return router
}

func newCorsHandler() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")

	return cors.New(config)
}
