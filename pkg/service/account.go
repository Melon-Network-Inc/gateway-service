package service

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type AccountService interface {
	HandleGetRequest(ctx *gin.Context)
	HandlePostRequest(ctx *gin.Context)
	HandleUpdateRequest(ctx *gin.Context)
	HandleDeleteRequest(ctx *gin.Context)
}

type accountService struct {}

func NewAccountService() AccountService {
	return &accountService{}
}

func (s *accountService) HandleGetRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	if ctx.GetString("username") != "" && ctx.GetString("user_id") != "" {
		userData["user"] = ctx.Value("username").(string)
		userData["user_id"] = ctx.Value("user_id").(string)
	}

	resp, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        Get("http://localhost:6000" + ctx.Request.URL.String())

    if err != nil {
    	log.Println("Account Service: unable to connect AccountService due to", err)
        return
    }
	ctx.JSON(resp.StatusCode(), resp.Body())
}

func (s *accountService) HandlePostRequest(ctx *gin.Context) {
	client := resty.New()

	userData := make(map[string]string)
	if ctx.GetString("username") != "" && ctx.GetString("user_id") != "" {
		userData["user"] = ctx.Value("username").(string)
		userData["user_id"] = ctx.Value("user_id").(string)
	}

	resp, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        Post("http://localhost:6000" + ctx.Request.URL.String())

    if err != nil {
    	log.Println("Account Service: unable to connect AccountService due to", err)
        return
    }
	ctx.JSON(resp.StatusCode(), resp.Body())
}

func (s *accountService) HandleUpdateRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	if ctx.GetString("username") != "" && ctx.GetString("user_id") != "" {
		userData["user"] = ctx.Value("username").(string)
		userData["user_id"] = ctx.Value("user_id").(string)
	}
	resp, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        Put("http://localhost:6000" + ctx.Request.URL.String())

    if err != nil {
    	log.Println("Account Service: unable to connect AccountService due to", err)
        return
    }
	ctx.JSON(resp.StatusCode(), resp.Body())
}

func (s *accountService) HandleDeleteRequest(ctx *gin.Context) {
	client := resty.New()
	userData := make(map[string]string)
	userData["user"] = ctx.Value("username").(string)
	userData["user_id"] = ctx.Value("user_id").(string)

	resp, err := client.R().
        SetBody(ctx.Request.Body).
		SetHeaders(userData).
        Delete("http://localhost:6000" + ctx.Request.URL.String())

    if err != nil {
    	log.Println("Account Service: unable to connect AccountService due to", err)
        return
    }
	ctx.JSON(resp.StatusCode(), resp.Body())
}