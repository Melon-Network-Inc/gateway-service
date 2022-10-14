package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/Melon-Network-Inc/common/pkg/config"
	"github.com/Melon-Network-Inc/common/pkg/log"
	"github.com/golang-jwt/jwt/v4"
)

type tokenManager struct {
	secretKey                  []byte
	logger                     log.Logger
}

type HashedTokenManager interface {
	GetSecretKey() []byte
	ValidateAuthToken(tokenString string) (int64, string, error)
}

func NewHashedTokenManager(tokenConfig *config.TokenConfig, logger log.Logger) HashedTokenManager {
	return &tokenManager{
		secretKey:                  tokenConfig.GetSecretKey(),
		logger:                     logger,
	}
}

func (m *tokenManager) GetSecretKey() []byte {
	return m.secretKey
}

func (m *tokenManager) ValidateAuthToken(tokenString string) (int64, string, error) {
	// set up a parser that doesn't validate expiration time
	parser := jwt.NewParser()

	token, err := parser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.secretKey, nil
	})

	if err != nil {
		m.logger.Error("Failed to parse the token [" + tokenString + "].")
		return 0, "", errors.New("invalid authentication token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimUserID, isSetID := claims["userID"]
		userID, ok := claimUserID.(float64)
		if !ok || !isSetID {
			return 0, "", errors.New("token does not contain required data")
		}

		claimLogin, isSetLogin := claims["lastLogin"]
		login, ok := claimLogin.(string)
		if !ok || !isSetLogin {
			return 0, "", errors.New("token does not contain required data")
		}

		// check if token contains expiry date
		if unexpired := claims.VerifyExpiresAt(time.Now().Unix(), true); !unexpired {
			return 0, "", errors.New("token has expired")
		}

		return int64(userID), login, nil
	}

	return 0, "", errors.New("malformed authentication token")
}