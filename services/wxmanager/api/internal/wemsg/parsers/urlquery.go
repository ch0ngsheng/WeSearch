package parsers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/zeromicro/go-zero/core/logx"

	"chongsheng.art/wesearch/services/userdocument/rpc/userdocument"
)

// NewUrlQueryParser 处理prefix开头的消息
func NewUrlQueryParser(prefix string) Parser {
	return &urlQuery{prefix: prefix}
}

type urlQuery struct {
	prefix string
}

func (u urlQuery) Parse(obj *HandlerObj, msg *message.MixMessage) (string, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	resp, err := obj.UserDocRpc.SearchDoc(ctx, &userdocument.DocumentSearchReq{
		WxUID:    string(msg.FromUserName),
		Keywords: u.getSearchKeywords(msg.Content),
	})
	if err != nil {
		logx.Errorf("search error for %s, url: %s, err: %v\n", msg.OpenID, msg.Content, err)
		return "", errors.Wrap(err, "rpc search doc")
	}
	return buildRespContent(resp), nil
}

func (u urlQuery) Prefix() string {
	return u.prefix
}

func buildRespContent(entry *userdocument.DocumentSearchResp) string {
	if len(entry.List) == 0 {
		return "not found"
	}

	item := `title: %s
url: %s

`

	sb := strings.Builder{}
	for _, e := range entry.List {
		sb.WriteString(fmt.Sprintf(item, e.GetTitle(), e.GetURL()))
	}

	return sb.String()
}

func (u urlQuery) getSearchKeywords(str string) []string {
	str = strings.TrimLeft(str, u.Prefix())
	str = strings.TrimSpace(str)
	return strings.Split(str, " ")
}
