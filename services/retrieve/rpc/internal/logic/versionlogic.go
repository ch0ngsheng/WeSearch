package logic

import (
	"context"

	"chongsheng.art/wesearch/services/retrieve/rpc/internal/svc"
	"chongsheng.art/wesearch/services/retrieve/rpc/pb"

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
	// todo: add your logic here and delete this line

	return &pb.VersionResp{}, nil
}
