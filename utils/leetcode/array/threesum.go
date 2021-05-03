package array

func threeSum(a []int) [][]int {
	result := make([][]int, 0)
	leng := len(a)
	for i := 0; i < leng-2; i++ {
		for j := i + 1; j < leng-1; j++ {
			for k := j + 1; k < leng; k++ {
				if a[i]+a[j]+a[k] == 0 {
					if i != j && j != k {
						result = append(result, []int{i, j, k})
					}
				}
			}
		}
	}

	return result
}
