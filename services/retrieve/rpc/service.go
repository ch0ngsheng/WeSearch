package main

import (
	"flag"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"chongsheng.art/wesearch/internal/common/interceptor"
	"chongsheng.art/wesearch/internal/common/xerror"
	"chongsheng.art/wesearch/services/retrieve/rpc/internal/config"
	"chongsheng.art/wesearch/services/retrieve/rpc/internal/mqmsg"
	"chongsheng.art/wesearch/services/retrieve/rpc/internal/server"
	"chongsheng.art/wesearch/services/retrieve/rpc/internal/svc"
	"chongsheng.art/wesearch/services/retrieve/rpc/pb"
)

var configFile = flag.String("f", "etc/service.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterRetrieveServer(grpcServer, server.NewRetrieveServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 服务端拦截器
	// 错误处理
	s.AddUnaryInterceptors(interceptor.NewServerErrInterceptor(func(err xerror.SearchErr) proto.Message {
		return &pb.ErrorResp{
			ErrCode: uint32(err.GetCode()),
			ErrMsg:  err.GetMsg(),
			Detail:  err.GetExt(),
		}
	}))

	// 消息消费
	mqmsg.MustStartConsumer(c.Kafka, mqmsg.NewConsumer(ctx).ReadCreateDocMessage)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
