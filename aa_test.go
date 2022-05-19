/*
 * @Author: symbol
 * @Date: 2022-05-04 10:22:52
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-19 17:15:31
 * @FilePath: /todb/aa_test.go
 * @Description:
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */
package main

import (
	"ToDb/kit"
	"ToDb/service"
	"context"
	"fmt"
	"testing"
)

func TestTt(t *testing.T) {
	aa := `{"alias":"11","hostURL":"11","port":"11","username":"11","password":"11","savePassword":true}`
	service.Redis().Ok(context.Background(), aa)
}

func TestReadFile(t *testing.T) {
	//communication.LoadingBaseHistoryInfo()
	service.Redis().LoadingHistoryInfo("这是2")
}

func TestGetHistoryInfo(t *testing.T) {
	fmt.Println(service.Redis().LoadingHistoryInfo("这是2"))
}

func TestFilePath(t *testing.T) {
	fmt.Println(service.Redis().LoadingBaseHistoryInfo(context.Background()))
}

func TestTest(t *testing.T) {
	//redis返回数据: [1:2:3, 1:2:4, 1111, 12312]
	//sl := []string{"1:2:3:9:5", "1:2:4:6", "1111", "12312:333"}
	sl := []string{"234234", "1:2:4", "1:2:3", "1111:33333", "12312", "1:2:5", "1111:2222"}
	v := kit.StrKit().PackageTree(sl)
	fmt.Printf("%s\n", v)
}

func TestMP(t *testing.T) {
	//v := []int{0, 2, 2, 0, 0}
	sl := []string{"234234", "1:2:4", "1:2:3", "1111:33333", "12312", "1:2:5", "1111:2222"}
	fmt.Println(kit.StrKit().BubbleDescSort(sl))
}

func TestApp_GetNodeData(t *testing.T) {
	fmt.Println(service.Redis().GetNodeData(context.Background(), "redis", "localhost", "13"))
}
