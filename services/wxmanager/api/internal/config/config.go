package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	WeChat   WeChat
	WeSearch WeSearch

	UserDocRpcConf zrpc.RpcClientConf
}
