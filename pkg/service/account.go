package service

import (
	"github.com/Melon-Network-Inc/common/pkg/log"
	"github.com/Melon-Network-Inc/gateway-service/pkg/processor"
	"github.com/gin-gonic/gin"
)

type AccountService interface {
	HandleGetRequest(ctx *gin.Context)
	HandlePostRequestWithAttachment(ctx *gin.Context)
	HandlePostRequest(ctx *gin.Context)
	HandleUpdateRequest(ctx *gin.Context)
	HandleDeleteRequest(ctx *gin.Context)
}

type accountService struct {
	serviceUrlPrefix string
	logger           log.Logger
}

func NewAccountService(serviceUrlPrefix string, logger log.Logger) AccountService {
	return &accountService{
		serviceUrlPrefix: serviceUrlPrefix,
		logger:           logger,
	}
}

func (s *accountService) HandleGetRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Get(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}

func (s *accountService) HandlePostRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Post(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}

func (s *accountService) HandlePostRequestWithAttachment(ctx *gin.Context) {
	req, err := processor.PrepareRequestWithAttachment(ctx, CreateRetryRestyClient())
	if err != nil {
		processor.HandleFileAttachmentNotFound(ctx, err, s.logger)
		return
	}

	resp, err := req.Post(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}

func (s *accountService) HandleUpdateRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Put(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}

func (s *accountService) HandleDeleteRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Delete(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}
