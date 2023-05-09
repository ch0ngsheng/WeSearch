package svc

import (
	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"

	"chongsheng.art/wesearch/internal/common/interceptor"
	"chongsheng.art/wesearch/internal/common/xerror"
	"chongsheng.art/wesearch/internal/mq"
	"chongsheng.art/wesearch/services/retrieve/rpc/pb"
	"chongsheng.art/wesearch/services/retrieve/rpc/retrieve"
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
		UserModel:    model.NewUsersModel(conn),
		DocModel:     model.NewDocumentsModel(conn),
		RetrieveRpc: retrieve.NewRetrieve(
			zrpc.MustNewClient(
				c.RetrieveRpcConf,
				zrpc.WithUnaryClientInterceptor(interceptor.NewClientErrInterceptor(func(message proto.Message) xerror.SearchErr {
					errResp := message.(*pb.ErrorResp)
					return xerror.NewSearchErr(errResp.GetErrCode(), errResp.GetErrMsg(), errResp.GetDetail())
				})),
			),
		),
	}
}
