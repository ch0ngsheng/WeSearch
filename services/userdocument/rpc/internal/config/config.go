package config

import (
	"github.com/zeromicro/go-zero/zrpc"

	"chongsheng.art/wesearch/services/userdocument/message"
)

type Config struct {
	zrpc.RpcServerConf

	Kafka message.KafkaConfig

	Mysql struct {
		DataSource string
	}

	Version string
}
