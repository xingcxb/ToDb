package lib

import "strings"

// TreeKeys 节点键位
type TreeKeys struct {
	Title    string    `json:"title"`    //别名
	Key      string    `json:"key"`      //key
	Count    string    `json:"count"`    //数量
	TreeKeys *TreeKeys `json:"children"` //子集
}

// KeyToTree value
func KeyToTree(value string, treeKeysRoots *TreeKeys) {
	// 获取切片数据
	title := treeKeysRoots.Title
	if title == "" {
		//表示为根节点
		treeKeysRoots = &TreeKeys{
			Title:    value,
			Key:      value,
			Count:    "1",
			TreeKeys: nil,
		}
	}
	// 判断value中是否还能够继续分割
	index := strings.Index(value, ":")
	if index == -1 {
		// 表示不能继续分割
		treeKeysRoots.TreeKeys = &TreeKeys{
			Title:    value,
			Key:      value,
			Count:    "1",
			TreeKeys: nil,
		}
		return
	}
	// 分割
	key := value[index+1:]
	treeKeysRoots.TreeKeys = &TreeKeys{
		Title:    key,
		Key:      key,
		Count:    "1",
		TreeKeys: nil,
	}
	// 递归
	KeyToTree(value[:index], treeKeysRoots)
	//fmt.Println(treeKeysRoots)
}
