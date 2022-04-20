package communication

import (
	"ToDb/core/redisKit"
	"ToDb/lib"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

// 连接信息
type connectionType struct {
	Type     string `tag:"类型" json:"type"`      //类型
	Alias    string `tag:"连接名" json:"alias"`    //别名
	HostURL  string `tag:"连接地址" json:"hostURL"` //连接地址
	Port     string `tag:"端口号" json:"port"`     //端口号
	Username string `tag:"用户名" json:"username"` //用户名
	Password string `tag:"密码" json:"password"`  //密码
}

// 检查参数
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
		Type:     parameter["type"].String(),
		Alias:    parameter["alias"].String(),
		HostURL:  parameter["hostURL"].String(),
		Port:     parameter["port"].String(),
		Username: parameter["username"].String(),
		Password: parameter["password"].String(),
	}
	if parameter["savePassword"].Bool() {
		var dirBuild strings.Builder
		dirBuild.WriteString(lib.GetProgramSafePath())
		dirBuild.WriteString(info.Alias)
		dirBuild.WriteString(".json")
		filename := dirBuild.String()
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			newFile, err := os.Create(filename)
			if err != nil {
				return code, message
			}
			defer newFile.Close()
			f, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		}
		_v, _ := json.MarshalIndent(info, "", "    ")
		f.WriteString(string(_v))
		defer f.Close()
	}
	return code, message
}

type BaseConnInfo struct {
	Title        string `json:"title"`        //别名
	Key          string `json:"key"`          //适配tree
	ConnType     string `json:"connType"`     //类型
	IconPath     string `json:"iconPath"`     //图标路径
	ConnFileAddr string `json:"ConnFileAddr"` //连接信息文件存放地址
	//Children     string `json:"children"`
}

// LoadingBaseHistoryInfo 加载已经存储的连接别名
func LoadingBaseHistoryInfo() string {
	// 获取所有连接文件的路径
	allFilesPath := lib.GetProgramSafePath()
	files, _ := ioutil.ReadDir(allFilesPath)
	datas := make([]BaseConnInfo, 0, 1)
	for _, f := range files {
		fileName := f.Name()
		var filePath strings.Builder
		filePath.WriteString(allFilesPath)
		filePath.WriteString(fileName)
		valueByte, _ := ioutil.ReadFile(filePath.String())
		t := gjson.Get(string(valueByte), "type").String()
		alias := gjson.Get(string(valueByte), "alias").String()
		var ipt strings.Builder
		ipt.WriteString("leftNavigation/")
		ipt.WriteString(t)
		ipt.WriteString(".png")
		var key strings.Builder
		key.WriteString(t)
		key.WriteString(",")
		key.WriteString(alias)
		bci := BaseConnInfo{
			Title:        alias,
			Key:          key.String(),
			ConnType:     t,
			IconPath:     ipt.String(),
			ConnFileAddr: filePath.String(),
		}
		datas = append(datas, bci)
	}
	jb, _ := json.Marshal(datas)
	return string(jb)
}

type Children struct {
	Title string `json:"title"` //别名
	Key   string `json:"key"`   //key
	//Children []*Children `json:"children"` //子集
}

// LoadingHistoryInfo 加载已经存储的连接信息
func LoadingHistoryInfo(key string) (int, string) {
	valueByte := initRedis(key)
	err := redisKit.Ping(context.Background())
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	t := gjson.Get(string(valueByte), "type").String()
	var data []Children
	switch t {
	case "redis":
		//如果是redis则直接显示15个库
		for i := 0; i < 16; i++ {
			var dbName strings.Builder
			dbName.WriteString("db")
			dbName.WriteString(strconv.Itoa(i))
			dbName.WriteString("(")
			dbName.WriteString(strconv.Itoa(redisKit.GetDbCount(context.Background(), i)))
			dbName.WriteString(")")
			data = append(data, Children{
				Title: dbName.String(),
				Key:   strconv.Itoa(i),
			})
		}
	default:
	}
	vb, _ := json.Marshal(data)
	return http.StatusOK, string(vb)
}

// 连接redis
func initRedis(key string) []byte {
	// 获取所有连接文件的路径
	allFilesPath := lib.GetProgramSafePath()
	var filePath strings.Builder
	filePath.WriteString(allFilesPath)
	filePath.WriteString(key)
	filePath.WriteString(".json")
	valueByte, _ := ioutil.ReadFile(filePath.String())

	//优先初始化redis链接
	redisKit.Port = gjson.Get(string(valueByte), "port").String()
	redisKit.Username = gjson.Get(string(valueByte), "username").String()
	redisKit.Password = gjson.Get(string(valueByte), "password").String()
	redisKit.Addr = gjson.Get(string(valueByte), "hostURL").String()
	redisKit.InitDb()
	return valueByte
}

// LoadingDbResource 加载数据库资源消耗
func LoadingDbResource(key string) string {
	initRedis(key)
	return redisKit.GetMainViewInfo(context.Background())
}

// GetNodeData 获取节点数据
func GetNodeData(connType, connName, nodeIdStr string) (string, error) {
	var value strings.Builder
	if connType == "" ||
		connName == "" {
		return value.String(), errors.New("parameter is missing")
	}
	ctx := context.Background()
	switch connType {
	case "redis":
		initRedis(connName)
		nodeId, _ := strconv.Atoi(nodeIdStr)
		redisKit.ChangeDb(ctx, nodeId)
		arr, err := redisKit.GetDbKeys(ctx, 0)
		if err != nil {
			return "", err
		}
		value := lib.PackageTree(arr)
		return value, nil
	default:
		return "", errors.New("unknown error")
	}
}

// RedisGetData 通过key获取连接信息
func RedisGetData(connType, connName, nodeIdStr, key string) (string, error) {
	var value strings.Builder
	if connType == "" ||
		connName == "" {
		return value.String(), errors.New("parameter is missing")
	}
	ctx := context.Background()
	switch connType {
	case "redis":
		initRedis(connName)
		nodeId, _ := strconv.Atoi(nodeIdStr)
		redisKit.ChangeDb(ctx, nodeId)
		// 通过键获取值
		v := redisKit.GetKeyInfo(ctx, key)
		return v, nil
	default:
		return "", errors.New("unknown error")
	}
}

// 重命名key
func RedisReName(connType, connName, nodeIdStr, oldKey, newKey string) string {
	if connType == "" ||
		connName == "" {
		return "parameter is missing"
	}
	ctx := context.Background()
	switch connType {
	case "redis":
		initRedis(connName)
		nodeId, _ := strconv.Atoi(nodeIdStr)
		redisKit.ChangeDb(ctx, nodeId)
		// 通过键获取值
		v := redisKit.RenName(ctx, oldKey, newKey)
		if v != nil {
			return v.Error()
		}
		return "success"
	default:
		return "unknown error"
	}
}

// RedisUpTtl 更新redis剩余时间
func RedisUpTtl(connType, connName, nodeIdStr, key string, ttlStr string) string {
	//todo 当ttl=-1时会出现数据直接丢失的情况
	ttl, err := strconv.Atoi(ttlStr)
	if err != nil {
		return "ttl is not number"
	}
	if connType == "" ||
		connName == "" {
		return "parameter is missing"
	}
	ctx := context.Background()
	switch connType {
	case "redis":
		initRedis(connName)
		nodeId, _ := strconv.Atoi(nodeIdStr)
		redisKit.ChangeDb(ctx, nodeId)
		// 通过键获取值
		var v error
		if ttl == -1 {
			// 表示需要永久存储
			v = redisKit.UpPermanent(ctx, key)
		} else {
			v = redisKit.UpTtl(ctx, key, ttl)
		}
		if v != nil {
			return v.Error()
		}
		return "success"
	default:
		return "unknown error"
	}
}

// RedisDel 删除redis数据
func RedisDel(connType, connName, nodeIdStr, key string) string {
	if connType == "" ||
		connName == "" {
		return "parameter is missing"
	}
	ctx := context.Background()
	switch connType {
	case "redis":
		initRedis(connName)
		nodeId, _ := strconv.Atoi(nodeIdStr)
		redisKit.ChangeDb(ctx, nodeId)
		// 通过键获取值
		v := redisKit.Del(ctx, key)
		if v == 0 {
			return "del error"
		}
		return "success"
	default:
		return "unknown error"
	}
}

// BuildCommand 构建命令
func BuildCommand(connType, connName, nodeIdStr, key string) string {

	return ""
}
