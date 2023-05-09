package logic

import (
	"chongsheng.art/wesearch/services/retrieve/rpc/retrieve"
	"chongsheng.art/wesearch/services/userdocument/rpc/internal/svc"
	"chongsheng.art/wesearch/services/userdocument/rpc/pb"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type VersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VersionLogic {
	return &VersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VersionLogic) Version(in *pb.VersionReq) (*pb.VersionResp, error) {
	resp, err := l.svcCtx.RetrieveRpc.Version(l.ctx, &retrieve.VersionReq{})
	if err != nil {
		logx.Errorf("rpc retrieve get version, %+v", err)
		return nil, errors.Wrap(err, "rpc retrieve get version.")
	}

	return &pb.VersionResp{
		Version: resp.Version,
	}, nil
}
