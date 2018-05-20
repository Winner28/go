package server

import (
	"web/server/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type database struct {
	DB *gorm.DB
}

func getDB() *database {
	db, err := gorm.Open("mysql", "root:12345678@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return &database{
		DB: db,
	}
}

func (db *database) getBook(ID int) models.Book {
	var book models.Book
	db.DB.First(&book, ID)

	return book
}

func (db *database) getBooks() []models.Book {
	return nil
}
