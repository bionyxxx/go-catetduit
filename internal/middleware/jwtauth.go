package middleware

import (
	"catetduit/internal/helper"
	"context"
	"net/http"
)

type contextKey string

const UserClaimsKey contextKey = "userClaims"

type AuthMiddleware struct {
	jwtHelper *helper.JWTHelper
}

func NewAuthMiddleware(jwtHelper *helper.JWTHelper) *AuthMiddleware {
	return &AuthMiddleware{
		jwtHelper: jwtHelper,
	}
}

func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString string
		var err error

		// Prioritas 1: Baca dari cookie (production/browser)
		cookie, errCookie := r.Cookie("access_token")
		if errCookie == nil {
			tokenString = cookie.Value
		} else {
			// Prioritas 2: Baca dari Authorization header (testing/API client)
			authHeader := r.Header.Get("Authorization")
			tokenString, err = helper.ExtractTokenFromHeader(authHeader)
			if err != nil {
				err := helper.ResponseBadRequest(w, "Authorization token required", err)
				if err != nil {
					panic("Failed to send unauthorized response: " + err.Error())
				}
				return
			}
		}

		claims, err := m.jwtHelper.ValidateToken(tokenString)
		if err != nil {
			err = helper.ResponseUnauthorized(w, "Invalid or expired token")
			if err != nil {
				panic("Failed to send unauthorized response: " + err.Error())
			}
			return
		}

		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
