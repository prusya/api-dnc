package mergesort

import (
	"reflect"
	"testing"
)

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
		{"3 1 2", s, args{[]int{3, 1, 2}}, []int{1, 2, 3}, false},
		{"empty", s, args{[]int{}}, []int{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DistributedSort(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Sort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
