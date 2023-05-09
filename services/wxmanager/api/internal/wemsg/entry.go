package wemsg

import (
	"strings"

	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/zeromicro/go-zero/core/logx"

	"chongsheng.art/wesearch/services/wxmanager/api/internal/config"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/wemsg/parsers"
)

type Handler interface {
	Do(msg *message.MixMessage) *message.Reply
}

// NewWxMsgHandler 构造消息处理器
func NewWxMsgHandler(cfg config.WeSearch, obj *parsers.HandlerObj) Handler {
	return &handler{
		config:     cfg,
		obj:        obj,
		parsers:    buildMatchTable(cfg),
		matchOrder: buildMatchOrder(cfg),
	}
}

type handler struct {
	config     config.WeSearch
	obj        *parsers.HandlerObj
	parsers    map[string]parsers.Parser
	matchOrder []string
}

func (h handler) Do(msg *message.MixMessage) *message.Reply {
	logx.Infof("recv WX message: %s", msg.Content)

	var key = ""
	for _, pp := range h.matchOrder {
		if strings.HasPrefix(msg.Content, pp) {
			key = pp
		}
	}

	resp, err := h.parsers[key].Parse(h.obj, msg)
	if err != nil {
		logx.Errorf("Handle WX message, %+v", err)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("system error")}
	}

	logx.Infof("response WX message.")
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText(resp)}
}

func buildMatchOrder(cfg config.WeSearch) []string {
	return []string{cfg.KeyPrefix.UrlCollector, cfg.KeyPrefix.UrlQuery}
}

func buildMatchTable(cfg config.WeSearch) map[string]parsers.Parser {
	matchTable := make(map[string]parsers.Parser)
	matchTable[cfg.KeyPrefix.UrlCollector] = parsers.NewUrlCollectorParser(cfg.KeyPrefix.UrlCollector)
	matchTable[cfg.KeyPrefix.UrlQuery] = parsers.NewUrlQueryParser(cfg.KeyPrefix.UrlQuery)
	matchTable[""] = parsers.NewDefaultParser()

	return matchTable
}
