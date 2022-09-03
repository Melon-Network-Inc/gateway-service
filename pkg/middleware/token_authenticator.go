package middleware

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Melon-Network-Inc/gateway-service/pkg/storage"
	"github.com/gin-gonic/gin"
)

const (
	UsernameKey      = "username"
	UserIDKey        = "user_id"
	AuthorizationKey = "Authorization"
	RegistrationKey  = "RegistrationSession"
)

// TokenForwarder check if token is valid and forward token to backend service.
func TokenForwarder() gin.HandlerFunc {
	return func(context *gin.Context) {
		registrationTokenString := context.Request.Header.Get(RegistrationKey)
		if registrationTokenString == "" {
			context.AbortWithError(http.StatusUnauthorized, errors.New("registration is timed out"))
			return
		}
		context.Set(RegistrationKey, registrationTokenString)
		context.Next()
	}
}

// TokenAuthenticator check if token is valid and sets context key and value
// appropriately. Aborts when there was a problem in validating token.
func TokenAuthenticator(cache storage.Accessor) gin.HandlerFunc {
	return func(context *gin.Context) {
		fullAuthTokenString := context.Request.Header.Get(AuthorizationKey)
		if fullAuthTokenString != "" {
			context.Set(AuthorizationKey, fullAuthTokenString)
		}
		tokenString := strings.TrimPrefix(fullAuthTokenString, "Bearer ")

		user, err := cache.GetUser(context.Request.Context(), tokenString)
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		context.Set(UserIDKey, strconv.FormatUint(uint64(user.ID), 10))
		context.Set(UsernameKey, user.Username)
		context.Next()
	}
}
