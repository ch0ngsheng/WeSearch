package parsers

import (
	"context"
	
	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/zeromicro/go-zero/core/logx"

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
		WxUID: string(msg.FromUserName),
		URL:   msg.Content,
	}
	_, err := obj.UserDocRpc.CreateDoc(context.Background(), req)
	if err != nil {
		logx.Errorf("create doc, %v\n", err)
		return "", errors.Wrap(err, "create doc.")
	}

	return "ok", nil
}

func (u urlCollector) Prefix() string {
	return u.prefix
}
