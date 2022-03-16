package main

import (
	"ToDb/appview"
	"ToDb/core/redisKit"
	"ToDb/kit"
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	a.ctx = ctx
	appview.InitMenu(ctx)
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
}

// TestConnection 测试连接
func (a *App) TestConnection(connectionInfo string) string {
	var responseJson kit.JsonResponse
	var message string
	//用于标记是否要继续匹配
	fail := false
	parameter := gjson.Parse(connectionInfo).Map()
	for k, v := range parameter {
		if k == "savePassword" || k == "username" {
			continue
		}
		if v.String() == "" {
			message = "参数" + k + "不存在值"
			fail = true
			break
		}
	}

	if !fail {
		redisKit.Addr = parameter["hostURL"].String()
		redisKit.Port = parameter["port"].String()
		redisKit.Username = parameter["username"].String()
		redisKit.Password = parameter["password"].String()
		redisKit.InitDb()
		err := redisKit.Ping(context.Background())
		if err != nil {
			message = err.Error()
		} else {
			message = "连接成功"
		}
	}
	fmt.Println(message)
	responseJson = kit.JsonResponse{
		Code:    http.StatusBadRequest,
		Message: message,
	}
	return responseJson.String()
}
