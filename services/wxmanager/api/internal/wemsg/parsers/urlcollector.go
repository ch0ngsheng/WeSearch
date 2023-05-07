package parsers

import (
	"context"
	"log"

	"github.com/silenceper/wechat/v2/officialaccount/message"

	"chongsheng.art/wesearch/services/userdocument/rpc/pb"
)

// NewUrlCollectorParser 处理prefix开头的消息
func NewUrlCollectorParser(prefix string) Parser {
	return &urlCollector{prefix: prefix}
}

type urlCollector struct {
	prefix string
}

func (u urlCollector) Parse(obj *HandlerObj, msg *message.MixMessage) (string, error) {
	req := &pb.DocumentCollectReq{
		WxUID: msg.OpenID,
		URL:   msg.Content,
	}
	_, err := obj.UserDocRpc.CreateDoc(context.Background(), req)
	if err != nil {
		log.Printf("create doc, %v\n", err)
		return "", err
	}

	return "ok", nil
}

func (u urlCollector) Prefix() string {
	return u.prefix
}
