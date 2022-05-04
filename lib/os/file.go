// package os
// @Author: symbol
// @Date: 2022-05-03
// @LastEditors: symbol
// @LastEditTime: 2022-05-03 18:06:31
// @FilePath: \ToDb\lib\os\file.go
// @Description: 针对操作系统文件/文件夹进行封装

// Copyright (c) 2022 by symbol, All Rights Reserved.

package os

import (
	"context"
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
)

var (
	insFile  = sFile{}
	lastPath = "/.ToDb"
)

type sFile struct {
	ctx context.Context
}

func File() *sFile {
	return &insFile
}

// HomeDir
// 获取系统当前使用的用户名
// @param {[type]} ctx context.Context
// @return string,error
func (o *sFile) HomeDir(ctx context.Context) (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.HomeDir + lastPath + "/", nil
}

// CreateFile 创建文件夹或文件
// @param {[type]} ctx context.Context [description]
// @param {string} fileName 文件名 如果文件名不存在则创建文件夹
// @param {bool}   hide 	是否隐藏 true 隐藏 false 不隐藏
// @return error
func (o *sFile) CreateFile(ctx context.Context, fileName string, hide bool) error {
	// 文件夹路径
	folderPath, err := o.HomeDir(ctx)
	if err != nil {
		return err
	}
	// 判断文件夹是否存在
	if !o.Exists(ctx, folderPath) {
		// 文件夹不存在，创建文件夹，并且对文件夹设置隐藏属性
		err = os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if fileName != "" {
		// 文件名如果存在，就创建文件
		filePath := folderPath + fileName + ".json"
		if !o.Exists(ctx, filePath) {
			// 创建文件
			file, err := os.Create(filePath)
			if err != nil {
				return err
			}
			defer file.Close()
		}
	}
	return nil
}

// Exists 文件或文件夹是否存在
// @param {[type]} ctx context.Context [description]
// @param {string} path 文件夹路径
// @return bool true 存在 false 不存在
func (o *sFile) Exists(ctx context.Context, path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// SaveFile 保存信息到文件
// @param {[type]} ctx context.Context [description]
// @param {string} fileName 文件名
// @param {string} content 文件内容
// @return error
func (o *sFile) SaveFile(ctx context.Context, fileName, content string) error {
	// 文件夹路径
	folderPath, err := o.HomeDir(ctx)
	if err != nil {
		return err
	}
	folderPath = folderPath + fileName + ".json"
	f, err := os.OpenFile(folderPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	f.WriteString(content)
	defer f.Close()
	return nil
}

// ReadFiles 读取文件内容
// @param {[type]} ctx context.Context [description]
// @return []fs.FileInfo,error
func (o *sFile) ReadFiles(ctx context.Context) ([]fs.FileInfo, error) {
	// 文件夹路径
	folderPath, err := o.HomeDir(ctx)
	if err != nil {
		return nil, err
	}
	f, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}
	for i, v := range f {
		if v.Name() == ".DS_Store" {
			f = append(f[:i], f[i+1:]...)
		}
	}
	return f, nil
}
