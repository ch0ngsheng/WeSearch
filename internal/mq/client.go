package mq

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func MustNewMqProducer(addr []string) Producer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // waits for all in-sync replicas to commit before responding
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true // 同步生产者必须同时开启 Return.Successes

	client, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		log.Fatalf("fail to start producer, err: %v\n", err)
	}

	return &kafkaProducer{client: client}
}

func MustNewMqConsumer(addr []string) Consumer {
	config := sarama.NewConfig()
	// 开启自动提交 offset，samara 库会定时把最新的 offset 信息提交给 kafka
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	consumer, err := sarama.NewClient(addr, config)
	if err != nil {
		log.Fatalf("fail to start consumer, err:%v\n", err)
	}

	return &kafkaConsumer{client: consumer}
}
