package service

import (
	"github.com/Melon-Network-Inc/gateway-service/pkg/processor"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type PaymentService interface {
	HandlePostRequest(ctx *gin.Context)
	HandleUpdateRequest(ctx *gin.Context)
	HandleGetRequest(ctx *gin.Context)
	HandleDeleteRequest(ctx *gin.Context)
}

type paymentService struct {
	serviceUrlPrefix string
}

func NewPaymentService(serviceUrlPrefix string) PaymentService {
	return &paymentService{
		serviceUrlPrefix: serviceUrlPrefix}
}

func (s *paymentService) HandlePostRequest(ctx *gin.Context) {
	client := resty.New()
	resp, err := processor.PrepareRequest(ctx, client).
		Post(s.serviceUrlPrefix + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService due to", err)
        return
    }
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *paymentService) HandleUpdateRequest(ctx *gin.Context) {
	client := resty.New()
	resp, err := processor.PrepareRequest(ctx, client).
		Put(s.serviceUrlPrefix + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService due to", err)
        return
    }
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *paymentService) HandleGetRequest(ctx *gin.Context) {
	client := resty.New()
	resp, err := processor.PrepareRequest(ctx, client).
		Get(s.serviceUrlPrefix + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService due to", err)
        return
    }
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *paymentService) HandleDeleteRequest(ctx *gin.Context) {
	client := resty.New()
	resp, err := processor.PrepareRequest(ctx, client).
		Delete(s.serviceUrlPrefix + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService due to", err)
        return
    }
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}