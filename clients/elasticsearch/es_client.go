package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string) (*elastic.GetResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetSniff(false),
	)

	if err != nil {
		panic(err.Error())
	}
	Client.setClient(client)
}

func (es *esClient) setClient(client *elastic.Client)  {
	es.client = client
}

func (es *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error)  {

	ctx := context.Background()
	res, err := es.client.
		Index().
		Index(index).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", res.Id, res.Index, res.Type)
	return res, nil
}

func (es *esClient) Get(index string, id string) (*elastic.GetResult, error)  {
	ctx := context.Background()
	result, err := es.client.Get().
		Index(index).
		Id(id).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}

	if !result.Found {
		return nil, nil
	}

	fmt.Printf("Got document %s in version %d from index %s\n", result.Id, result.Version, result.Index)
	return result, nil
}
