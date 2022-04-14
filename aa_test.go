package main

import (
	"ToDb/communication"
	"ToDb/lib"
	"context"
	"encoding/json"
	"fmt"
	"strings"
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
	// 声明切片存放树形数据
	var treeNode []lib.Node

	//redis返回数据: [1:2:3, 1:2:4, 1111, 12312]
	sl := []string{"234234", "1:2:4", "1:2:3", "1111", "12312"}
	for _, val := range sl {
		var node lib.Node
		sl := strings.SplitN(val, ":", 2)
		// 查找treeNode切片中是否已经存在当前key
		for _, v := range treeNode {
			if v.Key == sl[0] {
				node = v
			}
		}
		flag := node.Key == ""
		if flag {
			node.Title = sl[0]
			node.Key = sl[0]
			node.Count = 0
		}

		if len(sl) > 1 {
			lib.GetChildren(sl[1], &node)
		}

		if flag {
			treeNode = append(treeNode, node)
		}
	}
	res, err := json.Marshal(treeNode)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", res)
}

func TestApp_GetNodeData(t *testing.T) {
	fmt.Println(communication.GetNodeData("redis", "localhost", "13"))
}
