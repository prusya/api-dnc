package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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
	err := s.router.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Service) Stop() {
}
