package services

import "github.com/prusya/api-dnc/models"

var (
	Mergesort MergesortService
	Http      HttpService
)

type MergesortService interface {
	DistributedSort([]int) (*models.MergesortResult, error)
	JobQueueSort([]int) (string, error)
	JobQueueSortResults(string) ([]int, error)
}

type HttpService interface {
	Start()
	Stop()
}
