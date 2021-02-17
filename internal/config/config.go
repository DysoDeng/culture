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
	Connection string
	Host       string
	Port       string
	DataBase   string
	UserName   string
	Password   string
	Prefix     string
	// 数据库连接池中最大闲置连接数
	MaxIdleConn int
	// 数据库最大连接数量
	MaxOpenConn int
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

	maxIdleConnString := os.Getenv("mysql_max_idle_conn")
	maxOpenConnString := os.Getenv("mysql_max_open_conn")

	maxIdleConn, err := strconv.Atoi(maxIdleConnString)
	if err != nil {
		maxIdleConn = 1
	}
	maxOpenConn, err := strconv.Atoi(maxOpenConnString)
	if err != nil {
		maxOpenConn = 5
	}

	redisDatabaseString := os.Getenv("redis_database")
	redisDatabase, err := strconv.Atoi(redisDatabaseString)
	if err != nil {
		redisDatabase = 0
	}

	Config = &AppConfig{
		AppName:   "culture",
		Env:       Env(os.Getenv("env")),
		TokenKey:  os.Getenv("token_secret"),
		AppDomain: os.Getenv("app_domain"),
		DataBase: DataBase{
			Connection:  "mysql",
			Host:        os.Getenv("mysql_host"),
			Port:        os.Getenv("mysql_port"),
			DataBase:    os.Getenv("mysql_database"),
			UserName:    os.Getenv("mysql_user"),
			Password:    os.Getenv("mysql_password"),
			Prefix:      os.Getenv("mysql_table_prefix"),
			MaxIdleConn: maxIdleConn,
			MaxOpenConn: maxOpenConn,
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
