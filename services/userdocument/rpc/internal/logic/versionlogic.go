package logic

import (
	"chongsheng.art/wesearch/services/retrieve/rpc/retrieve"
	"context"
	"log"

	"chongsheng.art/wesearch/services/userdocument/rpc/internal/svc"
	"chongsheng.art/wesearch/services/userdocument/rpc/pb"

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
		log.Printf("failed to request retriever, err: %v", err)
		return nil, err
	}

	return &pb.VersionResp{
		Version: resp.Version,
	}, nil
}
