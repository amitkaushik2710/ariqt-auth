package server

import (
	"ariqt-auth/pkg/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	r := gin.Default()
	srv := Server{
		router: r,
	}
	return &srv
}

func (s *Server) StartPort(port int) {
	host := fmt.Sprintf(":%d", port)
	err := s.router.Run(host)
	if err != nil {
		logger.Logger.Fatal("StartPort: failed to start server at port", zap.Int("port", port), zap.Error(err))
	}
}

func (s *Server) Routes(method, path string, handler func(c *gin.Context)) {
	switch method {
	case http.MethodGet:
		s.router.GET(path, handler)
	case http.MethodPost:
		s.router.POST(path, handler)
	case http.MethodPut:
		s.router.PUT(path, handler)
	case http.MethodDelete:
		s.router.DELETE(path, handler)
	}
}
