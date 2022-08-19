package service

import (
	"github.com/Melon-Network-Inc/common/pkg/log"
	"github.com/Melon-Network-Inc/gateway-service/pkg/processor"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type AccountService interface {
	HandleGetRequest(ctx *gin.Context)
	HandlePostRequest(ctx *gin.Context)
	HandleUpdateRequest(ctx *gin.Context)
	HandleDeleteRequest(ctx *gin.Context)
}

type accountService struct {
	serviceUrlPrefix string
	logger log.Logger
}

func NewAccountService(serviceUrlPrefix string, logger log.Logger) AccountService {
	return &accountService{
		serviceUrlPrefix: serviceUrlPrefix,
		logger: logger,
	}
}

func (s *accountService) HandleGetRequest(ctx *gin.Context) {
	client := resty.New()
	resp, err := processor.PrepareRequest(ctx, client).
		Get(s.serviceUrlPrefix + ctx.Request.URL.String())

    if err != nil {
    	s.logger.Errorf("Account Service: unable to connect AccountService due to", err)
        return
    }
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *accountService) HandlePostRequest(ctx *gin.Context) {
	client := resty.New()
	resp, err := processor.PrepareRequest(ctx, client).
		Post(s.serviceUrlPrefix + ctx.Request.URL.String())

    if err != nil {
    	s.logger.Errorf("Account Service: unable to connect AccountService due to", err)
        return
    }
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *accountService) HandleUpdateRequest(ctx *gin.Context) {
	client := resty.New()
	resp, err := processor.PrepareRequest(ctx, client).
		Put(s.serviceUrlPrefix + ctx.Request.URL.String())

    if err != nil {
    	s.logger.Errorf("Account Service: unable to connect AccountService due to", err)
        return
    }
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *accountService) HandleDeleteRequest(ctx *gin.Context) {
	client := resty.New()
	resp, err := processor.PrepareRequest(ctx, client).
		Delete(s.serviceUrlPrefix + ctx.Request.URL.String())

    if err != nil {
    	s.logger.Errorf("Account Service: unable to connect AccountService due to", err)
        return
    }
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}