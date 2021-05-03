package array

import (
	"testing"
)

func Test_findPairs(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args:args{
				nums: []int{3,1,4,1,5},
				k:    2,
			},
			want:2,
		},
		{
			args:args{
				nums: []int{1, 2, 3, 4, 5},
				k:    1,
			},
			want:4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPairs(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}