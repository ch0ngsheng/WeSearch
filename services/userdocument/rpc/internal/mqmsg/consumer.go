package mqmsg

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"chongsheng.art/wesearch/internal/message"
	"chongsheng.art/wesearch/internal/mq"
	"chongsheng.art/wesearch/services/userdocument/model"
	"chongsheng.art/wesearch/services/userdocument/rpc/internal/svc"
)

type Consumer interface {
	ReadDocAnalysisMessage(data []byte) error
}

type docAnalysis struct {
	svcCtx *svc.ServiceContext
}

func NewConsumerObj(svcCtx *svc.ServiceContext) Consumer {
	return docAnalysis{svcCtx: svcCtx}
}

func (d docAnalysis) ReadDocAnalysisMessage(data []byte) error {
	msg := &message.DocAnalysis{}
	if err := json.Unmarshal(data, msg); err != nil {
		log.Printf("msg unmarshal %v", err)
		return err
	}
	log.Printf("rec: docID %d, title %s, description %s", msg.DocID, msg.Title, msg.Description)

	doc := &model.Documents{
		Id:          msg.DocID,
		Title:       sql.NullString{String: msg.Title, Valid: true},
		Description: sql.NullString{String: msg.Description, Valid: true},
	}
	err := d.svcCtx.DocModel.Trans(context.Background(), func(context context.Context, session sqlx.Session) error {
		if _, err := d.svcCtx.DocModel.Update(context, session, doc); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("failed to update doc info, err: %v\n", err)
	}

	return nil
}

func MustStartConsumer(cfg message.KafkaConfig, fn func(msg []byte) error) {
	consumer := mq.MustNewMqConsumer(cfg.Brokers)

	go func() {
		consumer.Consume(cfg.ConsumerGroup, cfg.TopicParseDoc, fn)
	}()
}
