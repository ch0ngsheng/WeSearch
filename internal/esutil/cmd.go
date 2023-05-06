package esutil

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

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
	DocIDList []string
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
		log.Printf("fail to read body, err: %v\n", err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		log.Printf("response status %d not expected, %s", res.StatusCode, respBody)
		return fmt.Errorf("response status not expected, %s", respBody)
	}

	createResp := &createDocResp{}
	if err = json.Unmarshal(respBody, createResp); err != nil {
		log.Printf("failed to unmarshal resp body when create doc %s", respBody)
		return err
	}

	log.Printf("doc created result %s\n", createResp.Result)
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
		log.Printf("failed to search %v\n", err)
		return nil, err
	}

	searchRes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("read search body, err: %v\n", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("seach status not ok, %s", searchRes)
		return nil, errors.New("search failed")
	}

	resp := &searchResp{}
	if err = json.Unmarshal(searchRes, resp); err != nil {
		log.Printf("failed to unmarshal search resp %s, %v", searchRes, err)
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

func buildRespDocIDList(resp *searchResp) []string {
	res := make([]string, 0, len(resp.Hits.Hits))
	for _, r := range resp.Hits.Hits {
		res = append(res, r.Id)
	}
	return res
}
