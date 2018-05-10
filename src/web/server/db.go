package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func getDB() *gorm.DB {
	if db == nil {
		var err error
		db, err = gorm.Open("mysql", "root:12345@/demo?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			panic("failed to connect database")
		}
	}

	return db
}
