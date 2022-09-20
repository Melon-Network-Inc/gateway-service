package middleware

import (
	"net/http"
	"strings"

	"github.com/Melon-Network-Inc/gateway-service/pkg/storage"
	"github.com/gin-gonic/gin"
)

const (
	UsernameKey      = "username"
	UserIDKey        = "user_id"
	UserRoleKey      = "role"
	AuthorizationKey = "Authorization"
	RegistrationKey  = "RegistrationSession"
)

// TokenForwarder check if token is valid and forward token to backend service.
func TokenForwarder() gin.HandlerFunc {
	return func(context *gin.Context) {
		registrationTokenString := context.Request.Header.Get(RegistrationKey)
		if registrationTokenString == "" {
			context.String(http.StatusUnauthorized, "registration is timed out")
			return
		}
		context.Set(RegistrationKey, registrationTokenString)
		context.Next()
	}
}

// TokenAuthenticator check if token is valid and sets context key and value
// appropriately. Aborts when there was a problem in validating token.
func TokenAuthenticator(cache storage.Accessor) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fullAuthTokenString := ctx.Request.Header.Get(AuthorizationKey)
		if fullAuthTokenString != "" {
			ctx.Set(AuthorizationKey, fullAuthTokenString)
		}
		tokenString := strings.TrimPrefix(fullAuthTokenString, "Bearer ")

		user, err := cache.GetCachedUserByToken(ctx, tokenString)
		if err != nil {
			ctx.String(http.StatusUnauthorized, err.Error())
			return
		}

		ctx.Set(UserIDKey, user.ID)
		ctx.Set(UsernameKey, user.Username)
		ctx.Set(UserRoleKey, user.Role)
		ctx.Next()
	}
}
