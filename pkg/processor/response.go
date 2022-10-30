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
	ctx.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}

func HandleServiceUnavailable(ctx *gin.Context, err error, logger log.Logger) {
	logger.Errorf("unable to connect to backend service: ", err)
	ctx.Status(http.StatusServiceUnavailable)
}

func HandleFileAttachmentNotFound(ctx *gin.Context, err error, logger log.Logger) {
	logger.Errorf("unable to fetch file attachment: ", err)
	ctx.String(http.StatusBadRequest, "unable to fetch file attachment: ", err)
}
