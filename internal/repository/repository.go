package repository

import (
	"github.com/renfy96/yy-layout-v1/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	// TODO: init db
	db, err := gorm.Open(mysql.Open(config.GetString("data.mysql.user")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
