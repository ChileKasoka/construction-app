package middleware

import (
	"fmt"
	"net/http"
	"strings"

	auth "github.com/ChileKasoka/construction-app/middleware/auth"
	"github.com/ChileKasoka/construction-app/model"
	"github.com/ChileKasoka/construction-app/repository"
)

// matchPathFromPermissions finds the matching permission path for the request
func matchPathFromPermissions(requestPath, requestMethod string, permissions []model.Permission) (string, bool) {
	reqSegments := strings.Split(strings.Trim(requestPath, "/"), "/")

	for _, perm := range permissions {
		if !strings.EqualFold(perm.Method, requestMethod) {
			continue
		}

		permSegments := strings.Split(strings.Trim(perm.Path, "/"), "/")

		if len(reqSegments) != len(permSegments) {
			continue
		}

		match := true
		for i := range permSegments {
			if strings.HasPrefix(permSegments[i], ":") || strings.HasPrefix(permSegments[i], "{") {
				// This is a path variable, match anything
				continue
			}
			if permSegments[i] != reqSegments[i] {
				match = false
				break
			}
		}

		if match {
			return perm.Path, true
		}
	}

	return "", false
}

func RoleMiddleware(repo *repository.RolePermissionRepo) func(http.Handler) http.Handler {
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

			roleFloat, ok := claims["role_id"].(float64)
			if !ok {
				http.Error(w, "role_id missing in token", http.StatusUnauthorized)
				return
			}
			roleID := int(roleFloat)

			method := r.Method
			requestPath := r.URL.Path

			permissions, err := repo.GetByRoleID(roleID)
			if err != nil {
				http.Error(w, "Error fetching permissions for role", http.StatusInternalServerError)
				return
			}
			if len(permissions) == 0 {
				http.Error(w, "No permissions found for role", http.StatusForbidden)
				return
			}

			normalizedPath, found := matchPathFromPermissions(requestPath, method, permissions)
			if !found {
				http.Error(w, "Forbidden: route not allowed", http.StatusForbidden)
				return
			}

			allowed, err := repo.HasPermission(roleID, normalizedPath)
			if err != nil {
				http.Error(w, "Error checking permissions", http.StatusInternalServerError)
				return
			}
			if !allowed {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			fmt.Printf("✅ Access granted: %s [%s] matched %s for role %d\n",
				method, requestPath, normalizedPath, roleID)

			next.ServeHTTP(w, r)
		})
	}
}
