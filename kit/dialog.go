/*
 * @Author: symbol
 * @Date: 2022-05-22 11:21:34
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-28 17:00:39
 * @FilePath: \ToDb\kit\dialog.go
 * @Description: 对话框控制工具
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */

package kit

import (
	"ToDb/common/consts"
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
func (d *sDiaLogKit) About(ctx context.Context, icon []byte) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:          runtime.InfoDialog,
		Title:         consts.AppName,
		Message:       consts.Description,
		Icon:          icon,
		DefaultButton: consts.BtnConfirmText,
		Buttons:       []string{consts.BtnConfirmText},
	})
}
