package main

import (
	"flag"
	"fmt"

	"chongsheng.art/wesearch/services/wxmanager/api/internal/config"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/handler"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/manager.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
