package redisKit

import (
	"ToDb/lib"
	"context"
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	Addr = "127.0.0.1"
	Port = "6379"
	Password = "123456"
	InitDb()
	fmt.Println(GetMainViewInfo(context.Background()))
}

func TestInfo(t *testing.T) {
	Addr = "127.0.0.1"
	Port = "6379"
	Password = "123456"
	InitDb()
	ChangeDb(context.Background(), 13)
	//fmt.Println(GetBaseAllInfo(context.Background()))
	v, _ := GetDbKeys(context.Background(), 0)
	fmt.Println(v)
	_v := lib.PackageTree(v)
	fmt.Println(_v)
}
