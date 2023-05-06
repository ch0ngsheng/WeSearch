package logic

import (
	"chongsheng.art/wesearch/services/userdocument/rpc/internal/svc"
	"chongsheng.art/wesearch/services/userdocument/rpc/pb"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"

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
	resp := &pb.FindUserDocResp{
		WxUID: in.WxUID,
	}

	err := l.svcCtx.UserModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		user, err := l.svcCtx.UserModel.FindOneByOpenid(l.ctx, session, in.WxUID)
		if err == sqlx.ErrNotFound {
			return nil
		}
		if err != nil {
			return err
		}

		docs, err := l.svcCtx.DocModel.FindByUID(l.ctx, session, user.Id)
		if err != nil {
			return err
		}

		resp.List = make([]*pb.FindDocInfo, 0, len(docs))
		for _, doc := range docs {
			resp.List = append(resp.List, &pb.FindDocInfo{
				URL:         doc.Url,
				Title:       doc.Title.String,
				CollectTime: timestamppb.New(doc.CreatedAt),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}
