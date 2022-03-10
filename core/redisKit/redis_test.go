package redisKit

import (
	"context"
	"testing"
)

func TestConnect(t *testing.T) {
	Addr = "127.0.0.1"
	Port = "6739"
	Password = "123456"
	InitDb()
	GetBaseInfo(context.Background())
}
