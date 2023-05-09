package mqmsg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"chongsheng.art/wesearch/internal/esutil"
	"chongsheng.art/wesearch/internal/htmlutil"
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
	// mark 将依赖的对象作为属性，在New方法中引入，这样就不需要加到接口参数中了。
	svcCtx *svc.ServiceContext
}

// ReadCreateDocMessage 读取userdocument服务发来的消息
func (d docCreator) ReadCreateDocMessage(data []byte) error {
	msg := &message.DocCollection{}
	if err := json.Unmarshal(data, msg); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to unmarshal kafka msg %s, err: %v", data, err))
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
		return errors.Wrap(err, fmt.Sprintf("failed to marshal doc analysis msg for kafka, err: %v\n", err))
	}
	return d.svcCtx.Producer.Send(d.svcCtx.Config.Kafka.TopicParseDoc, docMsg)
}

func MustStartConsumer(cfg message.KafkaConfig, fn func(msg []byte) error) {
	consumer := mq.MustNewMqConsumer(cfg.Brokers)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logx.Errorf("kafka consumer panic, %+v", err)
			}
		}()
		consumer.Consume(cfg.ConsumerGroup, cfg.TopicCreateDoc, fn)
	}()
}
