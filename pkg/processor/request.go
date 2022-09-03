package processor

import (
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
	return userData, (existsName && existsID && existsToken) || existsRegistrationToken
}
