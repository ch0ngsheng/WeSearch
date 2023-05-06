package logic

import (
	"context"
	"log"

	"github.com/zeromicro/go-zero/core/logx"

	"chongsheng.art/wesearch/internal/esutil"
	"chongsheng.art/wesearch/services/retrieve/rpc/internal/svc"
	"chongsheng.art/wesearch/services/retrieve/rpc/pb"
)

const (
	searchKey = "content"
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
	res, err := l.svcCtx.ESClient.Search(l.ctx, &esutil.SearchDocReq{
		DocIDList: in.DocIDs,
		Key:       searchKey,
		Keywords:  in.KeyWords,
	})
	if err != nil {
		log.Printf("search failed, err: %v\n", err)
		return nil, err
	}

	resp := &pb.SearchResp{
		UID: in.UID,
	}
	for _, item := range res.DocIDList {
		resp.List = append(resp.List, &pb.SearchItem{
			DocID: item.DocID,
			Score: item.Score,
		})
	}
	log.Printf("search for user %s, keyword %s\n", in.UID, in.KeyWords)
	return resp, nil
}
