package elasticsearch

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/olivere/elastic"
	"github.com/yogasab/bookstore_items-api/app/logger"
)

var (
	ESClient esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(index string, body interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

// func NewESClient(client *elasticsearch.Client) esClientInterface {
// 	return &esClient{client: client}
// }
// func Init() (*elasticsearch.Client, error) {
// 	cfg := elasticsearch.Config{
// 		Addresses: []string{"http://localhost:9200"},
// 		Transport: &http.Transport{
// 			MaxIdleConnsPerHost:   10,
// 			ResponseHeaderTimeout: time.Millisecond,
// 			// DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
// 			// TLSClientConfig: &tls.Config{
// 			// 	MinVersion: tls.VersionTLS12,
// 			// 	// ...
// 			// },
// 		},
// 	}
// 	es, err := elasticsearch.NewClient(cfg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Println(es.Info())
// 	return es, nil
// }

func Init() {
	logger := logger.NewLogger()
	es, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9206"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(logger),
		elastic.SetInfoLog(logger),
	)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Elastisearch node connected successfully", es)
	ESClient.setClient(es)
}

func (es *esClient) setClient(client *elastic.Client) {
	es.client = client
}

func (ec *esClient) Index(index string, body interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	response, err := ec.client.
		Index().
		Index(index).
		BodyJson(body).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return response, nil
}
