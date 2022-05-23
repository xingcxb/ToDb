/*
 * @Author: symbol
 * @Date: 2022-05-22 11:21:34
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-23 11:57:23
 * @FilePath: /todb/menu/winMenu.go
 * @Description:
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */
package menu

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	insWinMenu = sWinMenu{}
)

type sWinMenu struct {
	ctx context.Context
}

func WinMenu() *sWinMenu {
	return &insWinMenu
}

// File win文件
func (s *sWinMenu) File(ctx context.Context) *menu.MenuItem {
	return menu.SubMenu("文件",
		menu.NewMenuFromItems(
			menu.SubMenu("新建连接",
				menu.NewMenuFromItems(
					menu.Text("Redis...", keys.CmdOrCtrl("N"), func(data *menu.CallbackData) {
						runtime.EventsEmit(ctx, "createConn", "redis")
					}),
					// menu.Text("MySQL...", nil, func(data *menu.CallbackData) {
					// 	runtime.EventsEmit(ctx, "createConn", "mysql")
					// }),
				),
			),
			menu.Text("新建查询", nil, nil),
			menu.Separator(),
			menu.Text("导入连接...", keys.CmdOrCtrl("I"), func(data *menu.CallbackData) {
				runtime.EventsEmit(ctx, "importConn")
			}),
			menu.Text("导出连接...", nil, func(data *menu.CallbackData) {
				runtime.EventsEmit(ctx, "exportConn")
			}),
			menu.Separator(),
			menu.Text("退出", keys.CmdOrCtrl("Q"), func(cd *menu.CallbackData) {
				runtime.Quit(ctx)
			}),
		),
	)
}
