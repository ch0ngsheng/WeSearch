package wx

import (
	"chongsheng.art/wesearch/services/wxmanager/api/internal/svc"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type VersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VersionLogic {
	return &VersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VersionLogic) Version() (resp *types.VersionResp, err error) {
	resp = &types.VersionResp{Version: "0.0.1"}
	return
}
