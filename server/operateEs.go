package server

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io/ioutil"
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

// GetOneDoc 查询一条文档
func (es *esClient) GetOneDoc(name, docId string) ([]byte, error) {
	req := esapi.GetRequest{
		Index:      name,
		DocumentID: docId,
	}
	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, fmt.Errorf("查询请求失败：%s", res.Status())
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应内容失败：%s", err)
	}

	return data, nil
}

// UpdateDocument 更新一个文档  全量更新
func (es *esClient) UpdateDocument(index, documentID string, updateBody string) error {
	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       strings.NewReader(updateBody),
	}

	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("更新请求失败：%s", res.Status())
	}

	return nil
}

// DeleteDocument 删除一个文档
func (es *esClient) DeleteDocument(index, documentID string) error {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: documentID,
	}
	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("删除请求失败：%s", res.Status())
	}
	return nil
}
