package middleware

import (
    "net/http"
)

func RoleMiddleware(requiredRole string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // In production: extract from JWT or session
            userRole := r.Header.Get("X-User-Role")
            if userRole != requiredRole {
                http.Error(w, "Forbidden", http.StatusForbidden)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
