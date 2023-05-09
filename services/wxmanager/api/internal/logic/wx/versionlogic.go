package wx

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"chongsheng.art/wesearch/services/userdocument/rpc/userdocument"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/svc"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/types"
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
	res, err := l.svcCtx.UserDocRpc.Version(l.ctx, &userdocument.VersionReq{})
	if err != nil {
		logx.Errorf("rpc request Version, %+v", err)
		return nil, errors.Wrap(err, "rpc user doc.")
	}
	resp = &types.VersionResp{Version: res.Version}
	return resp, nil
}
