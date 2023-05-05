package main

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func main() {
	config := elasticsearch.Config{
		Addresses: []string{"https://192.168.0.140:9200"},
		APIKey:    "MGlXUzBvY0J6OXpERXFiQUd6ZEs6NDlNTlRiXzlSbHFXa2c1UWJESndGdw==",
		//CertificateFingerprint: caFingerPrint,
		// self-signed certificate
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	es, err := elasticsearch.NewTypedClient(config)
	if err != nil {
		log.Fatal(err)
	}

	testSearch(es)

	//testAdd(es)
}

func testAdd(es *elasticsearch.TypedClient) {
	document := struct {
		Content string `json:"content"`
	}{
		Content: "Foo",
	}

	res, err := es.Index("wxindex").Id("1005").
		Request(document).
		Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.StatusCode)
	all, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(all))
}

func testSearch(es *elasticsearch.TypedClient) {

	res, err := es.Search().
		Index("wxindex").
		Request(&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"content": {Query: "springboot"},
				},
			},
		}).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.StatusCode)

	all, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	log.Println(string(all))
}
