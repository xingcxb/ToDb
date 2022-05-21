/*
 * @Author: symbol
 * @Date: 2022-05-08 08:49:54
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-21 20:47:32
 * @FilePath: \ToDb\main.go
 * @Description:
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */
package main

import (
	"embed"
	"log"
	goruntime "runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// 判断系统如果是win需要增高窗口大小
	windowsHeight := 728
	osInfo := goruntime.GOOS
	if osInfo == "windows" {
		windowsHeight = 770
	}
	// Create an instance of the app structure
	// 创建一个App结构体实例
	app := NewApp()

	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:            "ToDb",
		Width:            1342,
		Height:           windowsHeight,
		DisableResize:    true,
		Assets:           assets,
		OnStartup:        app.startup,
		OnDomReady:       app.domReady,
		OnBeforeClose:    app.beforeClose,
		OnShutdown:       app.shutdown,
		WindowStartState: options.Normal,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
