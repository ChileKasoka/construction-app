package middleware

import (
	"fmt"
	"net/http"
	"strings"

	auth "github.com/ChileKasoka/construction-app/middleware/auth"
	"github.com/ChileKasoka/construction-app/repository"
	"github.com/go-chi/chi/v5"
)

func RoleMiddleware(repo *repository.RolePermissionRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			fmt.Println("Authorization Header:", authHeader)
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := auth.ParseJWT(tokenString)
			if err != nil {
				fmt.Println("JWT parsing error:", err)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			fmt.Printf("Claims: %+v\n", claims)

			roleFloat, ok := claims["role_id"].(float64)
			if !ok {
				http.Error(w, "role_id missing in token", http.StatusUnauthorized)
				return
			}

			roleID := int(roleFloat)
			if !ok {
				http.Error(w, "role_id must be a number", http.StatusUnauthorized)
				return
			}

			routeCtx := chi.RouteContext(r.Context())
			pattern := routeCtx.RoutePattern()

			fmt.Printf("Route pattern: %s, Method: %s, Role ID: %d\n", pattern, r.Method, roleID)

			allowed, err := repo.HasPermission(roleID, pattern)
			if err != nil {
				http.Error(w, "Error checking permissions", http.StatusInternalServerError)
				return
			}

			if !allowed {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
