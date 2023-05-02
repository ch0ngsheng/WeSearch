package wx

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"chongsheng.art/wesearch/services/wxmanager/api/internal/logic/wx"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/svc"
)

func VersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := wx.NewVersionLogic(r.Context(), svcCtx)
		resp, err := l.Version()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
