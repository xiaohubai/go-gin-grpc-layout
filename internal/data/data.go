package data

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/xiaohubai/go-gin-grpc-layout/pkg/config"
	pgorm "github.com/xiaohubai/go-gin-grpc-layout/pkg/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var data Data

func New() *Data {
	return &data
}

type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewModel(tx ...*pgorm.DB) *pgorm.Model {
	if len(tx) > 0 && tx[0] != nil {
		return pgorm.NewModel(tx[0])
	}

	return pgorm.NewModel(data.db)
}

func NewCrud[T any](tx ...*gorm.DB) *pgorm.Crud[T] {
	return pgorm.NewCrud[T](NewModel(tx...))
}

func Init(c *config.Data) error {
	mysqlConfig := mysql.Config{
		DSN:                       c.MySQL.Source, // DSN data source name
		DefaultStringSize:         191,            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,           // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,           // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,           // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,          // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		//Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true, //禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //禁用表复数形式
			TablePrefix:   "",   //表前缀
		},
	})
	if err != nil {
		return fmt.Errorf("mysql connect failed: %s", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("redis connect ping failed: %s", err)
	}

	data = Data{
		db:  db,
		rdb: rdb,
	}

	return nil
}
