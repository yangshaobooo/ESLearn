package server

import (
	"ESLearn/dao/mysql"
	"ESLearn/model"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

// 创建es客户端
func TestNewEsClient(t *testing.T) {
	// 创建新的es客户端连接
	_, err := NewEsClient()
	if err != nil {
		log.Fatalf("创建 Elasticsearch 客户端失败：%s", err)
	} else {
		fmt.Println("创建新的es客户端连接成功")
	}
}

// 创建一个索引表
func TestCreateIndex(t *testing.T) {
	client, err := NewEsClient()
	if err != nil {
		log.Fatalf("创建 Elasticsearch 客户端失败：%s", err)
	}
	if err = client.CreateIndex("hotel", model.HotelTable); err != nil {
		log.Fatalf("创建es表失败 %s", err)
	}
	fmt.Println("创建hotel索引表成功")
}

// 增加一个文档
func TestInsertOne(t *testing.T) {
	err := mysql.Init()
	client, err := NewEsClient()
	if err != nil {
		log.Fatalf("创建 Elasticsearch 客户端失败：%s", err)
	}
	// 获取数据库中的数据
	hotel := &model.HotelSql{
		ID: 36934,
	}
	if err = mysql.GetHotelOne(hotel); err != nil {
		fmt.Printf("get data failed:%v\n", err)
		return
	}
	// 数据库中数据转化为es索引表中数据
	toEsHotel := model.HotelSqlToHotel(*hotel)
	// 转化为json
	data, err := json.Marshal(toEsHotel)
	if err != nil {
		fmt.Printf("toJson failed %v\n", err)
		return
	}
	// 把数据插入es
	if err = client.InsertOne("hotel", string(data)); err != nil {
		fmt.Printf("数据插入es失败 %v\n", err)
	}
	fmt.Println("数据插入es成功")
}

// 查询一个文档
func TestGetOneDoc(t *testing.T) {
	err := mysql.Init()
	client, err := NewEsClient()
	if err != nil {
		log.Fatalf("创建 Elasticsearch 客户端失败：%s", err)
	}
	doc, err := client.GetOneDoc("hotel", "1")
	if err != nil {
		fmt.Printf("查询数据失败:%v\n", err)
	}
	fmt.Printf(string(doc))
}

// 删除一个文档
func TestDeleteDocument(t *testing.T) {
	err := mysql.Init()
	client, err := NewEsClient()
	if err != nil {
		log.Fatalf("创建 Elasticsearch 客户端失败：%s", err)
	}
	err = client.DeleteDocument("hotel", "1")
	if err != nil {
		fmt.Printf("查询数据失败:%v\n", err)
	}
	fmt.Println("删除完成")
}

// 批量插入
func TestBulkIndexDocuments(t *testing.T) {
	err := mysql.Init()
	client, err := NewEsClient()
	if err != nil {
		log.Fatalf("创建 Elasticsearch 客户端失败：%s", err)
	}
	data, err := mysql.GetHotel()
	if err != nil {
		fmt.Printf("getHotel failed: %v\n", err)
		return
	}
	for _, d := range data {
		h := model.HotelSqlToHotel(d)
		marshal, _ := json.Marshal(h)
		err = client.InsertOne("hotel", string(marshal))
		if err != nil {
			fmt.Printf("添加数据失败：%v\n", err)
			continue
		}
	}
	fmt.Println("end")
}
