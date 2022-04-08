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
	//maps := make(map[string]int, 0)
	//trees := make([]lib.TreeKeys, 0, 1)
	//trees := new(lib.TreeKeys)
	value := []string{
		"1:2:3",
		"1:2:4",
		"1111",
		"12312",
	}
	var trees []interface{}
	lib.KeyToTree2(value, trees)
	fmt.Println(trees)
}

func TestApp_GetNodeData(t *testing.T) {
	fmt.Println(communication.GetNodeData("redis", "localhost", 13))
}
