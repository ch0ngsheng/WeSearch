// Code generated by goctl. DO NOT EDIT.
// Source: service.proto

package server

import (
	"context"

	"chongsheng.art/wesearch/services/userdocument/rpc/internal/logic"
	"chongsheng.art/wesearch/services/userdocument/rpc/internal/svc"
	"chongsheng.art/wesearch/services/userdocument/rpc/pb"
)

type UserDocumentServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserDocumentServer
}

func NewUserDocumentServer(svcCtx *svc.ServiceContext) *UserDocumentServer {
	return &UserDocumentServer{
		svcCtx: svcCtx,
	}
}

func (s *UserDocumentServer) Version(ctx context.Context, in *pb.VersionReq) (*pb.VersionResp, error) {
	l := logic.NewVersionLogic(ctx, s.svcCtx)
	return l.Version(in)
}

func (s *UserDocumentServer) CreateDoc(ctx context.Context, in *pb.DocumentCollectReq) (*pb.DocumentCollectResp, error) {
	l := logic.NewCreateDocLogic(ctx, s.svcCtx)
	return l.CreateDoc(in)
}

func (s *UserDocumentServer) SearchDoc(ctx context.Context, in *pb.DocumentSearchReq) (*pb.DocumentSearchResp, error) {
	l := logic.NewSearchDocLogic(ctx, s.svcCtx)
	return l.SearchDoc(in)
}

func (s *UserDocumentServer) FindUserDoc(ctx context.Context, in *pb.FindUserDocReq) (*pb.FindUserDocResp, error) {
	l := logic.NewFindUserDocLogic(ctx, s.svcCtx)
	return l.FindUserDoc(in)
}
