package utils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

var secretKey = []byte("secret-key")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]
		fmt.Println("near token string", tokenString)
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		fmt.Println("token valid ", token)

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		fmt.Println("berore route func")
		role := (*claims)["role"].(string)

		// Print the role to the console
		fmt.Println("User Role:", role)
		r = r.WithContext(context.WithValue(r.Context(), "username", (*claims)["username"]))
		r = r.WithContext(context.WithValue(r.Context(), "role", (*claims)["role"]))

		next.ServeHTTP(w, r)
	})
}
func RoleAuthorizationMiddleware(role string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole, ok := r.Context().Value("role").(string)
			fmt.Println("user role is ", userRole)
			if !ok || userRole == "" {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			if userRole != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
