package parsers

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

// NewUrlQueryParser 处理prefix开头的消息
func NewUrlQueryParser(prefix string) Parser {
	return &urlQuery{prefix: prefix}
}

type urlQuery struct {
	prefix string
}

func (u urlQuery) Parse(obj *HandlerObj, msg *message.MixMessage) (string, error) {
	return "todo", nil
}

func (u urlQuery) Prefix() string {
	return u.prefix
}
