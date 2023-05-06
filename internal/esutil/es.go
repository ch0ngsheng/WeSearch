package esutil

type docCreate struct {
	Content string `json:"content"`
}

type createDocResp struct {
	Result string `json:"result"`
}

type searchResp struct {
	TimeOut bool `json:"timed_out"`

	Hits struct {
		Hits []searchRespHit `json:"hits"`
	} `json:"hits"`
}

type searchRespHit struct {
	Id    string  `json:"_id"`
	Score float64 `json:"_score"`
}
