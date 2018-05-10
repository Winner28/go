package server

import "github.com/gin-gonic/gin"

type server struct {
	router *gin.Engine
}

var s *server

func init() {
	s = &server{
		router: gin.Default(),
	}
}
