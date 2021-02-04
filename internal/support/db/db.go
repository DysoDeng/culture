package db

import (
	"culture/internal/config"
	"database/sql/driver"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var DB *gorm.DB

// 初始化数据库连接
func initDB() {
	var err error
	var dsn string

	defer func() {
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

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true, // 禁用默认事务
		PrepareStmt:                              true, // 预编译sql
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用创建外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁止表名复数
		},
	})
	if err != nil {
		panic("failed to connect database " + err.Error())
	}

	sqlDB, _ := DB.DB()

	// 连接池
	sqlDB.SetMaxIdleConns(config.Config.DataBase.MaxIdleConn) // 连接池最大连接数
	sqlDB.SetMaxOpenConns(config.Config.DataBase.MaxOpenConn) // 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于该值，超过的连接会被连接池关闭。
	sqlDB.SetConnMaxLifetime(time.Hour)
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

func init() {
	initDB()
}
