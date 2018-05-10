package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type server struct {
	router *gin.Engine
	db     *gorm.DB
}

var s *server

func init() {
	s = &server{
		router: gin.Default(),
		db:     getDB(),
	}
}

func (s *server) routes() {
	bookRoutes := s.router.Group("/books")
	s.setUpBookRoutes(bookRoutes)
}

func (s *server) setUpBookRoutes(bookRoutes *gin.RouterGroup) {
	bookRoutes.POST("/", createBook)
	bookRoutes.GET("/", fetchBooks)
	bookRoutes.GET("/:id", fetchBook)
	bookRoutes.PUT("/:id", updateBook)
	bookRoutes.DELETE("/:id", deleteBook)
}

func createBook(c *gin.Context) {

}

func fetchBooks(c *gin.Context) {

}

func fetchBook(c *gin.Context) {

}

func updateBook(c *gin.Context) {

}

func deleteBook(c *gin.Context) {

}
