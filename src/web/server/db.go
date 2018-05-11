package server

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func getDB() *gorm.DB {
	if db == nil {
		var err error
		db, err = gorm.Open("mysql", "root:12345678@/test?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic("failed to connect database")
		}
	}
	return db
}
