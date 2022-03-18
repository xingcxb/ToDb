package communication

import (
	"ToDb/core/redisKit"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"net/http"
	"os"
	"reflect"
	goruntime "runtime"
	"strings"
)

// 连接信息
type connectionType struct {
	Alias    string `tag:"连接名" json:"alias"`    //别名
	HostURL  string `tag:"连接地址" json:"hostURL"` //连接地址
	Port     string `tag:"端口号" json:"port"`     //端口号
	Username string `tag:"用户名" json:"username"` //用户名
	Password string `tag:"密码" json:"password"`  //密码
}

func checkParameter(parameter map[string]gjson.Result) (int, string, bool) {
	//返回连接信息
	message := "连接成功"
	//状态码
	code := http.StatusOK
	//用于标记是否要继续匹配
	fail := false
	for k, v := range parameter {
		if k == "savePassword" || k == "username" {
			continue
		}
		if v.String() == "" {
			code = http.StatusBadRequest
			fieldV, _ := reflect.TypeOf(connectionType{}).FieldByName(k)
			tag := fieldV.Tag.Get("tag")
			message = tag + "不能为空"
			fail = true
			break
		}
	}
	return code, message, fail
}

// RedisPing redis测试连接
func RedisPing(connectionInfo string) (int, string) {
	//返回连接信息
	message := "连接成功"
	//状态码
	code := http.StatusOK
	//用于标记是否要继续匹配
	fail := false
	parameter := gjson.Parse(connectionInfo).Map()
	code, message, fail = checkParameter(parameter)

	info := connectionType{
		Alias:    parameter["alias"].String(),
		HostURL:  parameter["hostURL"].String(),
		Port:     parameter["port"].String(),
		Username: parameter["username"].String(),
		Password: parameter["password"].String(),
	}

	if !fail {
		redisKit.Addr = info.HostURL
		redisKit.Port = info.Port
		redisKit.Username = info.Username
		redisKit.Password = info.Password
		redisKit.InitDb()
		err := redisKit.Ping(context.Background())
		if err != nil {
			code = http.StatusBadRequest
			message = err.Error()
		} else {
			code = http.StatusOK
			message = "连接成功"
		}
	}
	return code, message
}

// Ok 确定按钮
func Ok(ctx context.Context, connectionInfo string) (int, string) {
	//返回连接信息
	message := "连接成功"
	//状态码
	code := http.StatusOK
	parameter := gjson.Parse(connectionInfo).Map()
	code, message, _ = checkParameter(parameter)
	info := connectionType{
		Alias:    parameter["alias"].String(),
		HostURL:  parameter["hostURL"].String(),
		Port:     parameter["port"].String(),
		Username: parameter["username"].String(),
		Password: parameter["password"].String(),
	}
	if parameter["savePassword"].Bool() {
		var dirBuild strings.Builder
		dir, _ := os.Getwd()
		fmt.Println("=========================", dir)
		dirBuild.WriteString(dir)
		if goruntime.GOOS == "windows" {
			//windows下存放配置文件路径
			dirBuild.WriteString("\\safe\\")
		} else if goruntime.GOOS == "darwin" {
			//macOS下存放配置文件路径
			dirBuild.WriteString("/safe/")
		}
		dirBuild.WriteString(info.Alias)
		dirBuild.WriteString(".json")
		filename := dirBuild.String()
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err, "================================================")
			newFile, err := os.Create(filename)
			if err != nil {
				fmt.Println(err, "==============2==================================")
				return code, message
			}
			defer newFile.Close()
			f, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		}
		_v, _ := json.MarshalIndent(info, "", "    ")
		_, err = f.WriteString(string(_v))
		if err != nil {
			runtime.LogError(ctx, err.Error())
		}
		defer f.Close()
	}
	return code, message
}
