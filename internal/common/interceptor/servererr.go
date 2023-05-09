package interceptor

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"chongsheng.art/wesearch/internal/common/xerror"
)

type BuildStatusDetailFunc func(xerror.SearchErr) proto.Message

func NewServerErrInterceptor(fn BuildStatusDetailFunc) grpc.UnaryServerInterceptor {
	i := &severErrInterceptor{fn: fn}
	return i.do
}

type severErrInterceptor struct {
	fn BuildStatusDetailFunc
}

func (sei severErrInterceptor) do(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err == nil {
		return resp, nil
	}

	realErr := errors.Cause(err)

	switch realErr.(type) {
	case xerror.SearchErr:
		searchErr := realErr.(xerror.SearchErr)
		s, _ := status.New(codes.Code(searchErr.GetCode()), searchErr.GetMsg()).WithDetails(sei.fn(searchErr))
		return resp, s.Err()
	default:
		// internal error
		s := status.New(codes.Code(xerror.InternalErr.GetCode()), xerror.InternalErr.GetMsg())
		logx.Info("internal server error: %+v", err)
		return resp, s.Err()
	}
}
