package dp

// Time Complexity: O(n^2), Space Complexity: O(n)
func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxLen := 1 // 定义最大长度为1
	// d[i]表示以第i个数结尾的最长上升子序列的长度
	d := make([]int, len(nums))
	d[0] = 1                         // 初始化d[0],下标为0的数的最长递增子序列长度为1
	for i := 1; i < len(nums); i++ { // 遍历整个数组
		for j := 0; j < i; j++ { // 遍历第i个数字前面的所有数
			cur := 1
			if nums[i] > nums[j] { // nums[i]与i前面的数nums[j]一一对比
				cur = d[j] + 1 // 如果nums[i]比较大，则候选长度+1
			}
			d[i] = max(d[i], cur) // 取所有候选长度的最大值为第i个数字的最长递增长度
		}
		maxLen = max(maxLen, d[i]) // 更新最长递增子序列的长度
	}
	return maxLen
}

func lengthOfLIS2(nums []int) int {
	size := 0
	length := len(nums)
	tail := make([]int, length)

	for i := 0; i < length; i++ {
		if size == 0 || tail[size-1] < nums[i] {
			tail[size] = nums[i]
			size++
		} else {
			x, y := 0, size-1
			for x < y {
				mid := (x + y) / 2
				if tail[mid] < nums[i] {
					x = mid + 1
				} else {
					y = mid
				}
			}
			tail[x] = nums[i]
		}
	}
	return size
}
