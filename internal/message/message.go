package message

import "encoding/json"

type Msg interface {
	Build() ([]byte, error)
}

// DocCollection 收藏文档消息
type DocCollection struct {
	DocID int64  `json:"doc_id"`
	URL   string `json:"content"`
}

// DocAnalysis 文档摘要消息
type DocAnalysis struct {
	DocID       int64  `json:"doc_id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
}

func (d DocCollection) Build() ([]byte, error) {
	return json.Marshal(d)
}

func (d DocAnalysis) Build() ([]byte, error) {
	return json.Marshal(d)
}

func BuildDocMsg(body Msg) ([]byte, error) {
	return body.Build()
}
