package config

import (
	"github.com/zeromicro/go-zero/zrpc"

	"chongsheng.art/wesearch/internal/esutil"
	"chongsheng.art/wesearch/internal/message"
)

type Config struct {
	zrpc.RpcServerConf
	Version string

	Kafka message.KafkaConfig

	Elasticsearch esutil.Config
}
