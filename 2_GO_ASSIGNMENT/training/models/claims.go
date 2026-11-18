package models

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID    int64  `json:"user_id"`
	CompanyID int64  `json:"company_id"`
	Email     string `json:"email"`

	jwt.RegisteredClaims
}
