package service

import (
	"net/http"

	"github.com/Melon-Network-Inc/common/pkg/log"
	"github.com/Melon-Network-Inc/gateway-service/pkg/processor"

	"github.com/gin-gonic/gin"
)

type AccountService interface {
	HandleGetRequest(ctx *gin.Context)
	HandlePostRequest(ctx *gin.Context)
	HandleUpdateRequest(ctx *gin.Context)
	HandleDeleteRequest(ctx *gin.Context)
	HandleServiceUnavailable(ctx *gin.Context, err error)
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
	client := CreateRetryRestyClient()
	resp, err := processor.PrepareRequest(ctx, client).
		Get(s.serviceUrlPrefix + ctx.Request.URL.String())
	if err != nil {
		s.HandleServiceUnavailable(ctx, err)
		return
	}
	
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *accountService) HandlePostRequest(ctx *gin.Context) {
	client := CreateRetryRestyClient()
	resp, err := processor.PrepareRequest(ctx, client).
		Post(s.serviceUrlPrefix + ctx.Request.URL.String())
	if err != nil {
		s.HandleServiceUnavailable(ctx, err)
		return
	}
	
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *accountService) HandleUpdateRequest(ctx *gin.Context) {
	client := CreateRetryRestyClient()
	resp, err := processor.PrepareRequest(ctx, client).
		Put(s.serviceUrlPrefix + ctx.Request.URL.String())
	if err != nil {
		s.HandleServiceUnavailable(ctx, err)
		return
	}

	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *accountService) HandleDeleteRequest(ctx *gin.Context) {
	client := CreateRetryRestyClient()
	resp, err := processor.PrepareRequest(ctx, client).
		Delete(s.serviceUrlPrefix + ctx.Request.URL.String())
    if err != nil {
    	s.HandleServiceUnavailable(ctx, err)
        return
    }

	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func (s *accountService) HandleServiceUnavailable(ctx *gin.Context, err error) {
	s.logger.Errorf("Account Service: unable to connect AccountService due to", err)
	ctx.Status(http.StatusServiceUnavailable)
}