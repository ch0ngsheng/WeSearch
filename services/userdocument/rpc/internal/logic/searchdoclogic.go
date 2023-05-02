package logic

import (
	"context"

	"chongsheng.art/wesearch/services/userdocument/rpc/internal/svc"
	"chongsheng.art/wesearch/services/userdocument/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchDocLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchDocLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchDocLogic {
	return &SearchDocLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchDocLogic) SearchDoc(in *pb.DocumentSearchReq) (*pb.DocumentSearchResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DocumentSearchResp{}, nil
}
