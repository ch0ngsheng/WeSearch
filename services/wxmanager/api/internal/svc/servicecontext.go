package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"chongsheng.art/wesearch/services/userdocument/rpc/userdocument"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/config"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/wemsg"
	"chongsheng.art/wesearch/services/wxmanager/wx"
)

type ServiceContext struct {
	Config config.Config

	WeChatAccount *wx.OfficialAccount
	WxMsgHandler  wemsg.Handler

	UserDocRpc userdocument.UserDocument
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := &ServiceContext{
		Config:        c,
		WeChatAccount: wx.NewOfficialAccount(wx.InitWechat(), buildWxConfig(c)),
		UserDocRpc:    userdocument.NewUserDocument(zrpc.MustNewClient(c.UserDocRpcConf)),
	}
	ctx.WxMsgHandler = wemsg.NewWxMsgHandler(c.WeSearch, ctx)

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
