package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"basico-crud-go/infra/keycloack"
)

type UserClaims struct {
	Subject           string `json:"sub"`
	Email             string `json:"email"`
	PreferredUsername string `json:"preferred_username"`
	Name              string `json:"name"`
}

type contextKey string

const userContextKey contextKey = "user"

func Auth(keycloak *auth.Keycloak) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(
					w,
					"Token não informado",
					http.StatusUnauthorized,
				)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)

			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(
					w,
					"Formato de token inválido",
					http.StatusUnauthorized,
				)
				return
			}

			rawToken := parts[1]

			token, err := keycloak.Verifier.Verify(
				r.Context(),
				rawToken,
			)

			if err != nil {
				http.Error(
					w,
					"Token inválido ou expirado",
					http.StatusUnauthorized,
				)
				return
			}

			var claims UserClaims

			if err := token.Claims(&claims); err != nil {
				http.Error(
					w,
					"Erro ao ler informações do token",
					http.StatusUnauthorized,
				)
				return
			}

			ctx := context.WithValue(
				r.Context(),
				userContextKey,
				claims,
			)

			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)
		})
	}
}

func GetUser(r *http.Request) (UserClaims, bool) {

	user, ok := r.Context().Value(userContextKey).(UserClaims)

	return user, ok
}

func Me(w http.ResponseWriter, r *http.Request) {

	user, ok := GetUser(r)

	if !ok {
		http.Error(
			w,
			"Usuário não encontrado",
			http.StatusUnauthorized,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}