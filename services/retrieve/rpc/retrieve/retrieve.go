// Code generated by goctl. DO NOT EDIT.
// Source: service.proto

package retrieve

import (
	"context"

	"chongsheng.art/wesearch/services/retrieve/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DocumentCreateReq  = pb.DocumentCreateReq
	DocumentCreateResp = pb.DocumentCreateResp
	ErrorResp          = pb.ErrorResp
	SearchItem         = pb.SearchItem
	SearchReq          = pb.SearchReq
	SearchResp         = pb.SearchResp
	VersionReq         = pb.VersionReq
	VersionResp        = pb.VersionResp

	Retrieve interface {
		Version(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error)
		CreateDoc(ctx context.Context, in *DocumentCreateReq, opts ...grpc.CallOption) (*DocumentCreateResp, error)
		Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchResp, error)
	}

	defaultRetrieve struct {
		cli zrpc.Client
	}
)

func NewRetrieve(cli zrpc.Client) Retrieve {
	return &defaultRetrieve{
		cli: cli,
	}
}

func (m *defaultRetrieve) Version(ctx context.Context, in *VersionReq, opts ...grpc.CallOption) (*VersionResp, error) {
	client := pb.NewRetrieveClient(m.cli.Conn())
	return client.Version(ctx, in, opts...)
}

func (m *defaultRetrieve) CreateDoc(ctx context.Context, in *DocumentCreateReq, opts ...grpc.CallOption) (*DocumentCreateResp, error) {
	client := pb.NewRetrieveClient(m.cli.Conn())
	return client.CreateDoc(ctx, in, opts...)
}

func (m *defaultRetrieve) Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchResp, error) {
	client := pb.NewRetrieveClient(m.cli.Conn())
	return client.Search(ctx, in, opts...)
}
