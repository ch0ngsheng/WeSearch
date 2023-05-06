package svc

import (
	"chongsheng.art/wesearch/services/retrieve/rpc/retrieve"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"

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

	RetrieveRpc retrieve.Retrieve
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:       c,
		Producer:     mq.MustNewMqProducer(c.Kafka.Brokers),
		Consumer:     mq.MustNewMqConsumer(c.Kafka.Brokers),
		UserDocModel: model.NewUserDocsModel(conn),
		RetrieveRpc:  retrieve.NewRetrieve(zrpc.MustNewClient(c.RetrieveRpcConf)),
	}
}
