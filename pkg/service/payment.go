package service

import (
	"net/http"

	"github.com/Melon-Network-Inc/common/pkg/log"
	"github.com/Melon-Network-Inc/gateway-service/pkg/processor"

	"github.com/gin-gonic/gin"
)

type PaymentService interface {
	HandlePostRequest(ctx *gin.Context)
	HandleUpdateRequest(ctx *gin.Context)
	HandleGetRequest(ctx *gin.Context)
	HandleDeleteRequest(ctx *gin.Context)
	HandleServiceUnavailable(ctx *gin.Context, err error)
}

type paymentService struct {
	serviceUrlPrefix string
	logger log.Logger
}

func NewPaymentService(serviceUrlPrefix string, logger log.Logger) PaymentService {
	return &paymentService{
		serviceUrlPrefix: serviceUrlPrefix,
		logger: logger,
	}
}

func (s *paymentService) HandlePostRequest(ctx *gin.Context) {
	client := CreateRetryRestyClient()
	resp, err := processor.PrepareRequest(ctx, client).
		Post(s.serviceUrlPrefix + ctx.Request.URL.String())
	if err != nil {
		s.HandleServiceUnavailable(ctx, err)
		return
	}

	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *paymentService) HandleUpdateRequest(ctx *gin.Context) {
	client := CreateRetryRestyClient()
	resp, err := processor.PrepareRequest(ctx, client).
		Put(s.serviceUrlPrefix + ctx.Request.URL.String())
	if err != nil {
		s.HandleServiceUnavailable(ctx, err)
		return
	}

	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *paymentService) HandleGetRequest(ctx *gin.Context) {
	client := CreateRetryRestyClient()
	resp, err := processor.PrepareRequest(ctx, client).
		Get(s.serviceUrlPrefix + ctx.Request.URL.String())
	if err != nil {
		s.HandleServiceUnavailable(ctx, err)
		return
	}

	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *paymentService) HandleDeleteRequest(ctx *gin.Context) {
	client := CreateRetryRestyClient()
	resp, err := processor.PrepareRequest(ctx, client).
		Delete(s.serviceUrlPrefix + ctx.Request.URL.String())
    if err != nil {
    	s.HandleServiceUnavailable(ctx, err)
        return
    }
    
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *paymentService) HandleServiceUnavailable(ctx *gin.Context, err error) {
	s.logger.Errorf("Payment Service Failure: ", err)
	ctx.Status(http.StatusServiceUnavailable)
}