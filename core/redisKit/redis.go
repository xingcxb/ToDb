package redisKit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
)

var (
	Addr     = "" //redis链接地址
	Username = "" //用户名
	Password = "" //密码
	Port     = "" //端口号
	Db       = 0  //操作数据库
	rdb      *redis.Client
)

func InitDb() {
	var url strings.Builder
	url.WriteString(Addr)
	url.WriteString(":")
	url.WriteString(Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     url.String(), //redis链接地址
		Username: Username,     //设置用户名
		Password: Password,     //设置密码
		DB:       Db,           //设置默认的数据库
		//PoolSize: 10,           //设置连接池大小
	})
}

// Ping redis测试是否联通
func Ping(ctx context.Context) error {
	err := rdb.Ping(ctx).Err()
	//defer rdb.Close()
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}
	return nil
}

// GetDbCount 获取单个库的数量
func GetDbCount(ctx context.Context, dbId int) int {
	ChangeDb(ctx, dbId)
	count, err := rdb.DBSize(ctx).Result()
	//defer rdb.Close()
	if err != nil {
		fmt.Println("error:", err)
		return 0
	}
	return int(count)
}

// ChangeDb 切换数据库
func ChangeDb(ctx context.Context, dbId int) {
	pipe := rdb.Pipeline()
	_ = pipe.Select(ctx, dbId)
	_, _ = pipe.Exec(ctx)
}

//// Info redis所有信息
//type Info struct {
//	Server Server
//	Clients
//	Memory
//	Persistence
//	Stats
//	Replication
//	CPU
//	Cluster
//	Keyspace
//}
//
//// Server 服务信息
//type Server struct {
//	Redis_version     string `yaml:"redis_version"`  //Redis 服务器版本
//	Redis_git_sha1    string `yaml:"redis_git_sha1"` //Git SHA1
//	Redis_git_dirty   string //Git dirty flag
//	Redis_build_id    string
//	Redis_mode        string //运行模式，单机或者集群
//	Os                string //Redis 服务器的宿主操作系统
//	Arch_bits         string //架构（32 或 64 位）
//	Multiplexing_api  string //Redis 所使用的事件处理机制
//	Atomicvar_api     string //原子处理api
//	Gcc_version       string //	编译 Redis 时所使用的 GCC 版本
//	Process_id        string //服务器进程的 PID
//	Run_id            string //Redis 服务器的随机标识符（用于 Sentinel 和集群）
//	Tcp_port          string //TCP/IP 监听端口
//	Uptime_in_seconds string //自 Redis 服务器启动以来，经过的秒数
//	Uptime_in_days    string //自 Redis 服务器启动以来，经过的秒数
//	Hz                string //edis内部调度（进行关闭timeout的客户端，删除过期key等等）频率，程序规定serverCron每秒运行10次。
//	Configured_hz     string
//	Lru_clock         string //以分钟为单位进行自增的时钟，用于 LRU 管理
//	Executable        string //执行文件
//	Config_file       string //配置文件路径
//}
//
//// Clients 已连接客户端信息
//type Clients struct {
//	Connected_clients          string //已连接客户端的数量（不包括通过从属服务器连接的客户端）
//	Client_longest_output_list string //当前连接的客户端当中，最长的输出列表
//	Client_biggest_input_buf   string //当前连接的客户端当中，最大输入缓存
//	Blocked_clients            string //正在等待阻塞命令（BLPOP、BRPOP、BRPOPLPUSH）的客户端的数量
//}
//
//type Memory struct {
//	Used_memory               string //由 Redis 分配器分配的内存总量，以字节（byte）为单位
//	Used_memory_human         string //以人类可读的格式返回 Redis 分配的内存总量(mb)
//	Used_memory_rss           string //从操作系统的角度，返回 Redis 已分配的内存总量（俗称常驻集大小）。这个值和 top 、 ps等命令的输出一致。(byte)
//	Used_memory_rss_human     string //以人类可读的格式，从操作系统的角度，返回 Redis 已分配的内存总量（俗称常驻集大小）。这个值和 top 、 ps等命令的输出一致。(mb)
//	Used_memory_peak          string //redis的内存消耗峰值(byte)
//	Used_memory_peak_human    string //以人类可读的格式返回redis的内存消耗峰值(mb)
//	Used_memory_peak_perc     string //使用内存与峰值内存的百分比(used_memory/ used_memory_peak) *100%
//	Used_memory_overhead      string //Redis为了维护数据集的内部机制所需的内存开销，包括所有客户端输出缓冲区、查询缓冲区、AOF重写缓冲区和主从复制的backlog。
//	Used_memory_startup       string //Redis服务器启动时消耗的内存
//	Used_memory_dataset       string //数据占用的内存(used_memory—used_memory_overhead)
//	Used_memory_dataset_perc  string //数据占用的内存大小百分比,100%*(used_memory_dataset/(used_memory—used_memory_startup))
//	Allocator_allocated       string //分配器分配的内存
//	Allocator_active          string //分配器活跃的内存
//	Allocator_resident        string //分配器常驻的内存
//	Total_system_memory       string //主机内存总量(byte)
//	Total_system_memory_human string //以人类可读的格式，显示整个系统内存(mb)
//	Used_memory_lua           string //Lua引擎存储占用的内存(byte)
//	Used_memory_lua_human     string //以人类可读的格式，显示Lua脚本存储占用的内存(mb)
//	Used_memory_scripts       string
//	Used_memory_scripts_human string
//	Number_of_cached_scripts  string
//	Maxmemory                 string //配置中设置的最大可使用内存值(byte),默认0,不限制
//	Maxmemory_human           string //配置中设置的最大可使用内存值(mb)
//	Maxmemory_policy          string //当达到maxmemory时的淘汰策略
//	Allocator_frag_ratio      string //分配器的碎片率
//	Allocator_frag_bytes      string //分配器的碎片大小
//	Allocator_rss_ratio       string //分配器常驻内存比例
//	Allocator_rss_bytes       string //分配器的常驻内存大小
//	Rss_overhead_ratio        string //常驻内存开销比例
//	Rss_overhead_bytes        string //常驻内存开销大小
//	Mem_fragmentation_ratio   string //碎片率(used_memory_rss / used_memory),正常(1,1.6),大于比例说明内存碎片严重
//	Mem_fragmentation_bytes   string //内存碎片大小
//	Mem_not_counted_for_evict string //被驱逐的内存
//	Mem_replication_backlog   string //redis复制积压缓冲区内存
//	Mem_clients_slaves        string //Redis节点客户端消耗内存
//	Mem_clients_normal        string //Redis所有常规客户端消耗内存
//	Mem_aof_buffer            string //AOF使用内存
//	Mem_allocator             string //内存分配器
//	Active_defrag_running     string //活动碎片整理是否处于活动状态(0没有,1正在运行)
//	Lazyfree_pending_objects  string //0-不存在延迟释放的挂起对象
//}
//
//// Persistence RDB 和 AOF 的相关信息
//type Persistence struct {
//	Loading                      string //服务器是否正在载入持久化文件
//	Rdb_changes_since_last_save  string //离最近一次成功生成rdb文件，写入命令的个数，即有多少个写入命令没有持久化
//	Rdb_bgsave_in_progress       string //服务器是否正在创建rdb文件
//	Rdb_last_save_time           string //离最近一次成功创建rdb文件的时间戳。当前时间戳 - rdb_last_save_time=多少秒未成功生成rdb文件
//	Rdb_last_bgsave_status       string //最近一次rdb持久化是否成功
//	Rdb_last_bgsave_time_sec     string //最近一次成功生成rdb文件耗时秒数
//	Rdb_current_bgsave_time_sec  string //如果服务器正在创建rdb文件，那么这个域记录的就是当前的创建操作已经耗费的秒数
//	Rdb_last_cow_size            string //RDB过程中父进程与子进程相比执行了多少修改(包括读缓冲区，写缓冲区，数据修改等)。
//	Aof_enabled                  string //是否开启了aof
//	Aof_rewrite_in_progress      string //标识aof的rewrite操作是否在进行中
//	Aof_rewrite_scheduled        string //rewrite任务计划，当客户端发送bgrewriteaof指令，如果当前rewrite子进程正在执行，那么将客户端请求的bgrewriteaof变为计划任务，待aof子进程结束后执行rewrite
//	Aof_last_rewrite_time_sec    string //最近一次aof rewrite耗费的时长
//	Aof_current_rewrite_time_sec string //如果rewrite操作正在进行，则记录所使用的时间，单位秒
//	Aof_last_bgrewrite_status    string //上次bgrewriteaof操作的状态
//	Aof_last_write_status        string //上次aof写入状态
//	Aof_last_cow_size            string //AOF过程中父进程与子进程相比执行了多少修改(包括读缓冲区，写缓冲区，数据修改等)。
//}
//
//// Stats 一般统计信息
//type Stats struct {
//	Total_connections_received     string //新创建连接个数,如果新创建连接过多，过度地创建和销毁连接对性能有影响，说明短连接严重或连接池使用有问题，需调研代码的连接设置
//	Total_commands_processed       string //redis处理的命令数
//	Instantaneous_ops_per_sec      string //redis当前的qps，redis内部较实时的每秒执行的命令数
//	Total_net_input_bytes          string //redis网络入口流量字节数
//	Total_net_output_bytes         string //redis网络出口流量字节数
//	Instantaneous_input_kbps       string //redis网络入口kps
//	Instantaneous_output_kbps      string //redis网络出口kps
//	Rejected_connections           string //拒绝的连接个数，redis连接个数达到maxclients限制，拒绝新连接的个数
//	Sync_full                      string //主从完全同步成功次数
//	Sync_partial_ok                string //主从部分同步成功次数
//	Sync_partial_err               string //主从部分同步失败次数
//	Expired_keys                   string //运行以来过期的key的数量
//	Expired_stale_perc             string //过期的比率
//	Expired_time_cap_reached_count string //过期计数
//	Evicted_keys                   string //运行以来剔除(超过了maxmemory后)的key的数量
//	Keyspace_hits                  string //命中次数
//	Keyspace_misses                string //没命中次数
//	Pubsub_channels                string //当前使用中的频道数量
//	Pubsub_patterns                string //当前使用的模式的数量
//	Latest_fork_usec               string //最近一次fork操作阻塞redis进程的耗时数，单位微秒
//	Migrate_cached_sockets         string //是否已经缓存了到该地址的连接
//	Slave_expires_tracked_keys     string //从实例到期key数量
//	Active_defrag_hits             string //主动碎片整理命中次数
//	Active_defrag_misses           string //主动碎片整理未命中次数
//	Active_defrag_key_hits         string //主动碎片整理key命中次数
//	Active_defrag_key_misses       string //主动碎片整理key未命中次数
//}
//
//// Replication 主/从复制信息
//type Replication struct {
//	Role                           string `yaml:"role"`                           //实例的角色，是master or slave
//	Connected_slaves               string `yaml:"connected_slaves"`               //连接的slave实例个数
//	Master_replid                  string `yaml:"master_replid"`                  //主实例启动随机字符串
//	Master_replid2                 string `yaml:"master_replid2"`                 //主实例启动随机字符串2
//	Master_repl_offset             string `yaml:"master_repl_offset"`             //主从同步偏移量,此值如果和上面的offset相同说明主从一致没延迟，与master_replid可被用来标识主实例复制流中的位置。
//	Second_repl_offset             string `yaml:"second_repl_offset"`             //主从同步偏移量2,此值如果和上面的offset相同说明主从一致没延迟
//	Repl_backlog_active            string `yaml:"repl_backlog_active"`            //复制积压缓冲区是否开启
//	Repl_backlog_size              string `yaml:"repl_backlog_size"`              //复制积压缓冲大小
//	Repl_backlog_first_byte_offset string `yaml:"repl_backlog_first_byte_offset"` //复制缓冲区里偏移量的大小
//	Repl_backlog_histlen           string `yaml:"repl_backlog_histlen"`           //此值等于 master_repl_offset - repl_backlog_first_byte_offset,该值不会超过repl_backlog_size的大小
//}
//
////CPU 计算量统计信息
//type CPU struct {
//	Used_cpu_sys           string `yaml:"used_cpu_sys"`           //将所有redis主进程在核心态所占用的CPU时求和累计起来
//	Used_cpu_user          string `yaml:"used_cpu_user"`          //将所有redis主进程在用户态所占用的CPU时求和累计起来
//	Used_cpu_sys_children  string `yaml:"used_cpu_sys_children"`  //将后台进程在核心态所占用的CPU时求和累计起来
//	Used_cpu_user_children string `yaml:"used_cpu_user_children"` //将后台进程在用户态所占用的CPU时求和累计起来
//}
//
////Cluster Redis 集群信息
//type Cluster struct {
//	Cluster_enabled string `yaml:"cluster_enabled"` //实例是否启用集群模式
//}
//
////Keyspace 数据库相关的统计信息
//type Keyspace struct {
//	Db0  string `yaml:"db0"`  //db0
//	Db1  string `yaml:"db1"`  //db1
//	Db2  string `yaml:"db2"`  //db2
//	Db3  string `yaml:"db3"`  //db3
//	Db4  string `yaml:"db4"`  //db4
//	Db5  string `yaml:"db5"`  //db5
//	Db6  string `yaml:"db6"`  //db6
//	Db7  string `yaml:"db7"`  //db7
//	Db8  string `yaml:"db8"`  //db8
//	Db9  string `yaml:"db9"`  //db9
//	Db10 string `yaml:"db10"` //db10
//	Db11 string `yaml:"db11"` //db11
//	Db12 string `yaml:"db12"` //db12
//	Db13 string `yaml:"db13"` //db13
//	Db14 string `yaml:"db14"` //db14
//	Db15 string `yaml:"db15"` //db15
//}

// GetBaseAllInfo  获取redis基础信息
func GetBaseAllInfo(ctx context.Context) map[string]string {
	_info := rdb.Info(ctx).String()
	defer rdb.Close()
	_vs := strings.Split(_info, "\r\n")
	infoMap := make(map[string]string)
	for _, _str := range _vs {
		//_strs[0]:key  _strs[1]:value
		_strs := strings.Split(_str, ":")
		if len(_strs) != 2 || _strs[0] == "info" {
			continue
		}
		infoMap[_strs[0]] = _strs[1]
	}
	return infoMap
}

// GetMainViewInfo 获取主要信息展示信息
func GetMainViewInfo(ctx context.Context) string {
	allInfo := GetBaseAllInfo(ctx)
	if len(allInfo) == 0 {
		return ""
	}
	// 服务器
	server := make(map[string]string)
	//redis版本
	server["version"] = allInfo["redis_version"]
	//redis运行系统
	server["os"] = allInfo["os"]
	//redis进程id
	server["process"] = allInfo["process_id"]

	//内存
	memory := make(map[string]string)
	//已用内存
	memory["usedMemory"] = allInfo["used_memory_human"]
	//内存占用峰值
	memory["usedBigMemory"] = allInfo["used_memory_peak_human"]
	//Lua占用内存
	memory["luaMemory"] = allInfo["used_memory_lua_human"]

	//状态
	start := make(map[string]string)
	// 当前redis连接数
	start["connectCount"] = allInfo["connected_clients"]
	// 历史连接个数
	start["historyCount"] = allInfo["client_recent_max_input_buffer"]
	// 历史执行命令
	start["historyInstructions"] = allInfo["total_commands_processed"]

	//键值列表
	dbkv := make(map[string]string)
	for i := 0; i < 16; i++ {
		var _key strings.Builder
		_key.WriteString("db")
		_key.WriteString(strconv.Itoa(i))
		dbkv[_key.String()] = allInfo[_key.String()]
	}

	all := make(map[string]map[string]string)
	all["server"] = server
	all["memory"] = memory
	all["start"] = start
	all["dbkv"] = dbkv
	strByte, err := json.Marshal(all)
	if err != nil {
		return ""
	}
	return string(strByte)
}

// GetDbData 获取指定库中的数据
func GetDbData(ctx context.Context, cursor uint64) (string, error) {
	keys, cursor, err := rdb.Scan(ctx, cursor, "*", 100).Result()
	if err != nil {
		return "", err
	}
	for _, key := range keys {
		fmt.Println("|==========", key)
		//splitStr := strings.Split(key, ":")
	}
	return "", err
}
