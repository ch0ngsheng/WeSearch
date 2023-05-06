package config

import (
	"github.com/zeromicro/go-zero/zrpc"

	"chongsheng.art/wesearch/internal/message"
)

type Config struct {
	zrpc.RpcServerConf

	Kafka message.KafkaConfig

	Mysql struct {
		DataSource string
	}

	RetrieveRpcConf zrpc.RpcClientConf
}
