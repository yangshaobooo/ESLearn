package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200", // Elasticsearch 服务器的地址
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("创建 Elasticsearch 客户端失败：%s", err)
	}

	// 构建索引请求
	req := esapi.IndicesCreateRequest{
		Index: "hotel",
		Body:  strings.NewReader(`{"mappings": {"properties": {"id": {"type": "keyword"},"name": {"type": "text","analyzer": "ik_max_word","copy_to": "all"},"address": {"type": "keyword","index": false},"price": {"type": "integer"},"score": {"type": "integer"},"brand": {"type": "keyword"},"city": {"type": "keyword"},"starName": {"type": "keyword"},"business": {"type": "keyword","copy_to": "all"},"location": {"type": "geo_point"},"pic": {"type": "keyword","index": false},"all": {"type": "text","analyzer": "ik_max_word"}}}}`),
	}

	// 发送索引请求
	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("索引请求失败：%s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("索引请求失败：%s", res.Status())
	} else {
		fmt.Println("索引表创建成功")
	}
}
