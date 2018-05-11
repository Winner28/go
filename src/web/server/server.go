package server

import (
	"net/http"
	"strconv"

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

}

func (s *server) fetchBooks(c *gin.Context) {

}

func (s *server) fetchBook(c *gin.Context) {
	bookID, _ := strconv.Atoi(c.Param("id"))
	book := s.db.getBook(bookID)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   book,
	})

}

func (s *server) updateBook(c *gin.Context) {

}

func (s *server) deleteBook(c *gin.Context) {

}
