package wx

import (
	"net/http"

	"chongsheng.art/wesearch/services/wxmanager/api/internal/logic/wx"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/svc"
)

func WeChatValidationWithRawHTTPHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// generated by custom templates
		l := wx.NewWeChatValidationWithRawHTTPLogic(r.Context(), svcCtx)
		l.WeChatValidationWithRawHTTP(w, r)
	}
}
