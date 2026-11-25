package middleware

import (
	"context"
	"encoding/json"
	jwtlib "github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"training/models"
	reqctx "training/request"
)

type contextKey string

const claimsContextKey contextKey = "user_claim"

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				writeError(w, http.StatusUnauthorized, "authorization header required")
				return
			}
			parts := strings.Fields(authHeader)

			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				writeError(w, http.StatusUnauthorized, "invalid authorization format")
				return
			}
			tokenString := parts[1]

			claims := &models.JWTClaims{}

			token, err := jwtlib.ParseWithClaims(
				tokenString,
				claims,
				func(token *jwtlib.Token) (interface{}, error) {
					return []byte(secret), nil
				},
				jwtlib.WithValidMethods([]string{"HS256"}),
			)
			if err != nil || !token.Valid {
				writeError(w, http.StatusUnauthorized, "invalid or expired token")
				return
			}

			ctx := context.WithValue(
				r.Context(),
				claimsContextKey,
				claims,
			)

			// Validate authenticated identity.
			if claims.UserID <= 0 || claims.CompanyID <= 0 {
				writeError(
					w,
					http.StatusUnauthorized,
					"invalid user or company information",
				)
				return
			}

			// Store user and company information.
			ctx = reqctx.WithIdentity(
				ctx,
				reqctx.Identity{
					UserID:    claims.UserID,
					CompanyID: claims.CompanyID,
					Role:      claims.Role,
				},
			)

			// Pass updated context.
			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)

		})
	}
}

func GetClaims(ctx context.Context) (*models.JWTClaims, bool) {
	claims, ok := ctx.Value(claimsContextKey).(*models.JWTClaims)
	return claims, ok
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
