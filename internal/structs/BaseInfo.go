// package structs
// @Author: easymbol
// @Date: 2022/5/3
// @LastEditors: easymbol
// @LastEditTime: 2022/5/3 22:39
// @FilePath:
// @Description: TODO

// Copyright (c) 2022 by easymbol, All Rights Reserved.

package structs

// BaseConnInfo 基础连接信息
type BaseConnInfo struct {
	Title        string `json:"title"`        //别名
	Key          string `json:"key"`          //适配tree
	ConnType     string `json:"connType"`     //类型
	IconPath     string `json:"iconPath"`     //图标路径
	ConnFileAddr string `json:"ConnFileAddr"` //连接信息文件存放地址
	//Children     string `json:"children"`	 //子节点
}
