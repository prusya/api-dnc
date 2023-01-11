package mergesort

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/prusya/api-dnc/config"
	"github.com/prusya/api-dnc/models"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) DistributedSort(arr []int) (*models.MergesortResult, error) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(arr) < 2 {
		out := models.MergesortResult{
			Host:    hostname,
			Result:  arr,
			Request: arr,
		}
		return &out, nil
	}

	left, err := s.distributedSort(arr[:len(arr)/2])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	right, err := s.distributedSort(arr[len(arr)/2:])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res := merge(left.Result, right.Result)
	out := models.MergesortResult{
		Host:       hostname,
		Result:     res,
		Request:    arr,
		SubResults: []*models.MergesortResult{left, right},
	}

	return &out, nil
}

func merge(a, b []int) []int {
	out := []int{}

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	for i < len(a) {
		out = append(out, a[i])
		i++
	}
	for j < len(b) {
		out = append(out, b[j])
		j++
	}

	return out
}

func (s *Service) JobQueueSort(arr []int) (string, error) {
	return "", nil
}

func (s *Service) JobQueueSortResults(jobID string) ([]int, error) {
	return []int{}, nil
}

func (s *Service) distributedSort(arr []int) (*models.MergesortResult, error) {
	var out *models.MergesortResult
	var err error

	if config.MergesortUrl != "" {
		payload := fmt.Sprintf("%#v", arr)
		payload = strings.TrimPrefix(payload, "[]int{")
		payload = strings.TrimSuffix(payload, "}")
		payload = strings.ReplaceAll(payload, " ", "")
		urlLeft := fmt.Sprintf("%s?arr=%s", config.MergesortUrl, payload)
		respLeft, err := http.Get(urlLeft)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer respLeft.Body.Close()
		err = json.NewDecoder(respLeft.Body).Decode(&out)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		out, err = s.DistributedSort(arr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return out, nil
}
