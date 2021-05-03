package array

func containsNearbyDuplicate(nums []int, k int) bool {
	cache := map[int]int{}

	for i := 0; i < len(nums); i++ {
		if _, ok := cache[nums[i]]; ok {
			if isMatch(i, cache[nums[i]], k) {
				print(i, cache[nums[i]])
				return true
			} else {
				cache[nums[i]] = i
			}
		} else {
			cache[nums[i]] = i
		}
	}
	return false
}

// 2
func isMatch(i, j, k int) bool {
	if i - j <= k && i - j >= -k {
		return true
	}

	return false
}
