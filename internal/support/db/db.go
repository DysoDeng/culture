package db

import (
	"culture/internal/config"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/goava/di"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

// container DB服务容器
var container *di.Container

var mutex sync.Mutex

func init() {
	if config.Config.Env != config.Release {
		di.SetTracer(&di.StdTracer{})
	}
	var err error
	container, err = di.New(
		di.Provide(initDB),
	)
	if err != nil {
		panic(err)
	}
}

// 初始化数据库连接
func initDB() *gorm.DB {
	var err error
	var dsn string
	var DB *gorm.DB

	mutex.Lock()
	defer func() {
		mutex.Unlock()
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.Config.DataBase.UserName,
		config.Config.DataBase.Password,
		config.Config.DataBase.Host,
		config.Config.DataBase.Port,
		config.Config.DataBase.DataBase,
	) + "?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"

	// db日志
	logFilename := config.LogPath + "/db.log"
	dbLogFile, _ := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	dbLogger := logger.New(
		log.New(io.MultiWriter(os.Stdout, dbLogFile), "", log.LstdFlags),
		logger.Config{
			SlowThreshold: 200 * time.Millisecond, // 慢查询时间
			LogLevel:      logger.Warn,
			Colorful:      false,
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true, // 禁用默认事务
		PrepareStmt:                              true, // 预编译sql
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁止表名复数
		},
		Logger: dbLogger, // db日志
	})
	if err != nil {
		panic("failed to connect database " + err.Error())
	}

	sqlDB, _ := DB.DB()

	// 连接池
	sqlDB.SetMaxIdleConns(config.Config.DataBase.MaxIdleConn)                                     // 连接池最大连接数
	sqlDB.SetMaxOpenConns(config.Config.DataBase.MaxOpenConn)                                     // 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于该值，超过的连接会被连接池关闭。
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.Config.DataBase.ConnMaxLifetime)) // 连接空闲超时

	sqlDB.Stats()

	return DB
}

// 获取数据库连接
func DB() *gorm.DB {
	var orm *gorm.DB
	_ = container.Resolve(&orm)
	sqlDB, _ := orm.DB()
	data, _ := json.Marshal(sqlDB.Stats()) //获得当前的SQL配置情况
	fmt.Println(string(data))
	return orm
}

// get full table name
func FullTableName(tableName string) string {
	return config.Config.DataBase.Prefix + tableName
}

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	zero, _ := time.Parse("2006-01-02 15:04:05", "0001-01-01 00:00:00")
	zeroTime := JSONTime{Time: zero}
	if t == zeroTime {
		return []byte(fmt.Sprintf("\"%s\"", "")), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueOf time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// JSONDate format json date field by myself
type JSONDate struct {
	time.Time
}

func (t JSONDate) MarshalJSON() ([]byte, error) {
	zero, _ := time.Parse("2006-01-02", "0001-01-01")
	zeroTime := JSONDate{Time: zero}
	if t == zeroTime {
		return []byte(fmt.Sprintf("\"%s\"", "")), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueOf time.Time
func (t *JSONDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
