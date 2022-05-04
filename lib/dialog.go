// package lib
// @Author: symbol
// @Date: 2022-05-01 22:39:52
// @LastEditTime: 2022-05-01 22:54:13
// @LastEditors: symbol
// @Description: 对话框控制工具
// @FilePath: \ToDb\lib\dialog.go

// Copyright (c) 2022 by easymbol, All Rights Reserved.

package lib

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// DefaultDialog 显示默认的对话框
func DefaultDialog(ctx context.Context, title, message string, icon []byte) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type: func(title string) runtime.DialogType {
			if title == "错误" {
				return runtime.ErrorDialog
			}
			return runtime.InfoDialog
		}(title),
		Icon:          icon,
		Title:         title,
		Message:       message,
		Buttons:       []string{"确定"},
		DefaultButton: "确定",
	})
}
