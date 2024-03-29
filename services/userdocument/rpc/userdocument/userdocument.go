// Code generated by goctl. DO NOT EDIT.
// Source: userdoc.proto

package userdocument

import (
	"context"

	"chongsheng.art/wesearch/services/userdocument/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DocItem             = pb.DocItem
	DocumentCollectReq  = pb.DocumentCollectReq
	DocumentCollectResp = pb.DocumentCollectResp
	DocumentSearchReq   = pb.DocumentSearchReq
	DocumentSearchResp  = pb.DocumentSearchResp
	ErrorResp           = pb.ErrorResp
	FindDocInfo         = pb.FindDocInfo
	FindUserDocReq      = pb.FindUserDocReq
	FindUserDocResp     = pb.FindUserDocResp
	VersionReq          = pb.VersionReq
	VersionResp         = pb.VersionResp

	UserDocument interface {
		Version(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error)
		CreateDoc(ctx context.Context, in *DocumentCollectReq, opts ...grpc.CallOption) (*DocumentCollectResp, error)
		SearchDoc(ctx context.Context, in *DocumentSearchReq, opts ...grpc.CallOption) (*DocumentSearchResp, error)
		FindUserDoc(ctx context.Context, in *FindUserDocReq, opts ...grpc.CallOption) (*FindUserDocResp, error)
	}

	defaultUserDocument struct {
		cli zrpc.Client
	}
)

func NewUserDocument(cli zrpc.Client) UserDocument {
	return &defaultUserDocument{
		cli: cli,
	}
}

func (m *defaultUserDocument) Version(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error) {
	client := pb.NewUserDocumentClient(m.cli.Conn())
	return client.Version(ctx, in, opts...)
}

func (m *defaultUserDocument) CreateDoc(ctx context.Context, in *DocumentCollectReq, opts ...grpc.CallOption) (*DocumentCollectResp, error) {
	client := pb.NewUserDocumentClient(m.cli.Conn())
	return client.CreateDoc(ctx, in, opts...)
}

func (m *defaultUserDocument) SearchDoc(ctx context.Context, in *DocumentSearchReq, opts ...grpc.CallOption) (*DocumentSearchResp, error) {
	client := pb.NewUserDocumentClient(m.cli.Conn())
	return client.SearchDoc(ctx, in, opts...)
}

func (m *defaultUserDocument) FindUserDoc(ctx context.Context, in *FindUserDocReq, opts ...grpc.CallOption) (*FindUserDocResp, error) {
	client := pb.NewUserDocumentClient(m.cli.Conn())
	return client.FindUserDoc(ctx, in, opts...)
}
