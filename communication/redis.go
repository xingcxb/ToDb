package communication

import (
	"ToDb/internal/structs"
	"ToDb/kit"
	"ToDb/kit/os"
	"ToDb/kit/redis"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

var (
	insRedis = sRedis{}
)

type sRedis struct {
	ctx context.Context
}

func Redis() *sRedis {
	return &insRedis
}

// 检查参数
func (s *sRedis) checkParameter(parameter map[string]gjson.Result) (int, string, bool) {
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
			fieldV, _ := reflect.TypeOf(structs.ConnectionType{}).FieldByName(k)
			tag := fieldV.Tag.Get("tag")
			message = tag + "不能为空"
			fail = true
			break
		}
	}
	return code, message, fail
}

// RedisPing redis测试连接
func (s *sRedis) RedisPing(connectionInfo string) (int, string) {
	//返回连接信息
	message := "连接成功"
	//状态码
	code := http.StatusOK
	//用于标记是否要继续匹配
	fail := false
	parameter := gjson.Parse(connectionInfo).Map()
	code, message, fail = s.checkParameter(parameter)

	info := structs.ConnectionType{
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
		err := redisKit.Redis().Ping(context.Background())
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
func (s *sRedis) Ok(ctx context.Context, connectionInfo string) (int, string) {
	//返回连接信息
	message := "连接成功"
	//状态码
	code := http.StatusOK
	parameter := gjson.Parse(connectionInfo).Map()
	code, message, _ = s.checkParameter(parameter)
	info := structs.ConnectionType{
		Type:     parameter["connType"].String(),
		Alias:    parameter["alias"].String(),
		HostURL:  parameter["hostURL"].String(),
		Port:     parameter["port"].String(),
		Username: parameter["username"].String(),
		Password: parameter["password"].String(),
	}
	if parameter["savePassword"].Bool() {
		// 创建文件或文件夹
		err := os.File().CreateFile(ctx, info.Alias)
		if err != nil {
			return code, message
		}
		value, _ := json.MarshalIndent(info, "", "    ")
		err = os.File().SaveFile(ctx, info.Alias, string(value))
		if err != nil {
			return 0, "保存失败"
		}
	}
	return code, message
}

// LoadingBaseHistoryInfo 加载已经存储的连接别名
func (s *sRedis) LoadingBaseHistoryInfo(ctx context.Context) string {
	// 读取所有的文件
	files, _ := os.File().ReadFiles(ctx)
	datas := make([]structs.BaseTreeInfo, 0, 1)
	// 获取程序设定的基础路径
	homeDir, _ := os.File().HomeDir(ctx)
	// 循环组装数据，用于前端tree的显示
	for _, f := range files {
		fileName := f.Name()
		var filePath strings.Builder
		filePath.WriteString(homeDir)
		filePath.WriteString(fileName)
		// 读取文件内容
		valueByte, _ := ioutil.ReadFile(filePath.String())
		// 提取类型
		t := gjson.Get(string(valueByte), "type").String()
		// 提取别名
		alias := gjson.Get(string(valueByte), "alias").String()
		bci := structs.BaseTreeInfo{
			Title:        alias,
			Label:        alias,
			ConnType:     t,
			IconPath:     "",
			ConnFileAddr: filePath.String(),
		}
		datas = append(datas, bci)
	}
	jb, _ := json.Marshal(datas)
	return string(jb)
}

// LoadingHistoryInfo 加载已经存储的连接信息
func (s *sRedis) LoadingHistoryInfo(key string) (int, string) {
	valueByte := s.initRedis(key)
	err := redisKit.Redis().Ping(context.Background())
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}

	t := gjson.Get(string(valueByte), "type").String()
	var data []structs.DbTreeInfo
	switch t {
	case "redis":
		//如果是redis则直接显示15个库
		for i := 0; i < 16; i++ {
			var dbName strings.Builder
			dbName.WriteString("db")
			dbName.WriteString(strconv.Itoa(i))
			dbName.WriteString("(")
			dbName.WriteString(strconv.Itoa(redisKit.Redis().GetDbCount(context.Background(), i)))
			dbName.WriteString(")")
			data = append(data, structs.DbTreeInfo{
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
func (s *sRedis) initRedis(key string) []byte {
	// 获取所有连接文件的路径
	homeDir, _ := os.File().HomeDir(context.Background())
	var filePath strings.Builder
	filePath.WriteString(homeDir)
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
func (s *sRedis) LoadingDbResource(key string) string {
	s.initRedis(key)
	return redisKit.Redis().GetMainViewInfo(context.Background())
}

// GetNodeData 获取节点数据
func (s *sRedis) GetNodeData(connType, connName, nodeIdStr string) (string, error) {
	var value strings.Builder
	if connType == "" ||
		connName == "" {
		return value.String(), errors.New("parameter is missing")
	}
	ctx := context.Background()
	switch connType {
	case "redis":
		s.initRedis(connName)
		nodeId, _ := strconv.Atoi(nodeIdStr)
		redisKit.Redis().ChangeDb(ctx, nodeId)
		arr, err := redisKit.Redis().GetDbKeys(ctx, 0)
		if err != nil {
			return "", err
		}
		value := kit.StrKit().PackageTree(arr)
		return value, nil
	default:
		return "", errors.New("unknown error")
	}
}

// RedisGetData 通过key获取连接信息
func (s *sRedis) RedisGetData(connType, connName, nodeIdStr, key string) (structs.GetValue, error) {
	// var value strings.Builder
	var getValue structs.GetValue
	if connType == "" ||
		connName == "" {
		return getValue, errors.New("parameter is missing")
	}
	ctx := context.Background()
	switch connType {
	case "redis":
		s.initRedis(connName)
		nodeId, _ := strconv.Atoi(nodeIdStr)
		redisKit.Redis().ChangeDb(ctx, nodeId)
		// 获取数据类型
		valueType := redisKit.Redis().GetType(ctx, key)
		valueType = strings.ToLower(valueType)
		switch valueType {
		case "string":
			// 通过键获取值
			v := redisKit.Redis().GetValue(ctx, key)
			command := s.BuildCommand(key, "string", v)
			getValue.Type = "string"
			getValue.Key = key
			getValue.Ttl = redisKit.Redis().GetTTL(ctx, key)
			getValue.Value = v
			getValue.Size = len(v)
			getValue.CommandStr = command
			return getValue, nil
		default:
		}
		return getValue, errors.New("unknown error")
	default:
		return getValue, errors.New("unknown error")
	}
}

// RedisReName 重命名key
func (s *sRedis) RedisReName(connType, connName, nodeIdStr, oldKey, newKey string) string {
	if connType == "" ||
		connName == "" {
		return "parameter is missing"
	}
	ctx := context.Background()
	switch connType {
	case "redis":
		s.initRedis(connName)
		nodeId, _ := strconv.Atoi(nodeIdStr)
		redisKit.Redis().ChangeDb(ctx, nodeId)
		// 通过键获取值
		v := redisKit.Redis().RenName(ctx, oldKey, newKey)
		if v != nil {
			return v.Error()
		}
		return "success"
	default:
		return "unknown error"
	}
}

// RedisUpTtl 更新redis剩余时间
func (s *sRedis) RedisUpTtl(connType, connName, nodeIdStr, key string, ttlStr string) string {
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
		s.initRedis(connName)
		nodeId, _ := strconv.Atoi(nodeIdStr)
		redisKit.Redis().ChangeDb(ctx, nodeId)
		// 通过键获取值
		var v error
		if ttl == -1 {
			// 表示需要永久存储
			v = redisKit.Redis().UpPermanent(ctx, key)
		} else {
			v = redisKit.Redis().UpTtl(ctx, key, ttl)
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
func (s *sRedis) RedisDel(connType, connName, nodeIdStr, key string) string {
	if connType == "" ||
		connName == "" {
		return "parameter is missing"
	}
	ctx := context.Background()
	switch connType {
	case "redis":
		s.initRedis(connName)
		nodeId, _ := strconv.Atoi(nodeIdStr)
		redisKit.Redis().ChangeDb(ctx, nodeId)
		// 通过键获取值
		v := redisKit.Redis().Del(ctx, key)
		if v == 0 {
			return "del error"
		}
		return "success"
	default:
		return "unknown error"
	}
}

// RedisUpdateStringValue 更新redis数据
func (s *sRedis) RedisUpdateStringValue(connType, connName, nodeIdStr, key, value, ttlStr string) error {
	if connType == "" ||
		connName == "" {
		return errors.New("parameter is missing")
	}
	ctx := context.Background()
	switch connType {
	case "redis":
		s.initRedis(connName)
		nodeId, _ := strconv.Atoi(nodeIdStr)
		redisKit.Redis().ChangeDb(ctx, nodeId)
		// 通过键获取值
		ttl, _ := strconv.Atoi(ttlStr)
		err := redisKit.Redis().AddData(ctx, key, value, ttl)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("unknown error")
	}
}

// BuildCommand 构建命令
func (s *sRedis) BuildCommand(key, keyType, value string) string {
	lowerCaseKeyType := strings.ToLower(keyType)
	var command strings.Builder
	switch lowerCaseKeyType {
	case "string":
		// 构建set命令
		// SET "1:2:34" "你好啊😂"
		command.WriteString("SET ")
		command.WriteString("\"")
		command.WriteString(key)
		command.WriteString("\"")
		command.WriteString(" ")
		command.WriteString("\"")
		command.WriteString(value)
		command.WriteString("\"")
	case "hash":
		// 构建hash命令
		// HMSET "1:2:hash" "New field" "New value" "123" "321"
		// return "HMSET " + key + " " + value
		command.WriteString("HMSET ")
	case "list":
		// 构建list命令
		// RPUSH "1:2:list" "New member" "12312213"
		// return "RPUSH " + key + " " + value
		command.WriteString("RPUSH ")
	case "set":
		// 构建set命令
		// SADD "1:2:set" "New member" "sdfsdf"
		// return "SADD " + key + " " + value
		command.WriteString("SADD ")
	case "zset":
		// 构建zset命令
		// XADD "1:2:stream" 1650445322163-0  "New key" "New value"
		// XADD "1:2:stream" 21312312312312-0  "New key" "New value"
		// return "ZADD " + key + " " + value
		command.WriteString("ZADD ")
	case "stream":
		// 构建stream命令
		// ZADD "1:2:zset" 12 "321" 0 "New member"
		// return "XADD " + key + " " + value
		command.WriteString("XADD ")
	default:
		return ""
	}
	return command.String()
}
