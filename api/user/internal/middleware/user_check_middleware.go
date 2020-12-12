package middleware

import (
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"
)

type UserCheckMiddleware struct {
	log logx.Logger
}

func NewUserCheckMiddleware() *UserCheckMiddleware {
	return &UserCheckMiddleware{}
}

func (m *UserCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.log = logx.WithContext(r.Context())
		m.log.Info("UserCheckMiddleware")
		// Passthrough to next handler if need
		next(w, r)
	}
}
