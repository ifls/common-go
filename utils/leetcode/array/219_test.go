package array

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{[]int{1, 2, 3, 1, 2, 3}, 2},
			want: false,
		},
		{
			args: args{[]int{1, 0, 1, 1}, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMatch(t *testing.T) {
	type args struct {
		i int
		j int
		k int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{-1, 2, 1},
			want: false,
		},
		{
			args: args{-1, 2, 4},
			want: true,
		},
		{
			args: args{1, -2, 1},
			want: false,
		},
		{
			args: args{1, 6, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.i, tt.args.j, tt.args.k); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
