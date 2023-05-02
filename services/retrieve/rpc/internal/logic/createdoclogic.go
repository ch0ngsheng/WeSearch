package logic

import (
	"context"

	"chongsheng.art/wesearch/services/retrieve/rpc/internal/svc"
	"chongsheng.art/wesearch/services/retrieve/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDocLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDocLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDocLogic {
	return &CreateDocLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDocLogic) CreateDoc(in *pb.DocumentCreateReq) (*pb.DocumentCreateResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DocumentCreateResp{}, nil
}
