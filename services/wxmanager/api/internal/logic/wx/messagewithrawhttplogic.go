package wx

import (
	"context"
	"log"
	"net/http"

	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/zeromicro/go-zero/core/logx"

	"chongsheng.art/wesearch/services/wxmanager/api/internal/svc"
)

type MessageWithRawHTTPLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageWithRawHTTPLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageWithRawHTTPLogic {
	return &MessageWithRawHTTPLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageWithRawHTTPLogic) MessageWithRawHTTP(w http.ResponseWriter, r *http.Request) {
	// generated by custom templates
	log.Println("incoming..")
	server := l.svcCtx.WeChatAccount.OfficialAccount.GetServer(r, w)
	server.SkipValidate(true)
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		//回复消息：演示回复用户发送的消息
		log.Printf("recv: %s\n", msg.Content)
		text := message.NewText(msg.Content)
		log.Println(text)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Serve Error, err=%+v", err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Send Error, err=%+v", err)
		return
	}
}
