package menu

import (
	"ToDb/kit"
	"context"
	goruntime "runtime"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	osInfo  = ""
	insMenu = sMenu{}
)

type sMenu struct {
	ctx context.Context
}

func Menu() *sMenu {
	return &insMenu
}

// InitMenu 初始化菜单目录
func InitMenu(ctx context.Context, icon []byte) {
	var appMenu *menu.Menu
	// 判断系统类型
	osInfo = goruntime.GOOS
	switch osInfo {
	case "windows":
		//Windows目录
		appMenu = windows(ctx, icon)
	case "darwin":
		//macos
		appMenu = macOs(ctx, icon)
	default:
		//linux
		kit.DiaLogKit().DefaultDialog(ctx, "错误", "暂不支持该系统", nil)
	}
	runtime.MenuSetApplicationMenu(ctx, appMenu)
}

// macOS菜单
func macOs(ctx context.Context, icon []byte) *menu.Menu {
	return menu.NewMenuFromItems(
		MacOSMenu().About(ctx, icon),
		MacOSMenu().File(ctx),
		Menu().Edit(ctx),
		Menu().Find(ctx),
		Menu().Tools(ctx),
		Menu().Help(ctx, icon),
	)
}

// windows菜单
func windows(ctx context.Context, icon []byte) *menu.Menu {
	return menu.NewMenuFromItems(
		WinMenu().File(ctx),
		Menu().Edit(ctx),
		Menu().Find(ctx),
		Menu().Tools(ctx),
		Menu().Help(ctx, icon),
	)
}

// Edit 编辑
func (s *sMenu) Edit(ctx context.Context) *menu.MenuItem {
	return menu.SubMenu("编辑",
		menu.NewMenuFromItems(
			menu.Text("撤销", keys.CmdOrCtrl("Z"), nil),
			menu.Separator(),
			menu.Text("剪切", keys.CmdOrCtrl("X"), nil),
			menu.Text("复制", keys.CmdOrCtrl("C"), nil),
			menu.Text("粘贴", keys.CmdOrCtrl("V"), nil),
		),
	)
}

// Find 查找
func (s *sMenu) Find(ctx context.Context) *menu.MenuItem {
	return menu.SubMenu("查找",
		menu.NewMenuFromItems(
			menu.Text("运行", nil, nil),
			menu.Text("停止", nil, nil),
		),
	)
}

// Tools 工具
func (s *sMenu) Tools(ctx context.Context) *menu.MenuItem {
	return menu.SubMenu("工具",
		menu.NewMenuFromItems(
			menu.Text("历史日志", nil, nil),
			menu.Text("刷新", keys.CmdOrCtrl("R"), func(data *menu.CallbackData) {
				runtime.WindowReload(ctx)
			}),
		),
	)
}

// Help 帮助
func (s *sMenu) Help(ctx context.Context, icon []byte) *menu.MenuItem {
	return menu.SubMenu("帮助",
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
			menu.Text("关于", nil, func(cd *menu.CallbackData) {
				kit.DiaLogKit().About(ctx, icon)
			}),
		),
	)
}
