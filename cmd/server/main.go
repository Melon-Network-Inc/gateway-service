package main

import (
	"context"
	"os"

	"github.com/Melon-Network-Inc/gateway-service/pkg/config"
	"github.com/Melon-Network-Inc/gateway-service/pkg/middleware"
	"github.com/Melon-Network-Inc/gateway-service/pkg/service"
	"github.com/Melon-Network-Inc/gateway-service/pkg/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/Melon-Network-Inc/gateway-service/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Version indicates the current version of the application.
var Version = "1.0.0"
var swagHandler gin.HandlerFunc

func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerfiles.Handler)
	logrus.SetOutput(os.Stderr)
}

func main() {
	conf := config.New()
	if conf == nil {
		panic("Failed to get config.")
	}

	storage := storage.New(context.Background(), conf.Redis)

	setupRouter(storage).Run(":8080")
}

func setupRouter(s storage.Accessor) *gin.Engine {
	authenticator := middleware.TokenAuthenticator(s)
	accountService := service.NewAccountService()
	paymentService := service.NewPaymentService()
	corsHandler := newCorsHandler()

	router := gin.Default()
	router.Use(corsHandler)

	// Handle by Account Service
	router.Group("/auth", accountService.HandleRequest)
	router.Group("/whitelist", accountService.HandleRequest)
	router.Group("/account", authenticator, accountService.HandleRequest)
	router.Group("/activity", authenticator, accountService.HandleRequest)
	router.Group("/address", authenticator, accountService.HandleRequest)
	router.Group("/friend", authenticator, accountService.HandleRequest)

	// Handle by Payment Service
	router.Group("/transaction", authenticator, paymentService.HandleRequest)

	if swagHandler != nil {
		buildSwagger()
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	return router
}

func newCorsHandler() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")

	return cors.New(config)
}

func buildSwagger() {
	docs.SwaggerInfo.Title = "Melon Wallet Service API"
	docs.SwaggerInfo.Description = "This is backend server for Melon Wallet."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
