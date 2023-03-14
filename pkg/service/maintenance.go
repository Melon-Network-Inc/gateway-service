package service

import (
	"github.com/Melon-Network-Inc/common/pkg/log"
	"github.com/Melon-Network-Inc/gateway-service/pkg/processor"

	"github.com/gin-gonic/gin"
)

type MaintenanceService interface {
	HandlePostRequest(ctx *gin.Context)
	HandleUpdateRequest(ctx *gin.Context)
	HandleGetRequest(ctx *gin.Context)
	HandleDeleteRequest(ctx *gin.Context)
}

type maintenanceService struct {
	serviceUrlPrefix string
	logger           log.Logger
}

func NewMaintenanceService(serviceUrlPrefix string, logger log.Logger) MaintenanceService {
	return &maintenanceService{
		serviceUrlPrefix: serviceUrlPrefix,
		logger:           logger,
	}
}

func (s *maintenanceService) HandleGetRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Get(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}

func (s *maintenanceService) HandlePostRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Post(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}

func (s *maintenanceService) HandleUpdateRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Put(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}

func (s *maintenanceService) HandleDeleteRequest(ctx *gin.Context) {
	resp, err := processor.PrepareRequest(ctx, CreateRetryRestyClient()).
		Delete(s.serviceUrlPrefix + ctx.Request.URL.String())
	processor.HandleResponse(ctx, resp, err, s.logger)
}
