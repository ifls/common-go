package array

func findUnsortedSubarray(nums []int) int {
	r := len(nums) - 1

	if r <= 0 {
		return 0
	}

	i := 0
	for i < r && nums[i] <= nums[i+1] {
		i++
	}

	for i < r && nums[r-1] < nums[r] {
		r--
	}
	print(i, " ", r)

	if r == i {
		return r - i
	}

	return r - i + 1
}
