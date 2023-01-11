package models

type MergesortResult struct {
	Host       string
	Result     []int
	Request    []int
	SubResults []*MergesortResult
}
