package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	accessSecret    = []byte("your-access-secret")
	refreshSecret   = []byte("your-refresh-secret")
	ErrInvalidToken = errors.New("invalid token")
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

func ParseToken(tokenStr string, isAccess bool) (string, error) {
	secret := refreshSecret
	if isAccess {
		secret = accessSecret
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return "", ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		return "", ErrInvalidToken
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", ErrInvalidToken
	}

	return userID, nil
}
