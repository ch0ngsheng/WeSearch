package parsers

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"

	"chongsheng.art/wesearch/services/wxmanager/api/internal/svc"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/wemsg"
)

type defaultParser struct {
}

func (d defaultParser) Parse(ctx *svc.ServiceContext, msg *message.MixMessage) (string, error) {
	return msg.Content, nil
}

func (d defaultParser) Prefix() string {
	return ""
}

func NewDefaultParser() wemsg.Parser {
	return &urlCollector{prefix: ""}
}
