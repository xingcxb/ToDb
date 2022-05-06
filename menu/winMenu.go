package menu

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/menu"
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
					menu.Text("Redis...", nil, nil),
					menu.Text("MySQL...", nil, nil),
				),
			),
			menu.Text("新建查询", nil, nil),
			menu.Separator(),
			menu.Text("导入连接...", nil, nil),
			menu.Text("导出连接...", nil, nil),
			menu.Separator(),
			menu.Text("退出", nil, nil),
		),
	)
}
