package svc

import (
	"chongsheng.art/wesearch/services/wxmanager/api/internal/config"
	"chongsheng.art/wesearch/services/wxmanager/wx"
)

type ServiceContext struct {
	Config config.Config

	WeChatAccount *wx.OfficialAccount
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		WeChatAccount: wx.NewOfficialAccount(wx.InitWechat(), buildWxConfig(c)),
	}
}

func buildWxConfig(c config.Config) *wx.Config {
	return &wx.Config{
		AppID:          c.WeChat.AppID,
		AppSecret:      c.WeChat.AppSecret,
		Token:          c.WeChat.Token,
		EncodingAESKey: c.WeChat.EncodingAESKey,
	}
}
