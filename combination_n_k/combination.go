package combination

// https://leetcode.com/problems/combinations/

func combine(n int, k int) [][]int {
    var result [][]int 
    for i := 0; i < (1 << uint(n)); i++ {
		g := gray(i)
		if countOnes(g) == k {
            var subset []int
			for j := 0; j < n; j++ {
				if g & (1 << uint(j)) > 0 {
                    subset = append(subset, j+1)
				}
			}
			
            result = append(result, subset)
		}
	}
	
	return result
}

func countOnes(n int) int {
	count := 0
	for ; n > 0; n >>= 1 {
		count += n & 1
	}

	return count
}

func gray(n int) int {
	return n ^ (n >> 1)
}