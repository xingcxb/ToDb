package main

import (
	"ToDb/appview"
	"ToDb/communication"
	"ToDb/lib"
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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
func (a *App) TestConnection(connectionInfo string) {
	//var responseJson lib.JsonResponse
	code, message := communication.RedisPing(connectionInfo)
	//responseJson = lib.JsonResponse{
	//	Code:    code,
	//	Message: message,
	//}
	title := "成功"
	typeV := runtime.InfoDialog
	if code != http.StatusOK {
		title = "错误"
		typeV = runtime.ErrorDialog
	}
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          typeV,
		Title:         title,
		Message:       message,
		Buttons:       []string{"确定"},
		DefaultButton: "确定",
	})
	//return responseJson.String()
}

// Ok 确定按钮
func (a App) Ok(connectionInfo string) string {
	var responseJson lib.JsonResponse
	code, message := communication.Ok(a.ctx, connectionInfo)
	responseJson = lib.JsonResponse{
		Code:    code,
		Message: message,
	}
	if code != http.StatusOK {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "错误",
			Message:       message,
			Buttons:       []string{"确定"},
			DefaultButton: "确定",
		})
	}
	return responseJson.String()
}

// LoadingConnKey 加载已保存的连接信息
func (a *App) LoadingConnKey() string {
	return communication.LoadingBaseHistoryInfo()
}

// LoadingConnInfo 获取链接信息详情
func (a *App) LoadingConnInfo(key string) string {
	code, message := communication.LoadingHistoryInfo(key)
	if code != http.StatusOK {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "错误",
			Message:       message,
			Buttons:       []string{"确定"},
			DefaultButton: "确定",
		})
	}
	return message
}

// LoadingDbResource 加载数据库资源消耗
func (a *App) LoadingDbResource(key string) string {
	return communication.LoadingDbResource(key)
}

// GetNodeData 获取节点数据
func (a *App) GetNodeData(connType, connName string, nodeId int) string {
	return communication.GetNodeData(connType, connName, nodeId)
}
