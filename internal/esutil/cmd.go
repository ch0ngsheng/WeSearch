package esutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

var (
	minScore = 0.001
)

type CreateDocReq struct {
	ID      string
	Content string
}

type SearchDocReq struct {
	DocIDList []string

	Key      string
	Keywords []string
}

type SearchDocResp struct {
	DocIDList []*SearchDocItem
}
type SearchDocItem struct {
	DocID string
	Score float32
}

type esClient struct {
	client *elasticsearch.TypedClient
	index  string
}

func (e esClient) CreateDoc(ctx context.Context, req *CreateDocReq) error {
	doc := &docCreate{
		Content: req.Content,
	}

	res, err := e.client.Index(e.index).Id(req.ID).
		Request(doc).
		Do(ctx)
	if err != nil {
		return err
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("fail to read body, err: %v", err))
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return errors.New(fmt.Sprintf("response status %d not expected, %s", res.StatusCode, respBody))
	}

	createResp := &createDocResp{}
	if err = json.Unmarshal(respBody, createResp); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to unmarshal resp body when create doc %s", respBody))
	}

	return nil
}

func (e esClient) Search(ctx context.Context, req *SearchDocReq) (*SearchDocResp, error) {
	/*
		{
		  "query": {
		    "bool": {
		      "should":[
		        {"match_phrase": {"content":"授权"}},
		        {"match_phrase": {"content":"定时"}}
		      ],
		      "filter":{
		          "ids":{"values":["1001","1002","1003","1004"]}
		      }
		    }
		  },
		  "min_score":0.001
		}
	*/
	res, err := e.client.Search().
		Index(e.index).
		Request(&search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Filter: []types.Query{
						{Ids: &types.IdsQuery{Values: req.DocIDList}},
					},
					Should: buildMatchPhrase(req),
				},
			},
			MinScore: &minScore,
		}).Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to search %v", err))
	}

	searchRes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("read search resp body, err: %v", err))
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("search failed, status %d", res.StatusCode))
	}

	resp := &searchResp{}
	if err = json.Unmarshal(searchRes, resp); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to unmarshal search resp %s, %v", searchRes, err))
	}

	list := buildRespDocIDList(resp)
	return &SearchDocResp{DocIDList: list}, nil
}

func buildMatchPhrase(req *SearchDocReq) []types.Query {
	res := make([]types.Query, 0, len(req.Keywords))

	for _, word := range req.Keywords {
		q := types.Query{
			MatchPhrase: map[string]types.MatchPhraseQuery{
				req.Key: {Query: word},
			},
		}
		res = append(res, q)
	}
	return res
}

func buildRespDocIDList(resp *searchResp) []*SearchDocItem {
	res := make([]*SearchDocItem, 0, len(resp.Hits.Hits))
	for _, r := range resp.Hits.Hits {
		item := &SearchDocItem{
			DocID: r.Id,
			Score: r.Score,
		}
		res = append(res, item)
	}
	return res
}
