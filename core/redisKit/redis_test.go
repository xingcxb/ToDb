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
	Addr = "127.0.0.1"
	Port = "6379"
	Password = "redmiK40"
	InitDb()
	fmt.Println(GetMainViewInfo(context.Background()))
}
