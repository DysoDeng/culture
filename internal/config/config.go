package config

import (
	"os"
	"strconv"
)

// 应用环境
type Env string

const (
	// 生产环境
	Release Env = "release"
	// 开发环境
	Debug Env = "debug"
	// 测试环境
	Test Env = "test"
)

// rpc前缀
const RpcPrefix = "culture/cloud/grpc"

const (
	VarPath  string = "var"
	LogPath         = VarPath + "/logs"
	TempPath        = VarPath + "/tmp"
)

type AppConfig struct {
	AppName   string
	Env       Env
	TokenKey  string
	AppDomain string
	DataBase  DataBase
	Redis     Redis
	Etcd      Etcd
}

// 数据库配置
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

// redis配置
type Redis struct {
	Host      string
	Port      string
	Password  string
	DataBase  int
	KeyPrefix string
}

type Etcd struct {
	Addr string
	Port string
}

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
