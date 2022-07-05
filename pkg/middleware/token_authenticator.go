package middleware

import (
	"net/http"
	"strings"

	"github.com/Melon-Network-Inc/gateway-service/pkg/storage"
	"github.com/gin-gonic/gin"
)

// TokenAuthenticator check if token is valid and sets context key and value
// appropiretly. Aborts when there was a problem in validating token.
func TokenAuthenticator(cache storage.Accessor) gin.HandlerFunc {
	return func(context *gin.Context) {
		fullTokenString := context.Request.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(fullTokenString, "Bearer ")

		user, err := cache.GetUser(context.Request.Context(), tokenString)
		if err != nil {
			context.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		context.Set("user_id", user.ID)
		context.Set("username", user.Username)
		context.Next()
	}
}
