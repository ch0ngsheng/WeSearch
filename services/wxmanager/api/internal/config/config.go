package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	WeChat struct {
		AppID          string
		AppSecret      string
		Token          string
		EncodingAESKey string
	}
}
