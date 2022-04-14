package lib

import (
	"encoding/json"
	"strings"
)

type Node struct {
	Title    string  `json:"title"`
	Key      string  `json:"key"`
	Count    int     `json:"count"`
	Children []*Node `json:"children"`
}

func PackageTree(v []string) string {
	// 声明切片存放树形数据
	var treeNode []Node
	for _, val := range v {
		var node Node
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
			GetChildren(sl[1], &node)
		}

		if flag {
			treeNode = append(treeNode, node)
		}
	}
	res, err := json.Marshal(treeNode)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func GetChildren(nodeStr string, parentNode *Node) {
	node := &Node{}
	sl := strings.SplitN(nodeStr, ":", 2)
	for _, v := range parentNode.Children {
		if len(sl) > 1 && v.Key == sl[0] {
			node = v
		}
	}
	flag := node.Key == ""
	if flag {
		node.Title = sl[0]
		var sb strings.Builder
		sb.WriteString(sl[0])
		sb.WriteString(":*")
		node.Key = sb.String()
		node.Count = 0
	}

	if len(sl) > 1 {
		GetChildren(sl[1], node)
	}
	if flag {
		parentNode.Children = append(parentNode.Children, node)
		parentNode.Count = len(parentNode.Children)
	}
}
