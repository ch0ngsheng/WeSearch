package logic

import (
	"context"

	"chongsheng.art/wesearch/services/retrieve/rpc/internal/svc"
	"chongsheng.art/wesearch/services/retrieve/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLogic) Search(in *pb.SearchReq) (*pb.SearchResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchResp{}, nil
}
