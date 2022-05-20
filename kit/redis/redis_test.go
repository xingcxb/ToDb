/*
 * @Author: symbol
 * @Date: 2022-05-04 10:22:53
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-20 17:54:11
 * @FilePath: /todb/kit/redis/redis_test.go
 * @Description:
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */
package redisKit

import (
	"ToDb/kit"
	"context"
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	Addr = "127.0.0.1"
	Port = "6379"
	Password = "123456"
	InitDb()
	fmt.Println(Redis().GetMainViewInfo(context.Background()))
}

func TestInfo(t *testing.T) {
	Addr = "127.0.0.1"
	Port = "6379"
	Password = "123456"
	InitDb()
	Redis().ChangeDb(context.Background(), 13)
	//fmt.Println(GetBaseAllInfo(context.Background()))
	v, _ := Redis().GetDbKeys(context.Background(), 0)
	fmt.Println(v)
	_v := kit.StrKit().PackageTree(v)
	fmt.Println(_v)
}

func TestDel(t *testing.T) {
	Addr = "127.0.0.1"
	Port = "6379"
	Password = "123456"
	InitDb()
	Redis().ChangeDb(context.Background(), 13)
	Redis().Del(context.Background(), "1111")
}

func TestStream(t *testing.T) {
	Addr = "127.0.0.1"
	Port = "6379"
	Password = "123456"
	InitDb()
	Redis().ChangeDb(context.Background(), 13)
	Redis().GetStreamValue(context.Background(), "1:2:stream")
}

func TestZSet(t *testing.T) {
	Addr = "127.0.0.1"
	Port = "6379"
	Password = "123456"
	InitDb()
	Redis().ChangeDb(context.Background(), 13)
	Redis().GetZSetCount(context.Background(), "1:2:zset")
}
