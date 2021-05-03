package array

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{a: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{
				{-1, 0, 1},
				{-1, -1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
