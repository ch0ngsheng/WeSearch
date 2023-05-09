package svc

import (
	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/zrpc"

	"chongsheng.art/wesearch/internal/common/interceptor"
	"chongsheng.art/wesearch/internal/common/xerror"
	"chongsheng.art/wesearch/services/userdocument/rpc/pb"
	"chongsheng.art/wesearch/services/userdocument/rpc/userdocument"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/config"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/wemsg"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/wemsg/parsers"
	"chongsheng.art/wesearch/services/wxmanager/wx"
)

type ServiceContext struct {
	Config config.Config

	WeChatAccount *wx.OfficialAccount
	WxMsgHandler  wemsg.Handler

	UserDocRpc userdocument.UserDocument
}

func NewServiceContext(c config.Config) *ServiceContext {
	userDocRpc := userdocument.NewUserDocument(
		zrpc.MustNewClient(
			c.UserDocRpcConf,
			zrpc.WithUnaryClientInterceptor(
				interceptor.NewClientErrInterceptor(func(message proto.Message) xerror.SearchErr {
					respErr := message.(*pb.ErrorResp)
					return xerror.NewSearchErr(respErr.GetErrCode(), respErr.GetErrMsg(), respErr.GetDetail())
				})),
		),
	)
	ctx := &ServiceContext{
		Config:        c,
		WeChatAccount: wx.NewOfficialAccount(wx.InitWechat(), buildWxConfig(c)),
		UserDocRpc:    userDocRpc,
		WxMsgHandler:  wemsg.NewWxMsgHandler(c.WeSearch, &parsers.HandlerObj{UserDocRpc: userDocRpc}),
	}

	return ctx
}

func buildWxConfig(c config.Config) *wx.Config {
	return &wx.Config{
		AppID:          c.WeChat.AppID,
		AppSecret:      c.WeChat.AppSecret,
		Token:          c.WeChat.Token,
		EncodingAESKey: c.WeChat.EncodingAESKey,
	}
}
