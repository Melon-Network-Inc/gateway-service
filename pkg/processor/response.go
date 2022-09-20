package processor

import (
	"github.com/Melon-Network-Inc/common/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func HandleResponse(ctx *gin.Context, resp *resty.Response, err error, logger log.Logger) {
	if err != nil {
		HandleServiceUnavailable(ctx, err, logger)
		return
	}
	ctx.Data(resp.StatusCode(), "application/json", resp.Body())
}

func HandleServiceUnavailable(ctx *gin.Context, err error, logger log.Logger) {
	logger.Errorf("unable to connect to backend service: ", err)
	ctx.Status(http.StatusServiceUnavailable)
}
