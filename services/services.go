package services

var (
	Mergesort MergesortService
	Http      HttpService
)

type MergesortService interface {
	DistributedSort([]int) ([]int, error)
	JobQueueSort([]int) (string, error)
	JobQueueSortResults(string) ([]int, error)
}

type HttpService interface {
	Start()
	Stop()
}
