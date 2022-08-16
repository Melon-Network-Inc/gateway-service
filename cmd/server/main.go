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
	accountService := service.NewAccountService("http://localhost:6000")
	paymentService := service.NewPaymentService("http://localhost:7000")
	corsHandler := newCorsHandler()

	router := gin.Default()
	router.Use(corsHandler)

	v1 := router.Group("api/v1")

	// Handle by Account Service
	auth := v1.Group("/auth")
	auth.POST("/email/generate", accountService.HandlePostRequest)
	auth.POST("/email/veirfy", accountService.HandlePostRequest)
	auth.POST("/login", accountService.HandlePostRequest)
	auth.GET("/logout", authenticator, accountService.HandleGetRequest)

	address := v1.Group("/address")
	address.POST("/", authenticator, accountService.HandlePostRequest)
	address.GET("/", authenticator, accountService.HandleGetRequest)
	address.GET("/:id", authenticator, accountService.HandleGetRequest)
	address.PUT("/:id", authenticator, accountService.HandleUpdateRequest)
	address.DELETE("/:id", authenticator, accountService.HandleDeleteRequest)

	friend := v1.Group("/friend")
	friend.GET("/list/", authenticator, accountService.HandleGetRequest)
	friend.GET("/list/user/:id", authenticator, accountService.HandleGetRequest)
	friend.DELETE("/:id", authenticator, accountService.HandleDeleteRequest)

	request := v1.Group("/request")
	request.POST("/", authenticator, accountService.HandlePostRequest)
	request.DELETE("/:id", authenticator, accountService.HandleDeleteRequest)
	request.PUT("/confirm/:id", authenticator, accountService.HandleUpdateRequest)
	request.PUT("/reject/:id", authenticator, accountService.HandleUpdateRequest)
	request.GET("/", authenticator, accountService.HandleGetRequest)

	account := v1.Group("/account")
	account.POST("/", accountService.HandlePostRequest)
	account.GET("/:id", authenticator, accountService.HandleGetRequest)
	account.PUT("/:id", authenticator, accountService.HandleUpdateRequest)
	account.DELETE("/:id", authenticator, accountService.HandleDeleteRequest)
	account.PUT("/security/:id", authenticator, accountService.HandleUpdateRequest)
	account.PUT("/activate", authenticator, accountService.HandleUpdateRequest)
	account.PUT("/deactivate", authenticator, accountService.HandleUpdateRequest)

	activity := v1.Group("/activity")
	activity.GET("/", authenticator, accountService.HandleGetRequest)

	whitelist := v1.Group("/whitelist")
	whitelist.POST("/", accountService.HandlePostRequest)
	whitelist.GET("/name/:name", accountService.HandleGetRequest)
	whitelist.GET("/email/:email", accountService.HandleGetRequest)
	whitelist.GET("/phone/:phone", accountService.HandleGetRequest)
	whitelist.GET("/", accountService.HandleGetRequest)
	whitelist.DELETE("/:id", accountService.HandleDeleteRequest)

	// Handle by Payment Service
	transaction := v1.Group("/transactions")
	transaction.POST("/", authenticator, paymentService.HandlePostRequest)
	transaction.GET("/user/:id", authenticator, paymentService.HandleGetRequest)
	transaction.GET("/:id", authenticator, paymentService.HandleGetRequest)
	transaction.GET("/", authenticator, paymentService.HandleGetRequest)
	transaction.PUT("/:id", authenticator, paymentService.HandleUpdateRequest)
	transaction.DELETE("/:id", authenticator, paymentService.HandleDeleteRequest)

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
