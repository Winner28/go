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
}

func (s *server) setUpBookRoutes(bookRoutes *gin.RouterGroup) {

}
