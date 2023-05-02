package logic

import (
	"context"

	"chongsheng.art/wesearch/services/userdocument/rpc/internal/svc"
	"chongsheng.art/wesearch/services/userdocument/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserDocLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserDocLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserDocLogic {
	return &FindUserDocLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserDocLogic) FindUserDoc(in *pb.FindUserDocReq) (*pb.FindUserDocResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FindUserDocResp{}, nil
}
