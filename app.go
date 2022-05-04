package main

import (
	"ToDb/communication"
	"ToDb/lib"
	"ToDb/menu"
	"context"
	"encoding/json"
	"net/http"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	menu.InitMenu(ctx)
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

// ImportConn 导入
func (a *App) ImportConn() {
	err := communication.General().ImportConn(a.ctx)
	if err != nil {
		lib.DefaultDialog(a.ctx, "错误", err.Error(), icon)
		return
	}
	lib.DefaultDialog(a.ctx, "成功", "导入成功", icon)
	// 导入成功后窗口重载
	runtime.WindowReload(a.ctx)
}

// ExportConn 导出
func (a *App) ExportConn() {
	communication.General().ExportConn(a.ctx)
}

// TestConnection 测试连接
func (a *App) TestConnection(connectionInfo string) {
	//var responseJson lib.JsonResponse
	code, message := communication.RedisPing(connectionInfo)
	if code != http.StatusOK {
		lib.DefaultDialog(a.ctx, "错误", message, icon)
	} else {
		lib.DefaultDialog(a.ctx, "成功", message, icon)
	}
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
		lib.DefaultDialog(a.ctx, "错误", message, icon)
	}
	return responseJson.String()
}

// LoadingConnKey 加载已保存的连接信息
func (a *App) LoadingConnKey() string {
	return communication.LoadingBaseHistoryInfo(a.ctx)
}

// LoadingConnInfo 获取链接信息详情
func (a *App) LoadingConnInfo(key string) string {
	code, message := communication.LoadingHistoryInfo(key)
	if code != http.StatusOK {
		lib.DefaultDialog(a.ctx, "错误", message, icon)
	}
	return message
}

// LoadingDbResource 加载数据库资源消耗
func (a *App) LoadingDbResource(key string) string {
	return communication.LoadingDbResource(key)
}

// GetNodeData 获取节点数据
func (a *App) GetNodeData(connType, connName, nodeIdStr string) string {
	sts, _ := communication.GetNodeData(connType, connName, nodeIdStr)
	return sts
}

// RedisGetData 从redis获取数据
func (a *App) RedisGetData(connType, connName, nodeIdStr, key string) string {
	_v, _ := communication.RedisGetData(connType, connName, nodeIdStr, key)
	v, _ := json.Marshal(_v)
	return string(v)
}

// RedisReName redis key重命名
func (a *App) RedisReName(connType, connName, nodeIdStr, oldKey, newKey string) {
	v := communication.RedisReName(connType, connName, nodeIdStr, oldKey, newKey)
	if v != "success" {
		lib.DefaultDialog(a.ctx, "错误", v, icon)
	} else {
		lib.DefaultDialog(a.ctx, "成功", "修改成功", icon)
	}
}

// RedisUpTtl 更新redis剩余时间
func (a *App) RedisUpTtl(connType, connName, nodeIdStr, key, ttlStr string) {
	v := communication.RedisUpTtl(connType, connName, nodeIdStr, key, ttlStr)
	if v != "success" {
		lib.DefaultDialog(a.ctx, "错误", v, icon)
	} else {
		lib.DefaultDialog(a.ctx, "成功", "修改成功", icon)
	}
}

// RedisDelKey 删除键
func (a *App) RedisDelKey(connType, connName, nodeIdStr, key string) {
	selection, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.WarningDialog,
		Title:         "删除",
		Message:       "确定删除该key吗？",
		Buttons:       []string{"确定", "取消"},
		DefaultButton: "取消",
		CancelButton:  "取消",
	})
	if selection != "确定" {
		return
	}
	v := communication.RedisDel(connType, connName, nodeIdStr, key)
	if v != "success" {
		lib.DefaultDialog(a.ctx, "错误", v, icon)
	} else {
		lib.DefaultDialog(a.ctx, "成功", "删除成功", icon)
	}
}

// RedisSaveStringValue 更新redis值
func (a *App) RedisSaveStringValue(connType, connName, nodeIdStr, key, value, ttl string) {
	err := communication.RedisUpdateStringValue(connType, connName, nodeIdStr, key, value, ttl)
	if err != nil {
		lib.DefaultDialog(a.ctx, "错误", err.Error(), icon)
	} else {
		lib.DefaultDialog(a.ctx, "成功", "修改成功", icon)
	}
}
