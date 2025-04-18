package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// Define a type for context keys to avoid collisions
type contextKey string

const (
	UserIDKey = contextKey("id") // Key for storing user ID in context
)

// JWTAuth middleware function to handle JWT authentication and extract user ID
func JWTAuth(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			// Trim "Bearer " prefix from the token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, "Bearer token required", http.StatusUnauthorized)
				return
			}

			// Parse and validate the authorization token
			token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(jwtSecret), nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
				return
			}

			// Extract user ID from token claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}
			userID, ok := claims["sub"].(string)
			if !ok {
				http.Error(w, "User ID not found in token", http.StatusUnauthorized)
				return
			}

			// Add userID to request context
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
