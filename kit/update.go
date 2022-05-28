/*
 * @Author: symbol
 * @Date: 2022-05-28 23:13:40
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-28 23:48:38
 * @FilePath: \ToDb\kit\update.go
 * @Description:
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */
package kit

import (
	"context"
	"fmt"
)

var (
	insUpdate = sUpdate{}
)

type sUpdate struct {
	ctx context.Context
}

func Update() *sUpdate {
	return &insUpdate
}

const (
	githubUrl = "https://api.github.com/repos/xingcxb/todb/releases/latest"
)

// 检查是否存在更新
func (u *sUpdate) CheckUpdate() {
	//TODO 未发布版本暂缓实现
	fmt.Println(githubUrl)
}
