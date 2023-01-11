package mergesort

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"2; 1", args{[]int{2}, []int{1}}, []int{1, 2}},
		{"2,3; 1", args{[]int{2, 3}, []int{1}}, []int{1, 2, 3}},
		{"2; 1,3", args{[]int{2}, []int{1, 3}}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DistributedSort(t *testing.T) {
	type args struct {
		arr []int
	}
	s := New()
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    []int
		wantErr bool
	}{
		{"empty", s, args{[]int{}}, []int{}, false},
		{"2,1", s, args{[]int{2, 1}}, []int{1, 2}, false},
		{"1,2,3", s, args{[]int{1, 2, 3}}, []int{1, 2, 3}, false},
		{"3,2,1", s, args{[]int{3, 2, 1}}, []int{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DistributedSort(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.DistributedSort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Result, tt.want) {
				t.Errorf("Service.DistributedSort() = %v, want %v", got.Result, tt.want)
			}
		})
	}
}

func TestService_distributedSort(t *testing.T) {
	type args struct {
		arr []int
	}
	s := New()
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    []int
		wantErr bool
	}{
		{"empty", s, args{[]int{}}, []int{}, false},
		{"2,1", s, args{[]int{2, 1}}, []int{1, 2}, false},
		{"1,2,3", s, args{[]int{1, 2, 3}}, []int{1, 2, 3}, false},
		{"3,2,1", s, args{[]int{3, 2, 1}}, []int{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.distributedSort(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.distributedSort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Result, tt.want) {
				t.Errorf("Service.distributedSort() = %v, want %v", got.Result, tt.want)
			}
		})
	}
}
