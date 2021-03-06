package config

import (
	"os"
	"strconv"
)

// Env 应用环境
type Env string

const (
	// Release 生产环境
	Release Env = "release"
	// Debug 开发环境
	Debug Env = "debug"
	// Test 测试环境
	Test Env = "test"
)

// RpcPrefix rpc前缀
const RpcPrefix = "culture/cloud/grpc"

const (
	// VarPath var目录
	VarPath string = "var"
	// LogPath 日志目录
	LogPath = VarPath + "/logs"
	// TempPath 临时目录
	TempPath = VarPath + "/tmp"
)

// AppConfig 主配置
type AppConfig struct {
	AppName   string
	Env       Env
	TokenKey  string
	AppDomain string
	DataBase  DataBase
	Redis     Redis
	Etcd      Etcd
	HttpPort  string
	RpcPort   string
}

// DataBase 数据库配置
type DataBase struct {
	// 数据库类型
	Connection string
	Host       string
	Port       string
	DataBase   string
	UserName   string
	Password   string
	// 数据表前缀
	Prefix string
	// 数据库连接池中最大闲置连接数
	MaxIdleConn int
	// 数据库最大连接数量
	MaxOpenConn int
	// 数据库连接空闲超时时间
	ConnMaxLifetime int
}

// Redis redis配置
type Redis struct {
	Host      string
	Port      string
	Password  string
	DataBase  int
	KeyPrefix string
}

// Etcd 配置
type Etcd struct {
	Addr string
	Port string
}

// Config 配置实例
var Config *AppConfig

// 初始化配置
func initAppConfig() {

	dbConnection := os.Getenv("db_connection")
	maxIdleConnString := os.Getenv("db_max_idle_conn")
	maxOpenConnString := os.Getenv("db_max_open_conn")
	connMaxLifetimeString := os.Getenv("db_conn_max_lifetime")

	if dbConnection == "" {
		dbConnection = "mysql"
	}
	maxIdleConn, err := strconv.Atoi(maxIdleConnString)
	if err != nil {
		maxIdleConn = 1
	}
	maxOpenConn, err := strconv.Atoi(maxOpenConnString)
	if err != nil {
		maxOpenConn = 5
	}
	connMaxLifetime, err := strconv.Atoi(connMaxLifetimeString)
	if err != nil {
		connMaxLifetime = 300
	}

	redisDatabaseString := os.Getenv("redis_database")
	redisDatabase, err := strconv.Atoi(redisDatabaseString)
	if err != nil {
		redisDatabase = 0
	}

	var env Env
	if e := os.Getenv("env"); e == "" {
		env = Debug
	} else {
		env = Env(e)
	}

	Config = &AppConfig{
		AppName:   "culture",
		Env:       env,
		TokenKey:  os.Getenv("token_secret"),
		AppDomain: os.Getenv("app_domain"),
		HttpPort:  os.Getenv("http_port"),
		RpcPort:   os.Getenv("rpc_port"),
		DataBase: DataBase{
			Connection:      dbConnection,
			Host:            os.Getenv("db_host"),
			Port:            os.Getenv("db_port"),
			DataBase:        os.Getenv("db_database"),
			UserName:        os.Getenv("db_user"),
			Password:        os.Getenv("db_password"),
			Prefix:          os.Getenv("db_table_prefix"),
			MaxIdleConn:     maxIdleConn,
			MaxOpenConn:     maxOpenConn,
			ConnMaxLifetime: connMaxLifetime,
		},
		Redis: Redis{
			Host:      os.Getenv("redis_host"),
			Port:      os.Getenv("redis_port"),
			Password:  os.Getenv("redis_password"),
			DataBase:  redisDatabase,
			KeyPrefix: os.Getenv("redis_key_prefix"),
		},
		Etcd: Etcd{
			Addr: os.Getenv("etcd_host"),
			Port: os.Getenv("etcd_port"),
		},
	}
}

func init() {
	initAppConfig()
}
