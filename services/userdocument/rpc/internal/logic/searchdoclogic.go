package logic

import (
	"chongsheng.art/wesearch/services/retrieve/rpc/retrieve"
	"chongsheng.art/wesearch/services/userdocument/rpc/internal/svc"
	"chongsheng.art/wesearch/services/userdocument/rpc/pb"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"

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
	resp := &pb.DocumentSearchResp{WxUID: in.WxUID}

	err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user, err := l.svcCtx.UserModel.FindOneByOpenid(ctx, session, in.WxUID)
		if err == sqlx.ErrNotFound {
			return nil
		}
		if err != nil {
			return err
		}

		myDocs, err := l.svcCtx.DocModel.FindByUID(ctx, session, user.Id)
		if err != nil {
			return err
		}
		if len(myDocs) == 0 {
			return nil
		}

		// 获取用户的所有文档ID
		var docIDArr = make([]string, 0, len(myDocs))
		for _, doc := range myDocs {
			docIDArr = append(docIDArr, fmt.Sprintf("%d", doc.Id))
		}

		// 调用retrieve服务
		req := &retrieve.SearchReq{
			UID:      "",
			KeyWords: in.GetKeywords(),
			DocIDs:   docIDArr,
		}
		search, err := l.svcCtx.RetrieveRpc.Search(ctx, req)
		if err != nil {
			return err
		}
		if len(search.List) == 0 {
			return nil
		}

		// 解析结果
		var matchDocs = make([]*pb.DocItem, 0, len(search.List))
		for _, item := range search.List {
			matchDocID, _ := strconv.ParseInt(item.DocID, 10, 64)
			docInfo, err := l.svcCtx.DocModel.FindOneByUIDAndDocID(ctx, session, user.Id, matchDocID)
			if err != nil {
				log.Printf("using doc id %s to find doc, err: %v", matchDocID, err)
				return err
			}
			matchDocs = append(matchDocs, &pb.DocItem{
				URL:         docInfo.Url,
				Title:       docInfo.Title.String,
				Score:       item.Score,
				CollectTime: timestamppb.New(docInfo.CreatedAt),
			})
		}

		resp.List = matchDocs
		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
