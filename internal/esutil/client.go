package esutil

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

type Retriever interface {
	CreateDoc(context.Context, *CreateDocReq) error
	Search(context.Context, *SearchDocReq) (*SearchDocResp, error)
}

func MustNewRetriever(addr []string, apiKey string, index string) Retriever {
	config := elasticsearch.Config{
		Addresses: addr,
		APIKey:    apiKey,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 3 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			DisableKeepAlives: false,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // self-signed certificate
			},
		},
	}

	es, err := elasticsearch.NewTypedClient(config)
	if err != nil {
		log.Fatal(err)
	}
	return &esClient{client: es, index: index}
}
