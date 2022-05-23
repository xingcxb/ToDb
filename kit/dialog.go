/*
 * @Author: symbol
 * @Date: 2022-05-22 11:21:34
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-23 14:51:19
 * @FilePath: /todb/kit/dialog.go
 * @Description: 对话框控制工具
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */

package kit

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	insDiaLogKit = sDiaLogKit{}
)

type sDiaLogKit struct {
	ctx context.Context
}

func DiaLogKit() *sDiaLogKit {
	return &insDiaLogKit
}

// DefaultDialog 显示默认的对话框
func (d *sDiaLogKit) DefaultDialog(ctx context.Context, title, message string, icon []byte) {
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

// AboutDialog 关于对话框
func (d *sDiaLogKit) About(ctx context.Context) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type: runtime.InfoDialog,
		// Title: ,
	})
}
