package communication

import (
	"ToDb/internal/structs"
	"ToDb/kit"
	"ToDb/kit/os"
	"ToDb/kit/redis"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
func (s *sRedis) LoadingHistoryInfo(fileName string) (int, string) {
	s.initRedis(fileName)
	err := redisKit.Redis().Ping(context.Background())
	if err != nil {
		return http.StatusBadRequest, err.Error()
	}
	var data []structs.DbTreeInfo
	//redis则直接显示15个库
	for i := 0; i < 16; i++ {
		var dbName strings.Builder
		dbName.WriteString("db")
		dbName.WriteString(strconv.Itoa(i))
		dbName.WriteString("(")
		dbName.WriteString(strconv.Itoa(redisKit.Redis().GetDbCount(context.Background(), i)))
		dbName.WriteString(")")
		data = append(data, structs.DbTreeInfo{
			Label: dbName.String(),
			Title: dbName.String(),
			Key:   strconv.Itoa(i),
		})
	}

	vb, _ := json.Marshal(data)
	return http.StatusOK, string(vb)
}

// 连接redis
func (s *sRedis) initRedis(fileName string) []byte {
	// 获取所有连接文件的路径
	homeDir, _ := os.File().HomeDir(context.Background())
	var filePath strings.Builder
	filePath.WriteString(homeDir)
	filePath.WriteString(fileName)
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

// LoadingDbResource 加载数据库资源消耗信息
func (s *sRedis) LoadingDbResource(ctx context.Context, key string) string {
	s.initRedis(key)
	return redisKit.Redis().GetMainViewInfo(ctx)
}

// GetNodeData 获取节点数据
func (s *sRedis) GetNodeData(ctx context.Context, connType, connName, nodeIdStr string) (string, error) {
	var value string
	if connType == "" ||
		connName == "" {
		return value, errors.New("parameter is missing")
	}
	s.initRedis(connName)
	nodeId, _ := strconv.Atoi(nodeIdStr)
	redisKit.Redis().ChangeDb(ctx, nodeId)
	arr, err := redisKit.Redis().GetDbKeys(ctx, 0)
	if err != nil {
		return "", err
	}
	value = kit.StrKit().PackageTree(arr)
	return value, nil
}

// RedisGetData 通过key获取连接信息
func (s *sRedis) RedisGetData(ctx context.Context, connType, connName, nodeIdStr, key string) (structs.GetValue, error) {
	var getValue structs.GetValue
	if connType == "" ||
		connName == "" {
		return getValue, errors.New("parameter is missing")
	}
	s.initRedis(connName)
	nodeId, _ := strconv.Atoi(nodeIdStr)
	redisKit.Redis().ChangeDb(ctx, nodeId)
	// 获取数据类型
	valueType := redisKit.Redis().GetType(ctx, key)
	valueType = strings.ToLower(valueType)
	fmt.Println("测试数据：", connType, connName, nodeIdStr, key, valueType)
	switch valueType {
	case "string":
		// 获取类型为string的数据
		v := redisKit.Redis().GetStrValue(ctx, key)
		getValue.Type = "string"
		getValue.Key = key
		getValue.Ttl = redisKit.Redis().GetTTL(ctx, key)
		getValue.Value = v
		getValue.Size = len(v)
		command := s.BuildCommand(key, "string", v)
		getValue.CommandStr = command
		return getValue, nil
	case "list":
		// 获取类型为list的数据
		v := redisKit.Redis().GetListValue(ctx, key)
		getValue.Type = "list"
		getValue.Key = key
		getValue.Ttl = redisKit.Redis().GetTTL(ctx, key)
		var listValues []structs.RedisList
		for i, vv := range v {
			listValues = append(listValues, structs.RedisList{
				Id:    i + 1,
				Value: vv,
			})
		}
		getValue.Value = listValues
		command := s.BuildCommand(key, "list", v)
		getValue.CommandStr = command
		return getValue, nil
	case "set":
		// 获取类型为set的数据
		v := redisKit.Redis().GetSetValue(ctx, key)
		getValue.Type = "set"
		getValue.Key = key
		getValue.Ttl = redisKit.Redis().GetTTL(ctx, key)
		var setValues []structs.RedisList
		for i, vv := range v {
			setValues = append(setValues, structs.RedisList{
				Id:    i + 1,
				Value: vv,
			})
		}
		getValue.Value = setValues
		command := s.BuildCommand(key, "set", v)
		getValue.CommandStr = command
		return getValue, nil
	case "hash":
		// 获取类型为hash的数据
		v := redisKit.Redis().GetHashValue(ctx, key)
		getValue.Type = "hash"
		getValue.Key = key
		getValue.Ttl = redisKit.Redis().GetTTL(ctx, key)
		var hashValues []structs.RedisHash
		var i = 1
		for k, vv := range v {
			hashValues = append(hashValues, structs.RedisHash{
				Id:    i,
				Key:   k,
				Value: vv,
			})
			i++
		}
		getValue.Value = hashValues
		command := s.BuildCommand(key, "hash", v)
		getValue.CommandStr = command
		return getValue, nil
	//case "stream":
	default:
		return getValue, errors.New("unknown error")
	}
}

// GetValueType 获取指定key的数据类型
func (s *sRedis) GetValueType(ctx context.Context, connName, nodeIdStr, key string) (string, error) {
	if connName == "" {
		return "", errors.New("parameter is missing")
	}
	s.initRedis(connName)
	s.initRedis(connName)
	nodeId, _ := strconv.Atoi(nodeIdStr)
	redisKit.Redis().ChangeDb(ctx, nodeId)
	// 获取数据类型
	valueType := redisKit.Redis().GetType(ctx, key)
	return valueType, nil
}

// RedisReName 重命名key
func (s *sRedis) RedisReName(ctx context.Context, connType, connName, nodeIdStr, oldKey, newKey string) string {
	if connType == "" ||
		connName == "" {
		return "parameter is missing"
	}
	s.initRedis(connName)
	nodeId, _ := strconv.Atoi(nodeIdStr)
	redisKit.Redis().ChangeDb(ctx, nodeId)
	// 通过键获取值
	v := redisKit.Redis().RenName(ctx, oldKey, newKey)
	if v != nil {
		return v.Error()
	}
	return "success"
}

// RedisUpTtl 更新redis剩余时间
func (s *sRedis) RedisUpTtl(ctx context.Context, connType, connName, nodeIdStr, key string, ttlStr string) string {
	//todo 当ttl=-1时会出现数据直接丢失的情况
	ttl, err := strconv.Atoi(ttlStr)
	if err != nil {
		return "ttl is not number"
	}
	if connType == "" ||
		connName == "" {
		return "parameter is missing"
	}
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
}

// RedisDel 删除redis数据
func (s *sRedis) RedisDel(ctx context.Context, connType, connName, nodeIdStr, key string) string {
	if connType == "" ||
		connName == "" {
		return "parameter is missing"
	}
	s.initRedis(connName)
	nodeId, _ := strconv.Atoi(nodeIdStr)
	redisKit.Redis().ChangeDb(ctx, nodeId)
	// 通过键获取值
	v := redisKit.Redis().Del(ctx, key)
	if v == 0 {
		return "del error"
	}
	return "success"
}

// RedisUpdateStringValue 更新redis数据
func (s *sRedis) RedisUpdateStringValue(ctx context.Context, connType, connName, nodeIdStr, key, value, ttlStr string) error {
	if connType == "" ||
		connName == "" {
		return errors.New("parameter is missing")
	}
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
}

// BuildCommand 构建命令
func (s *sRedis) BuildCommand(key, keyType string, value interface{}) string {
	lowerCaseKeyType := strings.ToLower(keyType)
	var command strings.Builder
	switch lowerCaseKeyType {
	case "string":
		// 构建set命令
		// SET "1:2:34" "你好啊😂"
		command.WriteString("SET ")
		command.WriteString("\"")
		command.WriteString(key)
		command.WriteString("\" \"")
		command.WriteString(value.(string))
		command.WriteString("\"")
	case "hash":
		// 构建hash命令
		//HMSET "1:2:hash" "New field" "New value" "123" "321"
		command.WriteString("HMSET ")
		command.WriteString("\"")
		command.WriteString(key)
		arr := value.(map[string]string)
		for k, v := range arr {
			command.WriteString("\" \"")
			command.WriteString(k)
			command.WriteString("\" \"")
			command.WriteString(v)
		}
		command.WriteString("\"")
	case "list":
		// 构建list命令
		// RPUSH "1:2:list" "New member" "12312213" "1231" "测试"
		// return "RPUSH " + key + " " + value
		command.WriteString("RPUSH ")
		command.WriteString("\"")
		command.WriteString(key)
		command.WriteString("\"")
		arr := value.([]string)
		for _, v := range arr {
			command.WriteString(" \"")
			command.WriteString(v)
			command.WriteString("\"")
		}
	case "set":
		// 构建set命令
		// SADD "1:2:set" "New member" "sdfsdf"
		// return "SADD " + key + " " + value
		command.WriteString("SADD ")
		command.WriteString("\"")
		command.WriteString(key)
		command.WriteString("\"")
		arr := value.([]string)
		for _, v := range arr {
			command.WriteString(" \"")
			command.WriteString(v)
			command.WriteString("\"")
		}
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
