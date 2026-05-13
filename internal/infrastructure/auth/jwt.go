package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	GenerateToken(userID string) (string, error)
	ParseToken(tokenString string) (string, error)
}

type JWTManager struct {
	secretKey string
	tokenTTL  time.Duration
}

func NewJWTManager(secretKey string, tokenTTL time.Duration) *JWTManager {
	return &JWTManager{
		secretKey: secretKey,
		tokenTTL:  tokenTTL,
	}
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId string
}

func (j *JWTManager) GenerateToken(userID string) (string, error) {
	claims := tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: userID,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *JWTManager) ParseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.UserId, nil
}
