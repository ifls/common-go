package array

func findDisappearedNumbers1(nums []int) []int {
	cache := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		cache[nums[i]] = struct{}{}
	}

	ret := []int{}

	for i := 1; i <= len(nums); i++ {
		if _, ok := cache[i]; !ok {
			ret = append(ret, i)
		}
	}

	return ret
}

func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0, len(nums))
	for _, value := range nums {
		if value < 0 {
			value = -value
		}
		if nums[value-1] > 0 {
			nums[value-1] = -nums[value-1]
		}
	}
	for index, value := range nums {
		if value > 0 {
			result = append(result, index+1)
		}
	}
	return result
}
