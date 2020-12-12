package middleware

import (
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
)

type LoginCheckMiddleware struct {
	log logx.Logger
}

func NewLoginCheckMiddleware() *LoginCheckMiddleware {
	return &LoginCheckMiddleware{}
}

func (m *LoginCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.log = logx.WithContext(r.Context())
		m.log.Info("LoginCheckMiddleware")
		// Passthrough to next handler if need
		next(w, r)
	}
}
