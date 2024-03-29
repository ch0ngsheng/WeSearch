package mq

import (
	"io"

	"github.com/Shopify/sarama"
	"github.com/zeromicro/go-zero/core/logx"
)

type Producer interface {
	Send(topic string, data []byte) error
	io.Closer
}

type kafkaProducer struct {
	client sarama.SyncProducer
}

func (p kafkaProducer) Close() error {
	return p.client.Close()
}

func (p kafkaProducer) Send(topic string, data []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}
	partition, offset, err := p.client.SendMessage(msg)
	logx.Infof("kafka send message, partition: %d, offset: %d", partition, offset)
	return err
}
