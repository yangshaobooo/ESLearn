package server

import (
	"ESLearn/dao/mysql"
	"ESLearn/model"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestNewEsClient(t *testing.T) {
	// 创建新的es客户端连接
	_, err := NewEsClient()
	if err != nil {
		log.Fatalf("创建 Elasticsearch 客户端失败：%s", err)
	} else {
		fmt.Println("创建新的es客户端连接成功")
	}
}

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

func TestInsertOne(t *testing.T) {
	err := mysql.Init()
	client, err := NewEsClient()
	if err != nil {
		log.Fatalf("创建 Elasticsearch 客户端失败：%s", err)
	}
	// 获取数据库中的数据
	hotel := &model.HotelSql{
		ID: 1,
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
