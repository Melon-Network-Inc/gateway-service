package main

import (
	"context"
	"os"

	"github.com/Melon-Network-Inc/gateway-service/pkg/config"
	"github.com/Melon-Network-Inc/gateway-service/pkg/middleware"
	"github.com/Melon-Network-Inc/gateway-service/pkg/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stderr)
}

// New creates a new server.
func New() *gin.Engine {
	conf := config.New()
	if conf == nil {
		panic("Failed to get config.")
	}
	storage := storage.New(context.Background(), conf.Redis)

	return setupRouter(storage)
}

func setupRouter(s storage.Accessor) *gin.Engine {
	authenticator := middleware.TokenAuthenticator(s)
	corsHandler := newCorsHandler()

	router := gin.Default()
	router.Use(corsHandler)


	router.Group("/account")
	router.Group("/address", authenticator)
	router.Group("/friend", authenticator)
	router.Group("/transaction", authenticator)

	return router
}

func newCorsHandler() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")

	return cors.New(config)
}

func main() {
	s := New()
	s.Run(":5000")
}
