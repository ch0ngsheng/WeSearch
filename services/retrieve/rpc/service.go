package main

import (
	"flag"
	"fmt"

	"chongsheng.art/wesearch/services/retrieve/rpc/internal/config"
	"chongsheng.art/wesearch/services/retrieve/rpc/internal/server"
	"chongsheng.art/wesearch/services/retrieve/rpc/internal/svc"
	"chongsheng.art/wesearch/services/retrieve/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
