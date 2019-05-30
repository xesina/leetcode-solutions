package permutations

import (
	"fmt"
)

// https://leetcode.com/problems/permutations

func heapPermutation(a []int, size int) {
	if size == 1 {
		fmt.Println(a)
	}

	for i := 0; i < size; i++ {
		heapPermutation(a, size-1)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func permute(nums []int) [][]int {
	n := len(nums)
	var result [][]int
	var helper func([]int, int, int)

	helper = func(nums []int, left, size int) {
		if left == size {
			t := make([]int, len(nums))
			copy(t, nums)
			result = append(result, t)
		} else {
			for i := left; i < size; i++ {
				nums[left], nums[i] = nums[i], nums[left]
				helper(nums, left+1, size)
				nums[left], nums[i] = nums[i], nums[left]
			}
		}
	}

	helper(nums, 0, n)
	return result
}

/*

consider following example

nums = [1,2,3]
n = 3
result = []
swap(index1, index2)

helper(nums, 0, 3)
    l=0 s=3 i=0 swap(0,0) => nums = [1,2,3]
    helper(nums, 1, 3)
        l=1 s=3 i=1 swap(1,1) => nums = [1,2,3]
        helper(nums, 2, 3)
            l=2 s=3 i=2 swap (2,2) => nums[1,2,3]
            helper(nums, 3, 3)
                3 == 3 => result = [[123]]
            l=2 s=3 i=2 swap(2,2) => nums[1,2,3]

            l=2 s=3 i=3 i>=3 exit loop
        l=1 s=3 i=1 swap(1,1) => nums = [1,2,3]

        l=1 s=3 i=2 swap(1, 2) => nums = [1,3,2]
        helper(nums, 2, 3)
            l=2 s=3 i=2 swap(2, 2) => nums = [1,3,2]
            helper(nums, 3, 3)
                3 == 3 => result = [[1,2,3], [1,3,2]]
            l=2 s=3 i=2 swap(2, 2) => nums = [1,3,2]
        l=1 s=3 i=2 swap(1, 2) => nums = [1,2,3]

        l=1 s=3 i=3 i>=3 exit loop
    l=0 s=3 i=0 swap(1,1) => nums = [1,2,3]

    l=0 s=3 i=1 swap(0, 1) => nums = [2,1,3]
    helper(nums, 1, 3)
        l=1 s=3 i=1 swap(1, 1) => nums = [2,1,3]
        helper(nums, 2, 3)
            l=2 s=3 i=2 swap(2, 2) => nums[2,1,3]
            helper(nums, 3, 3)
                3 == 3 => result = [[1,2,3], [1,3,2], [2,1,3]]
            l=2 s=3 i=2 swap(2, 2) => nums[2,1,3]
        l=1 s=3 i=1 swap(1, 1) => nums = [2,1,3]

        l=1 s=3 i=2 swap(1, 2) => nums = [2,3,1]
        helper(nums, 2, 3)
            l=2 s=3 i=2 swap(2, 2) => nums[2,3,1]
            helper(nums, 3, 3)
                3 == 3 => result = [[1,2,3], [1,3,2], [2,1,3], [2,3,1]]
            l=2 s=3 i=2 swap(2, 2) => nums[2,3,1]
        l=1 s=3 i=2 swap(1, 2) => nums = [2,1,3]

        l=1 s=3 i=3 i>=3 exit loop

    l=0 s=3 i=1 swap(0, 1) => nums = [1,2,3]

    l=0 s=3 i=2 swap(0, 2) => nums = [3,2,1]
    helper(nums, 1, 3)
        l=1 s=3 i=1 swap(1, 1) => nums = [3,2,1]
        helper(nums, 2, 3)
            2=1 s=3 i=2 swap(2, 2) => nums = [3,2,1]
            helper(nums, 3, 3)
                3 == 3 => result = [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,2,1]]
            l=2 s=3 i=2 swap(2, 2) => nums = [3,2,1]
        l=1 s=3 i=1 swap(1, 1) => nums = [3,2,1]

        l=1 s=3 i=2 swap(1, 2) => nums = [3,1,2]
        helper(nums, 2, 3)
            l=2 s=3 i=2 swap(2, 2) => nums = [3,1,2]
            helper(nums, 3, 3)
                3 == 3 => result = [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,2,1], [3,1,2]]
            l=2 s=3 i=2 swap(2, 2) => nums = [3,1,2]
        l=1 s=3 i=2 swap(1, 2) => nums = [3,2,1]
    l=0 s=3 i=2 swap(0, 2) => nums = [1,2,3]

    l=0 s=3 i=3 i>=3 exit loop

result = [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,2,1], [3,1,2]]

*/
