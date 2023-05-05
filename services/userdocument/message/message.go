package message

import "encoding/json"

const (
	Topic = "we-doc"
)

type Body struct {
	DocID   int64  `json:"doc_id"`
	Content string `json:"content"`
}

func BuildMsg(body *Body) ([]byte, error) {
	return json.Marshal(body)
}
