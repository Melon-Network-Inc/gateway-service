package service

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const PaymentUrlPrefix = "http://localhost:7000"

type PaymentService interface {
	HandlePostRequest(ctx *gin.Context)
	HandleUpdateRequest(ctx *gin.Context)
	HandleGetRequest(ctx *gin.Context)
	HandleDeleteRequest(ctx *gin.Context)
}

type paymentService struct {}

func NewPaymentService() PaymentService {
	return &paymentService{}
}

func (s *paymentService) HandlePostRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	if ctx.GetString("username") != "" && ctx.GetString("user_id") != "" {
		userData["user"] = ctx.Value("username").(string)
		userData["user_id"] = ctx.Value("user_id").(string)
	}

	resp, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        Post(PaymentUrlPrefix + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService due to", err)
        return
    }
	ctx.JSON(resp.StatusCode(), resp.Body())
}

func (s *paymentService) HandleUpdateRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	if ctx.GetString("username") != "" && ctx.GetString("user_id") != "" {
		userData["user"] = ctx.Value("username").(string)
		userData["user_id"] = ctx.Value("user_id").(string)
	}

	resp, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        Put(PaymentUrlPrefix + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService due to", err)
        return
    }
	ctx.JSON(resp.StatusCode(), resp.Body())
}

func (s *paymentService) HandleGetRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	if ctx.GetString("username") != "" && ctx.GetString("user_id") != "" {
		userData["user"] = ctx.Value("username").(string)
		userData["user_id"] = ctx.Value("user_id").(string)
	}

	resp, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        Get(PaymentUrlPrefix + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService due to", err)
        return
    }
	ctx.JSON(resp.StatusCode(), resp.Body())
}

func (s *paymentService) HandleDeleteRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	if ctx.GetString("username") != "" && ctx.GetString("user_id") != "" {
		userData["user"] = ctx.Value("username").(string)
		userData["user_id"] = ctx.Value("user_id").(string)
	}

	resp, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        Delete(PaymentUrlPrefix + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService due to", err)
        return
    }
	ctx.JSON(resp.StatusCode(), resp.Body())
}