package middleware

import (
	"buffer/config"
	"fmt"
	"net/http"
)

// Обёртка авторизации прокси
type Auth struct {
	handler http.Handler
}

func (a *Auth) Authorize(token string) bool {
	return token == fmt.Sprintf("Bearer %s", config.GlobalConfig.BearerToken)
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	isAuth := a.Authorize(token)
	if !isAuth {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	a.handler.ServeHTTP(w, r)
}

func NewAuth(wrappedHandler http.Handler) *Auth {
	return &Auth{wrappedHandler}
}
