package appview

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	goruntime "runtime"
)

func init() {
	var appMenu *menu.Menu
	// 判断系统类型
	osInfo := goruntime.GOOS
	switch osInfo {
	case "windows":
	//Windows目录
	default:
		//linux
	}
	fmt.Println(appMenu)
}

// InitMenu 初始化菜单目录
func InitMenu(ctx context.Context) {

}

// 文件菜单
func files() {

}
