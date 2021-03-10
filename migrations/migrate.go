package main

import (
	"culture/cloud/base/internal/model"
	"culture/cloud/base/internal/model/message"
	"culture/cloud/base/internal/support/db"
	"flag"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
)

// 定义数据库迁移
var migration = []*gormigrate.Migration{
	{
		ID: "202103101456",
		Migrate: func(g *gorm.DB) error {
			err := g.AutoMigrate(&message.SmsConfig{}, &message.SmsTemplate{})
			if err != nil {
				return err
			}

			g.Exec("ALTER TABLE sms_config COMMENT=\"短信提供商配置表\"")
			g.Exec("ALTER TABLE sms_template COMMENT=\"短信模版配置表\"")
			return nil
		},
		Rollback: func(g *gorm.DB) error {
			return g.Migrator().DropTable(&message.SmsConfig{})
		},
	},
	{
		ID: "202103101459",
		Migrate: func(g *gorm.DB) error {
			return g.AutoMigrate(&model.User{})
		},
		Rollback: func(g *gorm.DB) error {
			return g.Migrator().DropTable(&model.User{})
		},
	},
}

func main() {

	dbSql := db.DB()

	var migrate = flag.Bool("m", false, "执行迁移 -m")
	var rollback = flag.Bool("r", false, "执行迁移回滚 -r versionID")
	var rollbackLast = flag.Bool("rl", false, "执行最后一次迁移回滚 -rl")

	flag.Parse()

	if *migrate {
		m := gormigrate.New(dbSql, gormigrate.DefaultOptions, migration)

		if err := m.Migrate(); err != nil {
			log.Fatalf("Could not migrate: %v", err)
		}
		log.Printf("Migration did run successfully")
	}

	if *rollback {
		arg := flag.Args()
		m := gormigrate.New(dbSql, gormigrate.DefaultOptions, migration)
		if len(arg) > 0 {
			if err := m.RollbackTo(arg[0]); err != nil {
				log.Fatalf("Could not rollback: %v", err)
			}
			log.Printf("Rollback to %s migrate successfully", arg)
		} else {
			log.Fatalf("请指定回滚版本号")
		}
	}

	if *rollbackLast {
		m := gormigrate.New(dbSql, gormigrate.DefaultOptions, migration)

		if err := m.RollbackLast(); err != nil {
			log.Fatalf("Could not rollback: %v", err)
		}
		log.Printf("Rollback to last migrate successfully")
	}
}
