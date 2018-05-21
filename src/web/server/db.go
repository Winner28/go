package server

import (
	"errors"
	"net/http"
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

func (db *database) fetchBook(ID int) models.Book {
	var book models.Book
	if err := db.DB.First(&book, ID).Error; err != nil {
		panic(err)
	}

	return book
}

func (db *database) fetchBooks() []models.Book {
	var books []models.Book
	if err := db.DB.Find(&books).Error; err != nil {
		panic(err)
	}

	return books
}

func (db *database) createBook(book *models.Book) {
	db.DB.Save(&book)
}

func (db *database) deleteBook(ID int) (status int, err error) {

	book := db.fetchBook(ID)
	if book.ID == 0 {
		return http.StatusNotFound, errors.New("We dont have book with such ID")
	}
	if err = db.DB.Delete(&book).Error; err != nil {
		return http.StatusInternalServerError, errors.New("Internal Server Error")
	}

	return http.StatusOK, nil
}

func (db *database) updateBook(book models.Book) {
	panic("Not implemented")
}
