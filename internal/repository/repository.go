package repository

import (
	"github.com/renfy96/yy-layout-v1/internal/repository/query"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func NewDb(conf *viper.Viper) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func Use(db *gorm.DB) query.Query {
	query.Use(db)
	return *query.Q
}
