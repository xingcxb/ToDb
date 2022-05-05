// package communication
// @Author: symbol
// @Date: 2022-05-04
// @LastEditors: symbol
// @LastEditTime: 2022-05-04 18:06:31
// @FilePath: \ToDb\communication\general.go
// @Description: 通用的通信模块

// Copyright (c) 2022 by symbol, All Rights Reserved.

package communication

import (
	"ToDb/lib/os"
	"context"
	"github.com/tidwall/gjson"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/ioutil"
	sysOs "os"
	"os/exec"
	"os/user"
	sysRuntime "runtime"
)

var (
	insGeneral = sGeneral{}
)

type sGeneral struct {
	ctx context.Context
}

func General() *sGeneral {
	return &insGeneral
}

// ImportConn 导入连接
// @param {[type]} ctx context.Context [description]
// @return error
func (s *sGeneral) ImportConn(ctx context.Context) error {
	user, err := user.Current()
	if err != nil {
		return err
	}
	defaultPath := user.HomeDir + string(sysOs.PathSeparator) + "Downloads"
	// 获取到文件的路径和错误信息
	selection, err := runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		// 默认对话框打开时显示的目录
		DefaultDirectory: defaultPath,
		// 对话框标题
		Title: "选择导入文件",
		// 默认文件名
		DefaultFilename: "localhost.json",
		// 文件过滤器列表
		Filters: []runtime.FileFilter{
			{
				DisplayName: "File (*.json)",
				Pattern:     "*.json",
			},
		},
	})
	if err != nil {
		return err
	}
	// 读取文件
	bData, err := ioutil.ReadFile(selection)
	if err != nil {
		return err
	}
	alias := gjson.Get(string(bData), "alias").String()
	return os.File().SaveFile(ctx, alias, string(bData))
}

// ExportConn 导出连接
// @param {[type]} ctx context.Context [description]
// @return error
func (s *sGeneral) ExportConn(ctx context.Context) error {
	homeDir, err := os.File().HomeDir(ctx)
	if err != nil {
		return err
	}
	// 打开文件资源管理器
	osName := sysRuntime.GOOS

	switch osName {
	case "windows":
		exec.Command("explorer", homeDir).Start()
	case "linux":
		exec.Command("xdg-open", homeDir).Start()
	case "darwin":
		exec.Command("open", homeDir).Start()
	}
	return nil
}