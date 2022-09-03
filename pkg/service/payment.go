package service

import (
	"github.com/Melon-Network-Inc/common/pkg/log"
	"github.com/Melon-Network-Inc/gateway-service/pkg/processor"

	"github.com/gin-gonic/gin"
)

type PaymentService interface {
	HandlePostRequest(ctx *gin.Context)
	HandleUpdateRequest(ctx *gin.Context)
	HandleGetRequest(ctx *gin.Context)
	HandleDeleteRequest(ctx *gin.Context)
}

type paymentService struct {
	serviceUrlPrefix string
	logger           log.Logger
}

func NewPaymentService(serviceUrlPrefix string, logger log.Logger) PaymentService {
	return &paymentService{
		serviceUrlPrefix: serviceUrlPrefix,
		logger:           logger,
	}
}

func (s *paymentService) HandleGetRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Get(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}

func (s *paymentService) HandlePostRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Post(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}

func (s *paymentService) HandleUpdateRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Put(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}

func (s *paymentService) HandleDeleteRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Delete(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}
