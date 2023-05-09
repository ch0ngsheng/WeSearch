package logic

import (
	"context"
	"sort"

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
		return nil, err
	}

	// 按score倒序
	sort.Sort(res)

	resp := &pb.SearchResp{
		UID: in.UID,
	}
	for _, item := range res.DocIDList {
		resp.List = append(resp.List, &pb.SearchItem{
			DocID: item.DocID,
			Score: item.Score,
		})
	}
	logx.Infof("search for user %s, keyword %s, docs %s, resp: %v", in.UID, in.KeyWords, in.DocIDs, resp)
	return resp, nil
}
