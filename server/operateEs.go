package server

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strconv"
	"strings"
)

type esClient struct {
	Client     *elasticsearch.Client
	documentId int
}

// NewEsClient 初始化es客户端
func NewEsClient() (*esClient, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200", // Elasticsearch 服务器的地址
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	result := &esClient{
		Client: client,
	}
	return result, nil
}

// CreateIndex 创建索引表
func (es *esClient) CreateIndex(name, body string) error {
	req := esapi.IndicesCreateRequest{
		Index: name,
		Body:  strings.NewReader(body),
	}
	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		log.Fatalf("索引请求失败：%s", err)
	}
	defer res.Body.Close()
	return nil
}

// InsertOne 插入一条doc
func (es *esClient) InsertOne(name, body string) error {
	es.documentId++
	req := esapi.IndexRequest{
		Index:      name,
		DocumentID: strconv.Itoa(es.documentId),
		Body:       strings.NewReader(body),
	}
	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
