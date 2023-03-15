package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Melon-Network-Inc/common/pkg/config"
	"github.com/Melon-Network-Inc/common/pkg/log"
	"github.com/Melon-Network-Inc/common/pkg/utils"
	"github.com/Melon-Network-Inc/gateway-service/docs"
	gatewayConfig "github.com/Melon-Network-Inc/gateway-service/pkg/config"
	"github.com/Melon-Network-Inc/gateway-service/pkg/lb"
	"github.com/Melon-Network-Inc/gateway-service/pkg/middleware"
	"github.com/Melon-Network-Inc/gateway-service/pkg/service"
	"github.com/Melon-Network-Inc/gateway-service/pkg/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	libredis "github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

var swagHandler gin.HandlerFunc

type Server struct {
	App                *gin.Engine
	Storage            storage.Accessor
	LoadBalancer       lb.Accessor
	RateLimiterSetting RateLimiterSetting
	logger             log.Logger
}

func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}

func main() {
	serverConfig := config.BuildServerConfig("config/gateway.yml")

	// create root logger tagged with server version
	logger := log.New(serverConfig.ServiceName).Default(context.Background(), serverConfig, "version", serverConfig.Version)

	tokenConfig := config.BuildTokenConfig("config/token.yml")
	if tokenConfig == nil {
		panic("Failed to get config.")
	}

	cache, err := storage.New(context.Background(), serverConfig, tokenConfig, logger)
	if err != nil {
		panic(err)
	}

	lbConf := gatewayConfig.NewLoadBalancerConfig("config/lb.yml")
	loadBalancer, err := lb.New(context.Background(), lbConf.LoadBalancer)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(log.GinLogger(logger), log.GinRecovery(logger, true))

	rateLimiterSetting, err := CreateRateLimiterSetting()
	if err != nil {
		panic(err)
	}

	s := Server{
		App:                router,
		Storage:            cache,
		LoadBalancer:       loadBalancer,
		RateLimiterSetting: rateLimiterSetting,
		logger:             logger,
	}
	s.SetupRouter()

	if !utils.IsProdEnvironment() {
		logger.Debug(router.Run(fmt.Sprintf(":%d", serverConfig.ServerPort)))
	} else {
		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", serverConfig.ServerPort),
			Handler: s.App,
		}

		go func() {
			// service connections
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Errorf("listen: %s\n", err)
			}
		}()

		// Wait for interrupt signal to gracefully shut down the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall. SIGKILL but can"t be caught, so don't need to add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Info("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.
		select {
		case <-ctx.Done():
			logger.Info("timeout of 5 seconds.")
		}
		logger.Info("Server exiting")
	}
}

func (s *Server) SetupRouter() *gin.Engine {
	forwarder := middleware.TokenForwarder()
	authenticator := middleware.TokenAuthenticator(s.Storage)

	accountService := service.NewAccountService(s.LoadBalancer.GetNextAccountServiceAddress(), s.logger)
	paymentService := service.NewPaymentService(s.LoadBalancer.GetNextPaymentServiceAddress(), s.logger)
	corsHandler := newCorsHandler()

	s.App.Use(corsHandler)
	s.App.Use(s.RateLimiterSetting.CreateRateLimiter())

	v1 := s.App.Group("api/v1")

	// Handle by Account Service
	auth := v1.Group("/auth")
	auth.POST("/email/generate", accountService.HandlePostRequest)
	auth.POST("/email/verify", accountService.HandlePostRequest)
	auth.POST("/username/validate", accountService.HandlePostRequest)
	auth.POST("/refresh", accountService.HandlePostRequest)
	auth.POST("/revoke", authenticator, accountService.HandlePostRequest)
	auth.POST("/login", accountService.HandlePostRequest)
	auth.POST("/logout", authenticator, accountService.HandlePostRequest)
	auth.GET("/forgetPassword/:email", accountService.HandleGetRequest)
	auth.POST("/resetPassword", accountService.HandlePostRequest)
	auth.GET("/google/login", accountService.HandleGetRequest)
	auth.GET("/google/callback", accountService.HandleGetRequest)

	address := v1.Group("/address")
	address.POST("/", authenticator, accountService.HandlePostRequest)
	address.GET("/query", authenticator, accountService.HandleGetRequest)
	address.GET("/recipient/:id", authenticator, accountService.HandleGetRequest)
	address.GET("/:id", authenticator, accountService.HandleGetRequest)
	address.GET("/", authenticator, accountService.HandleGetRequest)
	address.PUT("/:id", authenticator, accountService.HandleUpdateRequest)
	address.DELETE("/:id", authenticator, accountService.HandleDeleteRequest)

	friend := v1.Group("/friend")
	friend.GET("/list/user/:id", authenticator, accountService.HandleGetRequest)
	friend.GET("/list", authenticator, accountService.HandleGetRequest)
	friend.GET("/query/user/:id", authenticator, accountService.HandleGetRequest)
	friend.GET("/query", authenticator, accountService.HandleGetRequest)
	friend.DELETE("/:id", authenticator, accountService.HandleDeleteRequest)

	request := v1.Group("/request")
	request.POST("/", authenticator, accountService.HandlePostRequest)
	request.DELETE("/:id", authenticator, accountService.HandleDeleteRequest)
	request.PUT("/confirm/:id", authenticator, accountService.HandleUpdateRequest)
	request.PUT("/reject/:id", authenticator, accountService.HandleUpdateRequest)
	request.GET("/query/user/:id", authenticator, accountService.HandleGetRequest)
	request.GET("/query", authenticator, accountService.HandleGetRequest)
	request.GET("/", authenticator, accountService.HandleGetRequest)

	account := v1.Group("/account")
	account.POST("/uploadProfile", authenticator, accountService.HandlePostRequestWithAttachment)
	account.GET("/downloadProfile/:id", authenticator, accountService.HandleGetRequest)
	account.POST("/", forwarder, accountService.HandlePostRequest)
	account.GET("/self", authenticator, accountService.HandleGetRequest)
	account.GET("/:id", authenticator, accountService.HandleGetRequest)
	account.PUT("/:id", authenticator, accountService.HandleUpdateRequest)
	account.DELETE("/:id", authenticator, accountService.HandleDeleteRequest)
	account.PUT("/security/:id", authenticator, accountService.HandleUpdateRequest)
	account.PUT("/activate", authenticator, accountService.HandleUpdateRequest)
	account.PUT("/deactivate", authenticator, accountService.HandleUpdateRequest)
	account.POST("/addBackup", authenticator, accountService.HandlePostRequest)
	account.POST("/restoreBackup", authenticator, accountService.HandlePostRequest)
	account.PUT("/completeSurvey", authenticator, accountService.HandleUpdateRequest)
	account.PUT("/completeWalletCreation", authenticator, accountService.HandleUpdateRequest)

	search := v1.Group("/search")
	search.GET("/user/:keyword", authenticator, accountService.HandleGetRequest)

	notification := v1.Group("/notification")
	notification.GET("/query", authenticator, accountService.HandleGetRequest)
	notification.GET("/", authenticator, accountService.HandleGetRequest)
	notification.DELETE("/:id", authenticator, accountService.HandleDeleteRequest)

	whitelist := v1.Group("/whitelist")
	whitelist.POST("/email/generate", accountService.HandlePostRequest)
	whitelist.POST("/email/verify", accountService.HandlePostRequest)
	whitelist.POST("/", forwarder, accountService.HandlePostRequest)
	whitelist.GET("/name/:name", authenticator, accountService.HandleGetRequest)
	whitelist.GET("/email/:email", authenticator, accountService.HandleGetRequest)
	whitelist.GET("/phone/:phone", authenticator, accountService.HandleGetRequest)
	whitelist.GET("/query", authenticator, accountService.HandleGetRequest)
	whitelist.GET("/", authenticator, accountService.HandleGetRequest)
	whitelist.DELETE("/:id", authenticator, accountService.HandleDeleteRequest)

	referral := v1.Group("/referral")
	referral.GET("/create", authenticator, accountService.HandleGetRequest)
	referral.POST("/accept", accountService.HandlePostRequest)
	referral.POST("/nextAvailable", authenticator, accountService.HandlePostRequest)
	referral.GET("/revoke/:id", authenticator, accountService.HandleGetRequest)
	referral.GET("/list", authenticator, accountService.HandleGetRequest)
	referral.GET("/count/accepted", authenticator, accountService.HandleGetRequest)
	referral.GET("/count/left", authenticator, accountService.HandleGetRequest)
	referral.GET("/:id", authenticator, accountService.HandleGetRequest)
	referral.DELETE("/:id", authenticator, accountService.HandleDeleteRequest)

	setting := v1.Group("/setting")
	device := setting.Group("/device")
	device.GET("/", authenticator, accountService.HandleGetRequest)
	social := setting.Group("/socialMedia")
	social.POST("/link", authenticator, accountService.HandlePostRequest)
	social.DELETE("/link", authenticator, accountService.HandleDeleteRequest)
	social.GET("/link/:platform", authenticator, accountService.HandleGetRequest)
	social.GET("/link/list", authenticator, accountService.HandleGetRequest)

	// Handle by Payment Service
	transaction := v1.Group("/transaction")
	transaction.POST("/", authenticator, paymentService.HandlePostRequest)
	transaction.GET("/query/:id", authenticator, paymentService.HandleGetRequest)
	transaction.GET("/user/:id", authenticator, paymentService.HandleGetRequest)
	transaction.GET("/:id", authenticator, paymentService.HandleGetRequest)
	transaction.GET("/", authenticator, paymentService.HandleGetRequest)
	transaction.PUT("/:id", authenticator, paymentService.HandleUpdateRequest)
	transaction.DELETE("/:id", authenticator, paymentService.HandleDeleteRequest)

	activity := v1.Group("/activity")
	activity.GET("/query", authenticator, paymentService.HandleGetRequest)
	activity.GET("/", authenticator, paymentService.HandleGetRequest)

	news := v1.Group("/news")
	news.GET("/query", authenticator, paymentService.HandleGetRequest)

	

	if !utils.IsProdEnvironment() && swagHandler != nil {
		s.buildSwagger()
		s.App.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return s.App
}

type RateLimiterSetting struct {
	RateSetting limiter.Rate
	CacheOption libredis.Options
	Store       limiter.Store
}

func (r *RateLimiterSetting) CreateRateLimiter() gin.HandlerFunc {
	return mgin.NewMiddleware(limiter.New(r.Store, r.RateSetting))
}

func CreateRateLimiterSetting() (RateLimiterSetting, error) {
	rate := limiter.Rate{
		Period: 1 * time.Hour,
		Limit:  1000,
	}

	option, err := libredis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		return RateLimiterSetting{}, err
	}

	store, err := sredis.NewStoreWithOptions(libredis.NewClient(option), limiter.StoreOptions{
		Prefix:   "mw_rate_limit",
		MaxRetry: 3,
	})
	if err != nil {
		return RateLimiterSetting{}, err
	}

	setting := RateLimiterSetting{
		RateSetting: rate,
		CacheOption: *option,
		Store:       store,
	}
	return setting, nil
}

func (s *Server) buildSwagger() {
	docs.SwaggerInfo.Title = "Melon Wallet Service API"
	docs.SwaggerInfo.Description = "This is backend server for Melon Wallet."
	docs.SwaggerInfo.Version = "1.0"
	if utils.IsProdEnvironment() {
		docs.SwaggerInfo.Host = "prod.melonnetwork.io:8080"
	} else if utils.IsStagingEnvironment() {
		docs.SwaggerInfo.Host = "staging.melonnetwork.io"
	} else {
		docs.SwaggerInfo.Host = "localhost:8080"
	}
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func newCorsHandler() gin.HandlerFunc {
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowAllOrigins = true
	defaultConfig.AddAllowHeaders("Authorization", "AuthorizationToken", "RegistrationSession")

	return cors.New(defaultConfig)
}
