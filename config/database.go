package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitialDB() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/go_ktb_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
