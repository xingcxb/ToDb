/*
 * @Author: symbol
 * @Date: 2022-05-22 11:21:34
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-23 13:58:26
 * @FilePath: /todb/menu/macOSMenu.go
 * @Description:
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */
package menu

import (
	"ToDb/kit"
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	insMacOSMenu = sMacOSMenu{}
)

type sMacOSMenu struct {
	ctx context.Context
}

func MacOSMenu() *sMacOSMenu {
	return &insMacOSMenu
}

// About macOS菜单关于
func (s *sMacOSMenu) About(ctx context.Context) *menu.MenuItem {
	//todo 由于目前不可变暂时使用默认
	//看了别人的的代码，可以考虑使用自定义的对话框来实现
	return menu.SubMenu("关于",
		menu.NewMenuFromItems(
			menu.Text("关于ToDb", nil, func(cd *menu.CallbackData) {
				kit.DiaLogKit().DefaultDialog()
			}),
		),
	)
}

// File macOs文件
func (s *sMacOSMenu) File(ctx context.Context) *menu.MenuItem {
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
		),
	)
}
