package api

import (
	"github.com/alcastic/pwebsite-backend/internal/persistence"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *persistence.Store
	router *gin.Engine
}

func NewServer(store *persistence.Store) *Server {
	server := &Server{
		store: store,
	}
	server.router = setupRouter(server)
	return server
}

func setupRouter(s *Server) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.GET("/messages", s.getMessages)
	router.POST("/messages", s.addMessage)

	return router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
