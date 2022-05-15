package main

import (
	"ToDb/communication"
	"ToDb/kit"
	"ToDb/menu"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"net/http"
	"strings"

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
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", err.Error(), icon)
		return
	}
	kit.DiaLogKit().DefaultDialog(a.ctx, "成功", "导入成功", icon)
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
	code, message := communication.Redis().RedisPing(connectionInfo)
	if code != http.StatusOK {
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", message, icon)
	} else {
		kit.DiaLogKit().DefaultDialog(a.ctx, "成功", message, icon)
	}
}

// Ok 确定按钮进行的操作
func (a App) Ok(connectionInfo string) string {
	var responseJson kit.JsonResponse
	code, message := communication.Redis().Ok(a.ctx, connectionInfo)
	responseJson = kit.JsonResponse{
		Code:    code,
		Message: message,
	}
	if code != http.StatusOK {
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", message, icon)
	}
	return responseJson.String()
}

// LoadingConnKey 加载已保存的连接信息
func (a *App) LoadingConnKey() string {
	return communication.Redis().LoadingBaseHistoryInfo(a.ctx)
}

// LoadingConnInfo 获取链接信息详情
func (a *App) LoadingConnInfo(dbType, fileName string) string {
	var code int
	var message string
	switch dbType {
	case "redis":
		code, message = communication.Redis().LoadingHistoryInfo(fileName)
	default:
		code = http.StatusNotFound
		message = "暂不支持的数据库类型"
	}
	if code != http.StatusOK {
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", message, icon)
		return ""
	}
	return message
}

// LoadingDbResource 加载数据库资源消耗信息
func (a *App) LoadingDbResource(key string) string {
	return communication.Redis().LoadingDbResource(a.ctx, key)
}

// ChangeRightWindowStyle 改变右侧窗口样式
func (a *App) ChangeRightWindowStyle(parentNode, nextParentNode, node string) string {
	// 连接类型
	connType := gjson.Get(parentNode, "connType").String()
	// 文件名
	fileName := gjson.Get(parentNode, "title").String()
	switch connType {
	case "redis":
	default:
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", "暂不支持的数据库类型", icon)
	}
	// 获取db信息
	dbId := gjson.Get(nextParentNode, "key").String()
	// 获取指定的节点key
	fullKey := gjson.Get(node, "fullStr").String()
	// 指定的key数据类型
	dataType, err := communication.Redis().GetValueType(a.ctx, fileName, dbId, fullKey)
	fmt.Println(dataType)
	if err != nil {
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", err.Error(), icon)
	}
	return strings.ToLower(dataType)
}

// GetNodeData 获取节点数据
func (a *App) GetNodeData(connType, connName, nodeIdStr string) string {
	strs := ""
	var err error
	switch connType {
	case "redis":
		// 获取redis节点数据
		strs, err = communication.Redis().GetNodeData(a.ctx, connType, connName, nodeIdStr)
	default:
		err = errors.New("暂不支持此连接类型")
	}
	if err != nil {
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", err.Error(), icon)
	}
	if strs == "null" {
		strs = ""
	}
	return strs
}

// RedisGetData 从redis获取数据
func (a *App) RedisGetData(connType, connName, nodeIdStr, key string) string {
	v := ""
	switch connType {
	case "redis":
		// 获取redis节点数据
		getValue, _ := communication.Redis().RedisGetData(a.ctx, connType, connName, nodeIdStr, key)
		_v, _ := json.Marshal(getValue)
		v = string(_v)
	default:
		v = "暂不支持"
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", v, icon)
	}
	return v
}

// RedisReName redis key重命名
func (a *App) RedisReName(connType, connName, nodeIdStr, oldKey, newKey string) {
	v := ""
	switch connType {
	case "redis":
		v = communication.Redis().RedisReName(a.ctx, connType, connName,
			nodeIdStr, oldKey, newKey)
	default:
		v = "暂不支持"
	}
	if v != "success" {
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", v, icon)
	} else {
		kit.DiaLogKit().DefaultDialog(a.ctx, "成功", "修改成功", icon)
	}
}

// RedisUpTtl 更新redis剩余时间
func (a *App) RedisUpTtl(connType, connName, nodeIdStr, key, ttlStr string) {
	v := ""
	switch connType {
	case "redis":
		v = communication.Redis().RedisUpTtl(a.ctx, connType, connName, nodeIdStr, key, ttlStr)
	default:
		v = "暂不支持"
	}
	if v != "success" {
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", v, icon)
	} else {
		kit.DiaLogKit().DefaultDialog(a.ctx, "成功", "修改成功", icon)
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
	v := ""
	switch connType {
	case "redis":
		v = communication.Redis().RedisDel(a.ctx, connType, connName, nodeIdStr, key)
	default:
		v = "暂不支持"
	}
	if v != "success" {
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", v, icon)
	} else {
		kit.DiaLogKit().DefaultDialog(a.ctx, "成功", "删除成功", icon)
	}
}

// RedisSaveStringValue 更新redis值
func (a *App) RedisSaveStringValue(connType, connName, nodeIdStr, key, value, ttl string) {
	var err error
	switch connType {
	case "redis":
		err = communication.Redis().RedisUpdateStringValue(a.ctx, connType, connName,
			nodeIdStr, key, value, ttl)
	default:
		err = errors.New("暂不支持")
	}
	if err != nil {
		kit.DiaLogKit().DefaultDialog(a.ctx, "错误", err.Error(), icon)
	} else {
		kit.DiaLogKit().DefaultDialog(a.ctx, "成功", "修改成功", icon)
	}
}
