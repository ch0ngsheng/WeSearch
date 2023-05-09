package mqmsg

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"chongsheng.art/wesearch/internal/message"
	"chongsheng.art/wesearch/internal/mq"
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
		return errors.Wrap(err, "msg unmarshal")
	}
	logx.Infof("receive from kafka: docID %d, title %s, description %s.", msg.DocID, msg.Title, msg.Description)

	err := d.svcCtx.DocModel.Trans(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		doc, err := d.svcCtx.DocModel.FindOne(ctx, session, msg.DocID)
		if err == sqlx.ErrNotFound {
			return errors.New(fmt.Sprintf("doc id %d not found, title is %s", msg.DocID, msg.Title))
		}
		if err != nil {
			return errors.Wrap(err, "find one doc.")
		}

		// mark Update默认使用参数doc对象的所有字段值，更新数据库对应记录的所有字段
		// 要么先全查出来，要么写一个UpdateXXX方法
		doc.Title = sql.NullString{String: msg.Title, Valid: true}
		doc.Description = sql.NullString{String: msg.Description, Valid: true}

		if _, err := d.svcCtx.DocModel.Update(ctx, session, doc); err != nil {
			return errors.Wrap(err, "update doc.")
		}
		return nil
	})
	if err != nil {
		logx.Errorf("update doc transaction, err: %+v", err)
		return errors.Wrap(err, "update doc transaction.")
	}

	return nil
}

func MustStartConsumer(cfg message.KafkaConfig, fn func(msg []byte) error) {
	consumer := mq.MustNewMqConsumer(cfg.Brokers)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logx.Errorf("kafka consumer panic, %+v", err)
			}
		}()
		consumer.Consume(cfg.ConsumerGroup, cfg.TopicParseDoc, fn)
	}()
}
