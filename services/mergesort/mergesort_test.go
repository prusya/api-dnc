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
