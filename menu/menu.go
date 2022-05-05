package menu

import (
	"ToDb/lib"
	"context"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	goruntime "runtime"
)

var (
	osInfo = ""
)

// InitMenu 初始化菜单目录
func InitMenu(ctx context.Context) {
	var appMenu *menu.Menu
	// 判断系统类型
	osInfo = goruntime.GOOS
	switch osInfo {
	case "windows":
		//Windows目录
		appMenu = windows(ctx)
	case "darwin":
		//macos
		appMenu = macOs(ctx)
	default:
		//linux
		lib.DefaultDialog(ctx, "错误", "暂不支持该系统", nil)
	}
	runtime.MenuSetApplicationMenu(ctx, appMenu)
}

// macOS菜单
func macOs(ctx context.Context) *menu.Menu {
	return menu.NewMenuFromItems(
		macOsAbout(ctx),
		files(ctx),
		edit(ctx),
	)
}

// macOs关于
func macOsAbout(ctx context.Context) *menu.MenuItem {
	aboutMenu := menu.AppMenu()
	aboutMenu.Prepend(menu.Text("关于", nil, nil))
	return aboutMenu
}

// windows菜单
func windows(ctx context.Context) *menu.Menu {
	return menu.NewMenuFromItems(
		files(ctx),
		edit(ctx),
		about(ctx),
	)
}

// 文件菜单
func files(ctx context.Context) *menu.MenuItem {
	if osInfo == "darwin" {
		return menu.SubMenu("文件",
			menu.NewMenuFromItems(
				menu.SubMenu("新建连接",
					menu.NewMenuFromItems(
						menu.Text("Redis...", nil, nil),
						menu.Text("MySQL...", nil, nil),
					),
				),
				menu.Text("新建查询", nil, nil),
				menu.Separator(),
			),
		)
	} else {
		// 非darwin系统
		return menu.SubMenu("文件",
			menu.NewMenuFromItems(
				menu.SubMenu("新建连接",
					menu.NewMenuFromItems(
						menu.Text("Redis...", nil, nil),
						menu.Text("MySQL...", nil, nil),
					),
				),
				menu.Text("新建查询", nil, nil),
				menu.Separator(),
				menu.Text("退出", keys.CmdOrCtrl("Q"), nil),
			),
		)
	}
}

//编辑
func edit(ctx context.Context) *menu.MenuItem {
	return menu.SubMenu("编辑",
		menu.NewMenuFromItems(
			menu.Text("复制", keys.CmdOrCtrl("C"), nil),
			menu.Text("粘贴", keys.CmdOrCtrl("V"), nil),
			menu.Separator(),
			menu.Text("全选", keys.CmdOrCtrl("A"), nil),
		),
	)
}

// 关于
func about(ctx context.Context) *menu.MenuItem {
	if osInfo == "darwin" {
		return menu.SubMenu("ToDb",
			menu.NewMenuFromItems(

				menu.Text("帮助中心", nil, nil),
				menu.SubMenu("在线文档",
					menu.NewMenuFromItems(
						menu.Text("Redis", nil, func(_ *menu.CallbackData) {
							runtime.BrowserOpenURL(ctx, "http://doc.redisfans.com/")
						}),
						menu.Text("MySQL", nil, func(_ *menu.CallbackData) {
							runtime.BrowserOpenURL(ctx, "https://dev.mysql.com/doc/refman/8.0/en/")
						}),
					),
				),
				menu.Text("意见反馈", nil, nil),
				menu.Separator(),
				menu.Text("检查更新", nil, nil),
				menu.Separator(),
				//menu.Text("关于", nil, nil),
			),
		)
	} else {
		return menu.SubMenu("ToDb帮助",
			menu.NewMenuFromItems(
				menu.Text("帮助中心", nil, nil),
				menu.SubMenu("在线文档",
					menu.NewMenuFromItems(
						menu.Text("Redis", nil, func(_ *menu.CallbackData) {
							runtime.BrowserOpenURL(ctx, "http://doc.redisfans.com/")
						}),
						menu.Text("MySQL", nil, func(_ *menu.CallbackData) {
							runtime.BrowserOpenURL(ctx, "https://dev.mysql.com/doc/refman/8.0/en/")
						}),
					),
				),
				menu.Text("意见反馈", nil, nil),
				menu.Separator(),
				menu.Text("检查更新", nil, nil),
				menu.Separator(),
				menu.Text("关于", nil, nil),
			),
		)
	}
}
