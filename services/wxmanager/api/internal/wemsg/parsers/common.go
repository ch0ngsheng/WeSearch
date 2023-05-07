package parsers

import (
	"chongsheng.art/wesearch/services/userdocument/rpc/userdocument"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

type Parser interface {
	Prefix() string
	Parse(obj *HandlerObj, msg *message.MixMessage) (string, error)
}

type HandlerObj struct {
	UserDocRpc userdocument.UserDocument
}
