package parsers

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"

	"chongsheng.art/wesearch/services/wxmanager/api/internal/svc"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/wemsg"
)

// NewUrlQueryParser 处理prefix开头的消息
func NewUrlQueryParser(prefix string) wemsg.Parser {
	return &urlQuery{prefix: prefix}
}

type urlQuery struct {
	prefix string
}

func (u urlQuery) Parse(ctx *svc.ServiceContext, msg *message.MixMessage) (string, error) {
	return "todo", nil
}

func (u urlQuery) Prefix() string {
	return u.prefix
}
