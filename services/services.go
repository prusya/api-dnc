package services

var (
	Mergesort MergesortService
	Http      HttpService
)

type MergesortService interface {
	Sort([]int) ([]int, error)
}

type HttpService interface {
	Start()
	Stop()
}
