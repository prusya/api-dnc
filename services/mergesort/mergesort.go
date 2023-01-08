package mergesort

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/prusya/api-dnc/config"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) DistributedSort(arr []int) ([]int, error) {
	if len(arr) < 2 {
		return arr, nil
	}

	var left, right []int
	var err error
	if config.MergesortUrl != "" {
		arrLeft := fmt.Sprintf("%#v", arr[:len(arr)/2])
		arrLeft = strings.TrimPrefix(arrLeft, "[]int{")
		arrLeft = strings.TrimSuffix(arrLeft, "}")
		arrLeft = strings.ReplaceAll(arrLeft, " ", "")
		urlLeft := fmt.Sprintf("%s?arr=%s", config.MergesortUrl, arrLeft)
		respLeft, err := http.Get(urlLeft)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer respLeft.Body.Close()
		err = json.NewDecoder(respLeft.Body).Decode(&left)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		arrRight := fmt.Sprintf("%#v", arr[len(arr)/2:])
		arrRight = strings.TrimPrefix(arrRight, "[]int{")
		arrRight = strings.TrimSuffix(arrRight, "}")
		arrRight = strings.ReplaceAll(arrRight, " ", "")
		urlRight := fmt.Sprintf("%s?arr=%s", config.MergesortUrl, arrRight)
		respRight, err := http.Get(urlRight)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		defer respRight.Body.Close()
		err = json.NewDecoder(respRight.Body).Decode(&right)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		left, err = s.DistributedSort(arr[:len(arr)/2])
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		right, err = s.DistributedSort(arr[len(arr)/2:])
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	// fmt.Println("left", left)
	// fmt.Println("right", right)

	return merge(left, right), nil
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
