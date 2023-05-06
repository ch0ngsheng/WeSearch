package mqmsg

import (
	"chongsheng.art/wesearch/internal/htmlutil"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"chongsheng.art/wesearch/internal/esutil"
	"chongsheng.art/wesearch/internal/message"
	"chongsheng.art/wesearch/internal/mq"
	"chongsheng.art/wesearch/services/retrieve/rpc/internal/svc"
)

type Consumer interface {
	ReadCreateDocMessage(data []byte) error
}

func NewConsumer(svcCtx *svc.ServiceContext) Consumer {
	return docCreator{svcCtx: svcCtx}
}

type docCreator struct {
	svcCtx *svc.ServiceContext
}

// ReadCreateDocMessage 读取userdocument服务发来的消息
func (d docCreator) ReadCreateDocMessage(data []byte) error {
	msg := &message.DocCollection{}
	if err := json.Unmarshal(data, msg); err != nil {
		log.Printf("failed to unmarshal kafka msg %s, err: %v", data, err)
		return err
	}

	article, err := htmlutil.Parse(msg.URL)
	if err != nil {
		return err
	}

	req := &esutil.CreateDocReq{
		ID:      fmt.Sprintf("%d", msg.DocID),
		Content: article.Content,
	}

	// todo 两个操作的原子性保证？
	// ES
	if err = d.svcCtx.ESClient.CreateDoc(context.Background(), req); err != nil {
		return err
	}

	// Kafka
	docMsg, err := message.BuildDocMsg(message.DocAnalysis{
		DocID:       msg.DocID,
		Title:       article.Title,
		Description: article.Desc,
	})
	if err != nil {
		log.Printf("failed to marshal doc analysis msg, err: %v\n", err)
		return err
	}
	return d.svcCtx.Producer.Send(d.svcCtx.Config.Kafka.TopicParseDoc, docMsg)
}

func MustStartConsumer(cfg message.KafkaConfig, fn func(msg []byte) error) {
	consumer := mq.MustNewMqConsumer(cfg.Brokers)

	go func() {
		consumer.Consume(cfg.ConsumerGroup, cfg.TopicCreateDoc, fn)
	}()
}
