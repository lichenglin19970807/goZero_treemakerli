package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goZero_treemakerli/user-api/internal/logic"
	"goZero_treemakerli/user-api/internal/svc"
	"goZero_treemakerli/user-api/internal/types"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
