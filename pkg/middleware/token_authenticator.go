package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Melon-Network-Inc/gateway-service/pkg/storage"
	"github.com/gin-gonic/gin"
)

const (
	UsernameKey = "username"
	UserIDKey = "user_id"
	AuthorizationKey = "Authorization"
)

// TokenAuthenticator check if token is valid and sets context key and value
// appropiretly. Aborts when there was a problem in validating token.
func TokenAuthenticator(cache storage.Accessor) gin.HandlerFunc {
	return func(context *gin.Context) {
		fullTokenString := context.Request.Header.Get(AuthorizationKey)
		if fullTokenString != "" {
			context.Set(AuthorizationKey, fullTokenString)
		}
		tokenString := strings.TrimPrefix(fullTokenString, "Bearer ")

		user, err := cache.GetUser(context.Request.Context(), tokenString)
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		
		context.Set(UserIDKey,  strconv.FormatUint(uint64(user.ID), 10))
		context.Set(UsernameKey, user.Username)
		context.Next()
	}
}
