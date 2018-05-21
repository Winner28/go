package server

import (
	"net/http"
	"strconv"
	"web/server/models"

	"github.com/gin-gonic/gin"
)

type server struct {
	router *gin.Engine
	db     *database
}

// StartServer func
func StartServer() {
	s := &server{
		router: gin.Default(),
		db:     getDB(),
	}
	s.routes()
	s.router.Run("localhost:8080")
}

// SetUp routes
func (s *server) routes() {
	bookRoutes := s.router.Group("/books")
	s.setUpBookRoutes(bookRoutes)

}

func (s *server) setUpBookRoutes(bookRoutes *gin.RouterGroup) {
	bookRoutes.POST("/", s.createBook)
	bookRoutes.GET("/", s.fetchBooks)
	bookRoutes.GET("/:id", s.fetchBook)
	bookRoutes.PUT("/:id", s.updateBook)
	bookRoutes.DELETE("/:id", s.deleteBook)
}

func (s *server) createBook(c *gin.Context) {
	book := getBookModel(c.PostForm("title"), c.PostForm("author"))
	s.db.createBook(&book)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated,
		"message":    "Book created!",
		"resourceId": book.ID})

}

func (s *server) fetchBooks(c *gin.Context) {
	books := s.db.fetchBooks()
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   books,
	})
}

func (s *server) fetchBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	book := s.db.fetchBook(bookID)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   book,
	})

}

func (s *server) updateBook(c *gin.Context) {

}

func (s *server) deleteBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.PostForm("ID"))
	status, err := s.db.deleteBook(bookID)
	c.JSON(status, gin.H{
		"status":  status,
		"message": err.Error,
	})
}

func getBookModel(title, author string) models.Book {
	return models.Book{
		Title:  title,
		Author: author,
	}
}
