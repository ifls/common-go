package dp

import "testing"

type args struct {
	nums []int
}

type test struct {
	name string
	args args
	want int
}

var tests []test

func init() {
	tests = []test{
		{
			want: 4,
			args: args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}},
		},
	}
}

func Test_lengthOfLIS(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS2(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
