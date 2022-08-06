package processor

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const UsernameKey = "username"
const UserIDKey = "user_id"
const ContextUserKey = "Username"
const ContextUserIDKey = "User-Id"
const AuthorizationKey = "Authorization"

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
		userData[ContextUserIDKey] = userID.(string)
		userData[AuthorizationKey] = token.(string)
	}
	return userData, (existsName && existsID && existsToken)
}