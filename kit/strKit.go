package kit

import (
	"context"
	"encoding/json"
	"strings"
)

var (
	insStrKit = sStrKit{}
)

type sStrKit struct {
	ctx context.Context
}

func StrKit() *sStrKit {
	return &insStrKit
}

type Node struct {
	Label    string  `json:"label"`
	Key      string  `json:"key"`
	FullStr  string  `json:"fullStr"`
	Count    int     `json:"count"`
	Children []*Node `json:"children,omitempty"`
}

func (s *sStrKit) PackageTree(v []string) string {
	// 根据数组中的中的冒号数量进行排序，最多的冒号数量的元素放在最前面
	v = s.BubbleDescSort(v)

	// 声明切片存放树形数据
	var treeNode []Node
	for _, val := range v {
		var node Node
		sl := strings.SplitN(val, ":", 2)
		// 查找treeNode切片中是否已经存在当前key
		for _, v := range treeNode {
			var sb strings.Builder
			sb.WriteString(sl[0])
			sb.WriteString(":*")
			if v.Key == sb.String() {
				node = v
			}
		}
		flag := node.Key == ""
		if flag {
			node.Label = sl[0]
			var sb strings.Builder
			sb.WriteString(sl[0])
			sb.WriteString(":*")
			node.Key = sb.String()
			node.FullStr = val
			node.Count = 0
		}
		if len(sl) > 1 {
			s.GetChildren(sl[1], val, &node)
			if !flag {
				for i, v := range treeNode {
					if v.Key == node.Key {
						treeNode[i] = node
					}
				}
			}
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

func (s *sStrKit) GetChildren(nodeStr, fullStr string, parentNode *Node) {
	node := &Node{}
	sl := strings.SplitN(nodeStr, ":", 2)
	for _, v := range parentNode.Children {
		var sb strings.Builder
		sb.WriteString(sl[0])
		sb.WriteString(":*")
		if len(sl) > 1 && v.Key == sb.String() {
			node = v
		}
	}
	flag := node.Key == ""
	if flag {
		node.Label = sl[0]
		var sb strings.Builder
		sb.WriteString(sl[0])
		sb.WriteString(":*")
		node.Key = sb.String()
		node.FullStr = fullStr
		node.Count = 0
	}

	if len(sl) > 1 {
		s.GetChildren(sl[1], fullStr, node)
	} else {
		// 如果没有子节点说明是最终节点
		node.Key = fullStr
	}
	if flag {
		parentNode.Children = append(parentNode.Children, node)
		parentNode.Count = len(parentNode.Children)
	}
}

// BubbleDescSort 冒泡排序 倒序
func (s *sStrKit) BubbleDescSort(values []string) []string {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if strings.Count(values[i], ":") < strings.Count(values[j], ":") {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values
}

// BubbleAscSort 冒泡排序 正序
func (s *sStrKit) BubbleAscSort(values []string) []string {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if strings.Count(values[i], ":") > strings.Count(values[j], ":") {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values
}

// FirstUpper 字符串首字母大写
func (s *sStrKit) FirstUpper(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToUpper(str[:1]) + str[1:]
}

// FirstLower 字符串首字母小写
func (s *sStrKit) FirstLower(str string) string {
	if str == "" {
		return ""
	}
	return strings.ToLower(str[:1]) + str[1:]
}
