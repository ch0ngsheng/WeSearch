package parsers

import (
	"context"
	"log"

	"github.com/silenceper/wechat/v2/officialaccount/message"

	"chongsheng.art/wesearch/services/userdocument/rpc/pb"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/svc"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/wemsg"
)

// NewUrlCollectorParser 处理prefix开头的消息
func NewUrlCollectorParser(prefix string) wemsg.Parser {
	return &urlCollector{prefix: prefix}
}

type urlCollector struct {
	prefix string
}

func (u urlCollector) Parse(svcCtx *svc.ServiceContext, msg *message.MixMessage) (string, error) {
	req := &pb.DocumentCollectReq{
		WxUID: msg.OpenID,
		URL:   msg.Content,
	}
	_, err := svcCtx.UserDocRpc.CreateDoc(context.Background(), req)
	if err != nil {
		log.Printf("create doc, %v\n", err)
		return "", err
	}

	return "ok", nil
}

func (u urlCollector) Prefix() string {
	return u.prefix
}
