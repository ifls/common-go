package leetcode

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type data struct {
	arr []int
	k   int
	dst []int
}

var testCase189 = []data{
	{
		arr: []int{1, 2, 3, 4, 5, 6, 7},
		k:   3,
		dst: []int{5, 6, 7, 1, 2, 3, 4},
	},
	{
		arr: []int{-1},
		k:   2,
		dst: []int{-1},
	},
	{
		arr: []int{1, 2},
		k:   3,
		dst: []int{2, 1},
	},
}

func TestRun_189(t *testing.T) {
	for _, test := range testCase189 {
		rotate(test.arr, test.k)
		log.Printf("%v \n", test.arr)
		assert.Equal(t, test.dst, test.arr)
	}
}

func rotate(nums []int, k int) {
	swapArray(nums, 0, len(nums))
	swapArray(nums, 0, k%len(nums))
	swapArray(nums, k%len(nums), len(nums))
}

func swapArray(nums []int, start, end int) {
	for i := start; i < (end-start)/2+start; i++ {
		right := end - 1 - (i - start)
		temp := nums[i]
		nums[i] = nums[right]
		nums[right] = temp
	}
}
