package jwt

import (
	"errors"
	"time"
	"training/config"
	"training/models"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(
	cfg config.Config,
	userID int64,
	companyID int64,
	email string,
	role string,
) (string, error) {
	if cfg.JWTSecret == "" {
		return "", errors.New("JWT secret is not configured")
	}
	now := time.Now()

	claims := models.JWTClaims{
		UserID:    userID,
		CompanyID: companyID,
		Email:     email,
		Role:      role,

		RegisteredClaims: jwtlib.RegisteredClaims{
			IssuedAt: jwtlib.NewNumericDate(now),
			ExpiresAt: jwtlib.NewNumericDate(
				now.Add(cfg.JWTExpiry),
			),
		},
	}
	token := jwtlib.NewWithClaims(
		jwtlib.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString(
		[]byte(cfg.JWTSecret),
	)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
