package structs

// @Author: symbol
// @Date: 2022-04-28 10:57:39
// @LastEditors: symbol
// @LastEditTime: 2022-04-30 17:02:36
// @FilePath: /todb/model/redis.go
// @Description: redis返回信息结构体

// Copyright (c) 2022 by symbol, All Rights Reserved.

// GetValue	redis的详情页面返回结构体
type GetValue struct {
	Type       string      `json:"type"`       // value类型
	Key        string      `json:"key"`        // key键
	Ttl        string      `json:"ttl"`        // ttl剩余时间
	Value      interface{} `json:"value"`      // value值
	Size       int         `json:"size"`       // value大小
	CommandStr string      `json:"commandStr"` // redis命令
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
