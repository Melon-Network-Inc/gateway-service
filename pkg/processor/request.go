package processor

import (
	"errors"
	"fmt"

	"github.com/Melon-Network-Inc/common/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	UsernameKey = "username"
	UserIDKey   = "user_id"

	ContextUserKey              = "Username"
	ContextUserIDKey            = "UserID"
	AuthorizationKey            = "Authorization"
	RegistrationKey             = "RegistrationSession"
	ContextRoleKey              = "UserRole"
	ContextRegistrationTokenKey = "RegistrationSessionToken"
	ContextClientIP             = "X-Forwarded-For"
	ContextRequestID            = "X-Request-ID"
	ContextCorrelationID        = "X-Correlation-ID"
)

func PrepareRequest(ctx *gin.Context, client *resty.Client) *resty.Request {
	var req *resty.Request
	userData, exists := GetUserData(ctx)

	if exists {
		req = client.R().
			SetBody(ctx.Request.Body).
			SetHeaders(userData)
	} else {
		req = client.R().
			SetBody(ctx.Request.Body)
	}

	return req
}

func PrepareRequestWithAttachment(ctx *gin.Context, client *resty.Client) (*resty.Request, error) {
	var req *resty.Request
	userData, exists := GetUserData(ctx)

	if exists {
		req = client.R().
			SetBody(ctx.Request.Body).
			SetHeaders(userData)
	} else {
		req = client.R().
			SetBody(ctx.Request.Body)
	}

	// Check if request contains file. If not, skip the multipart form.
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		return req, errors.New(fmt.Sprintf("%s-%s-%s", file, fileHeader.Filename, err))
	}
	req.SetMultipartField(
		"file",
		fileHeader.Filename,
		ctx.Request.Header.Get("Content-Type"),
		file,
	)

	return req, nil
}

func GetUserData(ctx *gin.Context) (map[string]string, bool) {
	userData := make(map[string]string)
	username, existsName := ctx.Get(UsernameKey)
	userID, existsID := ctx.Get(UserIDKey)
	token, existsToken := ctx.Get(AuthorizationKey)
	if existsName && existsID && existsToken {
		userData[ContextUserKey] = username.(string)
		userData[ContextUserIDKey] = fmt.Sprintf("%d", userID.(uint))
		userData[AuthorizationKey] = token.(string)
	}
	userRole, existsRole := ctx.Get(ContextRoleKey)
	if existsRole {
		userData[ContextRoleKey] = userRole.(string)
	} else {
		userData[ContextRoleKey] = fmt.Sprintf("%d", entity.UserRole)
	}

	registrationToken, existsRegistrationToken := ctx.Get(RegistrationKey)
	if existsRegistrationToken {
		userData[ContextRegistrationTokenKey] = registrationToken.(string)
	}

	userData[ContextClientIP] = ctx.ClientIP()
	userData[ContextRequestID] = ctx.GetString(ContextRequestID)
	if correlationID, existsID := ctx.Get(ContextCorrelationID); existsID {
		userData[ContextCorrelationID] = correlationID.(string)
	}
	return userData, (existsName && existsID && existsToken) || existsRegistrationToken
}
