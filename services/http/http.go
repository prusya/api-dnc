package http

import "github.com/gin-gonic/gin"

type Service struct {
	router *gin.Engine
}

func New() *Service {
	r := gin.Default()
	s := Service{
		router: r,
	}
	s.routes()

	return &s
}

func (s *Service) Start() {
	s.router.Run()
}

func (s *Service) Stop() {
}
