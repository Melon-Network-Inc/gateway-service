package service

import (
	"net/http"

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
	var res []byte

	_, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        SetResult(&res).
        Post("http://localhost:7000" + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService")
        return
    }
	ctx.JSON(http.StatusOK, &res)
}

func (s *paymentService) HandleUpdateRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	if ctx.GetString("username") != "" && ctx.GetString("user_id") != "" {
		userData["user"] = ctx.Value("username").(string)
		userData["user_id"] = ctx.Value("user_id").(string)
	}
	var res []byte

	_, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        SetResult(&res).
        Patch("http://localhost:7000" + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService")
        return
    }
	ctx.JSON(http.StatusOK, &res)
}

func (s *paymentService) HandleGetRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	if ctx.GetString("username") != "" && ctx.GetString("user_id") != "" {
		userData["user"] = ctx.Value("username").(string)
		userData["user_id"] = ctx.Value("user_id").(string)
	}
	var res []byte

	_, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        SetResult(&res).
        Get("http://localhost:7000" + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService")
        return
    }
	ctx.JSON(http.StatusOK, &res)
}

func (s *paymentService) HandleDeleteRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	if ctx.GetString("username") != "" && ctx.GetString("user_id") != "" {
		userData["user"] = ctx.Value("username").(string)
		userData["user_id"] = ctx.Value("user_id").(string)
	}
	var res []byte

	_, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        SetResult(&res).
        Delete("http://localhost:7000" + ctx.Request.URL.String())

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService")
        return
    }
	ctx.JSON(http.StatusOK, &res)
}