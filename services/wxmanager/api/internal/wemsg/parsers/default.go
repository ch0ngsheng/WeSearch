package parsers

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

type defaultParser struct {
}

func (d defaultParser) Parse(obj *HandlerObj, msg *message.MixMessage) (string, error) {
	return msg.Content, nil
}

func (d defaultParser) Prefix() string {
	return ""
}

func NewDefaultParser() Parser {
	return &defaultParser{}
}
