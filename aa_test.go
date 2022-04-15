package main

import (
	"ToDb/communication"
	"ToDb/lib"
	"context"
	"fmt"
	"testing"
)

func TestTt(t *testing.T) {
	aa := `{"alias":"11","hostURL":"11","port":"11","username":"11","password":"11","savePassword":true}`
	communication.Ok(context.Background(), aa)
}

func TestReadFile(t *testing.T) {
	//communication.LoadingBaseHistoryInfo()
	communication.LoadingHistoryInfo("这是2")
}

func TestGetHistoryInfo(t *testing.T) {
	fmt.Println(communication.LoadingHistoryInfo("这是2"))
}

func TestFilePath(t *testing.T) {
	fmt.Println(communication.LoadingBaseHistoryInfo())
}

func TestTest(t *testing.T) {
	//redis返回数据: [1:2:3, 1:2:4, 1111, 12312]
	//sl := []string{"1:2:3:9:5", "1:2:4:6", "1111", "12312:333"}
	sl := []string{"234234", "1:2:4", "1:2:3", "1111:33333", "12312", "1:2:5", "1111:2222"}
	v := lib.PackageTree(sl)
	fmt.Printf("%s\n", v)
}

func TestMP(t *testing.T) {
	//v := []int{0, 2, 2, 0, 0}
	sl := []string{"234234", "1:2:4", "1:2:3", "1111:33333", "12312", "1:2:5", "1111:2222"}
	fmt.Println(lib.BubbleDescSort(sl))
}

func TestApp_GetNodeData(t *testing.T) {
	fmt.Println(communication.GetNodeData("redis", "localhost", "13"))
}
