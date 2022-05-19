// package structs
// @Author: easymbol
// @Date: 2022/5/3
// @LastEditors: easymbol
// @LastEditTime: 2022/5/3 22:39
// @FilePath:
// @Description: 对象结构体

// Copyright (c) 2022 by easymbol, All Rights Reserved.

package structs

// ConnectionType 基础连接信息
type ConnectionType struct {
	Type     string `tag:"类型" json:"type"`      //类型
	Alias    string `tag:"连接名" json:"alias"`    //别名
	HostURL  string `tag:"连接地址" json:"hostURL"` //连接地址
	Port     string `tag:"端口号" json:"port"`     //端口号
	Username string `tag:"用户名" json:"username"` //用户名
	Password string `tag:"密码" json:"password"`  //密码
}

// BaseTreeInfo 基础连接信息
type BaseTreeInfo struct {
	Label        string `json:"label"`        //适配tree
	Title        string `json:"title"`        //别名
	ConnType     string `json:"connType"`     //类型
	IconPath     string `json:"iconPath"`     //图标路径
	ConnFileAddr string `json:"ConnFileAddr"` //连接信息文件存放地址
	//Children     string `json:"children"`	 //子节点
}

// DbTreeInfo 数据库内部信息
type DbTreeInfo struct {
	Label    string `json:"label"`    //名称
	Title    string `json:"title"`    //别名
	Key      string `json:"key"`      //键
	IsLeaf   bool   `json:"isLeaf"`   //是否是叶子节点
	Children string `json:"children"` //子节点
}

// RedisList redis List数据
type RedisList struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}

// RedisHash redis hash数据
type RedisHash struct {
	Id    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
