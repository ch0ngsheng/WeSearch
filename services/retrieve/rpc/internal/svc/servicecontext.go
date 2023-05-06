package svc

import (
	"chongsheng.art/wesearch/internal/esutil"
	"chongsheng.art/wesearch/internal/mq"
	"chongsheng.art/wesearch/services/retrieve/rpc/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	Producer mq.Producer
	Consumer mq.Consumer
	ESClient esutil.Retriever
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Producer: mq.MustNewMqProducer(c.Kafka.Brokers),
		Consumer: mq.MustNewMqConsumer(c.Kafka.Brokers),
		ESClient: esutil.MustNewRetriever(c.Elasticsearch.Addresses, c.Elasticsearch.APIKey, c.Elasticsearch.Index),
	}
}
