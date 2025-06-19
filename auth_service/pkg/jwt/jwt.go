package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	accessSecret  = []byte("your-access-secret")
	refreshSecret = []byte("your-refresh-secret")
)

func GenerateTokens(userID string) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	accessStr, err := accessToken.SignedString(accessSecret)
	if err != nil {
		return "", "", err
	}

	refreshStr, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessStr, refreshStr, nil
}
