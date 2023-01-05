package http

func (s *Service) routes() {
	s.router.GET("/mergesort", getMergesort)
}
