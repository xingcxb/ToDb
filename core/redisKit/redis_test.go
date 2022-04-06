package redisKit

import (
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
	Addr = "192.168.10.59"
	Port = "6379"
	Password = "123456"
	InitDb()
	ChangeDb(context.Background(), 3)
	//fmt.Println(GetBaseAllInfo(context.Background()))
	fmt.Println(GetDbData(context.Background(), 0))
}
