package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"chongsheng.art/wesearch/internal/message"
	"chongsheng.art/wesearch/services/userdocument/rpc/internal/config"
	"chongsheng.art/wesearch/services/userdocument/rpc/internal/mqmsg"
	"chongsheng.art/wesearch/services/userdocument/rpc/internal/server"
	"chongsheng.art/wesearch/services/userdocument/rpc/internal/svc"
	"chongsheng.art/wesearch/services/userdocument/rpc/pb"
)

var configFile = flag.String("f", "etc/userdoc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserDocumentServer(grpcServer, server.NewUserDocumentServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 消息处理
	message.MustNewWorker(c.Kafka, mqmsg.NewConsumer(ctx).ReadDocAnalysisMessage)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
