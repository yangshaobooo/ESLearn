package main

import (
	"ESLearn/dao/mysql"
	"fmt"
)

func main() {
	err := mysql.Init()
	if err != nil {
		fmt.Printf("mysql 启动失败：%v\n", err)
	}
}
