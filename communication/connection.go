package communication

import (
	"ToDb/core/redisKit"
	"context"
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"net/http"
	"os"
)

// 连接信息
type connectionType struct {
	alias    string //别名
	hostURL  string //连接地址
	port     string //端口号
	username string //用户名
	password string //密码
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
			message = "参数" + k + "不存在值"
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
		alias:    parameter["alias"].String(),
		hostURL:  parameter["hostURL"].String(),
		port:     parameter["port"].String(),
		username: parameter["username"].String(),
		password: parameter["password"].String(),
	}

	if !fail {
		redisKit.Addr = info.hostURL
		redisKit.Port = info.port
		redisKit.Username = info.username
		redisKit.Password = info.password
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
func Ok(connectionInfo string) (int, string) {
	//返回连接信息
	message := "连接成功"
	//状态码
	code := http.StatusOK
	parameter := gjson.Parse(connectionInfo).Map()
	code, message, _ = checkParameter(parameter)
	info := connectionType{
		alias:    parameter["alias"].String(),
		hostURL:  parameter["hostURL"].String(),
		port:     parameter["port"].String(),
		username: parameter["username"].String(),
		password: parameter["password"].String(),
	}
	if parameter["savePassword"].Bool() {
		filename := info.alias + ".json"
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			//打开文件错误，创建文件
			newFile, err := os.Create(filename)
			if err != nil {
				return code, message
			}
			defer newFile.Close()
		} else {
			//此处上面已经绝对保证有文件
			f, err = os.Open(info.alias)
		}
		_v, _ := json.Marshal(info)
		_, err = f.WriteString(string(_v))
		if err != nil {
			runtime.LogError(context.Background(), err.Error())
		}
	}
	return code, message
}
