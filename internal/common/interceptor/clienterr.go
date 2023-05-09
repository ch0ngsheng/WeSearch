package interceptor

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"chongsheng.art/wesearch/internal/common/xerror"
)

type GetRealErrFunc func(proto.Message) xerror.SearchErr

type clientErrInterceptor struct {
	fn GetRealErrFunc
}

func NewClientErrInterceptor(fn GetRealErrFunc) grpc.UnaryClientInterceptor {
	i := &clientErrInterceptor{fn: fn}
	return i.do
}

func (cei clientErrInterceptor) do(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err == nil {
		return nil
	}

	var searchErr xerror.SearchErr

	s, ok := status.FromError(err)
	if !ok {
		panic(fmt.Sprintf("unrecognized error %+v", err))
	}

	if len(s.Details()) > 0 {
		d := s.Details()[0].(proto.Message)
		searchErr = cei.fn(d)
	} else {
		searchErr = xerror.InternalErr
	}

	return searchErr
}
