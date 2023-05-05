package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"chongsheng.art/wesearch/internal/mq"
	"chongsheng.art/wesearch/services/userdocument/model"
	"chongsheng.art/wesearch/services/userdocument/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Producer mq.Producer
	Consumer mq.Consumer

	UserDocModel model.UserDocsModel
	UserModel    model.UsersModel
	DocModel     model.DocumentsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:       c,
		Producer:     mq.MustNewMqProducer(c.Kafka.Brokers),
		Consumer:     mq.MustNewMqConsumer(c.Kafka.Brokers),
		UserDocModel: model.NewUserDocsModel(conn),
	}
}
