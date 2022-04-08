package lib

import (
	"encoding/json"
	"fmt"
	"strings"
)

// TreeKeys 节点键位
type Chilren struct {
	Title   string    `json:"title"`    //别名
	Key     string    `json:"key"`      //key
	Count   string    `json:"count"`    //数量
	Chilren []Chilren `json:"children"` //子集
}

// KeyToTree value
func KeyToTree(value string, treeKeysRoots Chilren) {
	// 判断value中是否还能够继续分割
	index := strings.Index(value, ":")
	if index == -1 {
		// 表示不能继续分割
		treeKeysRoots.Chilren = append(treeKeysRoots.Chilren, Chilren{
			Title:   value,
			Key:     value,
			Count:   "1",
			Chilren: nil,
		})
		sb, _ := json.Marshal(treeKeysRoots)
		fmt.Println(string(sb))
		return
	}
	// 获取切片数据
	title := treeKeysRoots.Title
	if title == "" {
		//表示为根节点
		treeKeysRoots = Chilren{
			Title:   value[:index],
			Key:     value,
			Count:   "1",
			Chilren: nil,
		}
	}
	// 分割
	key := value[index+1:]
	treeKeysRoots.Chilren = append(treeKeysRoots.Chilren, Chilren{
		Title:   value[:index],
		Key:     key,
		Count:   "1",
		Chilren: nil,
	})

	// 递归
	KeyToTree(value[:index], treeKeysRoots)
	//fmt.Println(treeKeysRoots)
}

func KeyToTree2(value []string, tree []interface{}) {

}
