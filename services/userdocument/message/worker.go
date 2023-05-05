package message

import (
	"encoding/json"
	"log"

	"chongsheng.art/wesearch/internal/mq"
)

type KafkaConfig struct {
	Brokers       []string
	Topic         string
	ConsumerGroup string
}

func MustNewWorker(cfg KafkaConfig, fn func(msg []byte) error) {
	consumer := mq.MustNewMqConsumer(cfg.Brokers)

	go func() {
		consumer.Consume(cfg.ConsumerGroup, cfg.Topic, fn)
	}()
}

func ReadMessage(data []byte) error {
	msg := &Body{}
	if err := json.Unmarshal(data, msg); err != nil {
		log.Printf("msg unmarshal %v", err)
		return err
	}

	log.Printf("rec: docID %d, content %s", msg.DocID, msg.Content)
	return nil
}
