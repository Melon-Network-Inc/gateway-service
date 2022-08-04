package service

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type PaymentService interface {
	HandleRequest(ctx *gin.Context)
}

type paymentService struct {}

func NewPaymentService() PaymentService {
	return &paymentService{}
}

func (s *paymentService) HandleRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	userData["user"] = ctx.Value("username").(string)
	userData["user_id"] = ctx.Value("user_id").(string)
	var res []byte

	_, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        SetResult(&res).
        Post("http://localhost:7000/payment")

    if err != nil {
        log.Println("Payment Service: unable to connect PaymentService")
        return
    }
	ctx.JSON(http.StatusOK, &res)
}