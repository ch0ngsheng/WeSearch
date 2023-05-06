package wemsg

import (
	"log"
	"strings"

	"github.com/silenceper/wechat/v2/officialaccount/message"

	"chongsheng.art/wesearch/services/wxmanager/api/internal/config"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/svc"
	"chongsheng.art/wesearch/services/wxmanager/api/internal/wemsg/parsers"
)

type Parser interface {
	Prefix() string
	Parse(ctx *svc.ServiceContext, msg *message.MixMessage) (string, error)
}

type Handler interface {
	Do(msg *message.MixMessage) *message.Reply
}

// NewWxMsgHandler 构造消息处理器
func NewWxMsgHandler(cfg config.WeSearch, ctx *svc.ServiceContext) Handler {
	return &handler{
		config:     cfg,
		svcCtx:     ctx,
		parsers:    buildMatchTable(cfg),
		matchOrder: buildMatchOrder(cfg),
	}
}

type handler struct {
	config     config.WeSearch
	svcCtx     *svc.ServiceContext
	parsers    map[string]Parser
	matchOrder []string
}

func (h handler) Do(msg *message.MixMessage) *message.Reply {
	log.Printf("recv: %s\n", msg.Content)

	var key = ""
	for _, pp := range h.matchOrder {
		if strings.HasPrefix(msg.Content, pp) {
			key = pp
		}
	}

	resp, err := h.parsers[key].Parse(h.svcCtx, msg)
	if err != nil {
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("system error")}
	}

	return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText(resp)}
}

func buildMatchOrder(cfg config.WeSearch) []string {
	return []string{cfg.KeyPrefix.UrlCollector, cfg.KeyPrefix.UrlQuery}
}

func buildMatchTable(cfg config.WeSearch) map[string]Parser {
	matchTable := make(map[string]Parser)
	matchTable[cfg.KeyPrefix.UrlCollector] = parsers.NewUrlCollectorParser(cfg.KeyPrefix.UrlCollector)
	matchTable[cfg.KeyPrefix.UrlQuery] = parsers.NewUrlQueryParser(cfg.KeyPrefix.UrlQuery)
	matchTable[""] = parsers.NewDefaultParser()

	return matchTable
}
