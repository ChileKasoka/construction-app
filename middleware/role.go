package middleware

import (
	"net/http"
	"strings"

	"slices"

	auth "github.com/ChileKasoka/construction-app/middleware/auth"
)

// Accepts one or more roles
func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := auth.ParseJWT(tokenString)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			roleClaim := claims["role"].(string)
			authorized := slices.Contains(allowedRoles, roleClaim)

			if !authorized {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			// Optional: inject user ID or role into context for later use
			next.ServeHTTP(w, r)
		})
	}
}
