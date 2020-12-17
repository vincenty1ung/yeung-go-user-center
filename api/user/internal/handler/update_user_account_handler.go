package handler

import (
	"net/http"

	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/logic"
	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/svc"
	"github.com/uncleyeung/yeung-go-user-center/api/user/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UpdateUserAccountHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAccountUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateUserAccountLogic(r.Context(), ctx)
		err := l.UpdateUserAccount(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
