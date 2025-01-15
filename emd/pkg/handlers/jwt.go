package handlers

import (
	"context"
	"emb/pkg/auth"
	"emb/pkg/tmpl"
	"net/http"
)

func JwtAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			tmpl.Tmpl.ExecuteTemplate(w, "Login", nil)
			return
		}

		isTokenValid, username := auth.VerifyToken(c.Value)
		if !isTokenValid {
			tmpl.Tmpl.ExecuteTemplate(w, "Login", nil)
			return
		}

		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func JwtAuthJsonMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		isTokenValid, username := auth.VerifyToken(c.Value)
		if !isTokenValid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
