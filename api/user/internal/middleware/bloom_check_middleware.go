package middleware

import (
	"errors"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"github.com/uncleyeung/yeung-user-center/api/user/internal/bloom"
	"net/http"
	"strconv"
	"strings"
)

type BloomCheckMiddleware struct {
	log     logx.Logger
	bloomer bloom.BloomerUserAbility
}

func NewBloomCheckMiddleware(cfg bloom.BloomerUserAbility) *BloomCheckMiddleware {
	return &BloomCheckMiddleware{bloomer: cfg}
}

func (m *BloomCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.log = logx.WithContext(r.Context())
		m.log.Info("BloomCheckMiddleware")
		uri := r.RequestURI
		if strings.Contains(uri, "/id/") {
			var req reqId
			if err := httpx.Parse(r, &req); err != nil {
				httpx.Error(w, err)
				return
			}
			exists, err := m.bloomer.Exists(strconv.FormatInt(req.Id, 10))
			if err != nil {
				httpx.Error(w, err)
				return
			}
			if exists {
				httpx.Error(w, errors.New("布隆过滤异常"))
				return
			}
		}

		next(w, r)
	}
}

//简单实现判断id
type reqId struct {
	Id int64 `path:"id"`
}
