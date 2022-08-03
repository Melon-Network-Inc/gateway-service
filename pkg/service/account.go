package service

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type AccountService interface {
	HandleRequest(ctx *gin.Context)
}

type accountService struct {}

func NewAccountService() AccountService {
	return &accountService{}
}

func (s *accountService) HandleRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	userData["user"] = ctx.Value("username").(string)
	userData["user_id"] = ctx.Value("user_id").(string)
	var res []byte

	_, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        SetResult(&res).
        Post("http://localhost:6000/account")

    if err != nil {
    	log.Println("Account Service: unable to connect AccountService")
        return
    }
	ctx.JSON(http.StatusOK, &res)
}