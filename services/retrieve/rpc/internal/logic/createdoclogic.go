package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"chongsheng.art/wesearch/internal/esutil"
	"chongsheng.art/wesearch/services/retrieve/rpc/internal/svc"
	"chongsheng.art/wesearch/services/retrieve/rpc/pb"
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
	err := l.svcCtx.ESClient.CreateDoc(l.ctx, &esutil.CreateDocReq{
		ID:      in.DocID,
		Content: in.DocURL,
	})
	if err != nil {
		logx.Errorf("create doc, %+v", err)
		return nil, err
	}

	logx.Infof("doc %s %s created.", in.DocID, in.DocURL)
	return &pb.DocumentCreateResp{DocID: in.DocID}, nil
}
