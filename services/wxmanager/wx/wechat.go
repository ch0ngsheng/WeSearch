package wx

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/zeromicro/go-zero/core/logx"
)

type Config struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
}

type OfficialAccount struct {
	Wechat          *wechat.Wechat
	OfficialAccount *officialaccount.OfficialAccount
}

func InitWechat() *wechat.Wechat {
	wc := wechat.NewWechat()
	c := cache.NewMemory()
	wc.SetCache(c)
	return wc
}

func NewOfficialAccount(wc *wechat.Wechat, cfg *Config) *OfficialAccount {
	offCfg := &config.Config{
		AppID:          cfg.AppID,
		AppSecret:      cfg.AppSecret,
		Token:          cfg.Token,
		EncodingAESKey: cfg.EncodingAESKey,
	}
	logx.Infof("new wechat client, appID: %s", offCfg.AppID)
	officialAccount := wc.GetOfficialAccount(offCfg)
	return &OfficialAccount{
		Wechat:          wc,
		OfficialAccount: officialAccount,
	}
}
