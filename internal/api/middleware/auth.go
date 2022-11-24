package middleware

import (
	"context"
	"net/http"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/internal/api/dto"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/users"
)

func NewAuth(authSvc auth.UseCase, ignoredEndpoints ...string) Middleware {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		for _, endpoint := range ignoredEndpoints {
			if r.URL.Path == endpoint {
				next(w, r)

				return
			}
		}

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write(dto.MessageJSON{Message: "authorization not found"}.Marshal())

			return
		}

		token := authHeader[len("Bearer")+1:]

		user, err := authSvc.Authorize(r.Context(), token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write(dto.MessageJSON{Message: err.Error()}.Marshal())
		}

		w.Header().Set("x-user-id", user.ID.String())

		ctx := context.WithValue(r.Context(), users.CtxKey{}, &user)

		r = r.WithContext(ctx)

		next(w, r)
	}
}
