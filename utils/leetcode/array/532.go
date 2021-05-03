package array

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	numsHas := make(map[int]bool)
	diffHas := make(map[int]bool)

	for _, num := range nums {
		if numsHas[num-k] {
			diffHas[num-k] = true
		}
		if numsHas[num+k] {
			diffHas[num] = true
		}
		numsHas[num] = true
	}
	return len(diffHas)
}
