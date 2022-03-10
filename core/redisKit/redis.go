package redisKit

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
)

var (
	Addr     = "" //redis链接地址
	Password = "" //密码
	Port     = "" //端口号
	Db       = 0  //操作数据库
	ctx      = context.Background()
	rdb      *redis.Client
)

func InitDb() {
	var url strings.Builder
	url.WriteString(Addr)
	url.WriteString(":")
	url.WriteString(Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     url.String(),
		Password: Password,
		DB:       Db,
	})
}

// GetBaseInfo 获取redis基础信息
func GetBaseInfo(ctx context.Context) {
	//fmt.Println(rdb.Set(ctx, "test", "11111", time.Second*100).Err())
	fmt.Println(rdb.Ping(ctx))
	//fmt.Println(rdb.Do(ctx, "info"))
}
