package mergesort

// https://leetcode.com/problems/sort-an-array/

func sortArray(nums []int) []int {
	return mergeSort(nums)
 }
 
 func mergeSort(nums []int) []int  {
	 if len(nums) < 2 {
		 return nums
	 }
	 
	 m := len(nums) / 2
 
	 l := mergeSort(nums[:m])
	 r := mergeSort(nums[m:])
	 
	 return merge(l, r)
 }
 
 func merge(l, r []int) []int {
	 tmp := make([]int, len(r)+len(l))
	 
	 leftIndex := 0
	 rightIndex := 0
	 
	 for i:= 0; i<len(tmp); i++ {
		 if leftIndex > len(l)-1 && rightIndex <= len(r)-1 {
			 tmp[i] = r[rightIndex]
			 rightIndex++
		 } else if rightIndex > len(r)-1 && leftIndex <= len(l) {
			 tmp[i] = l[leftIndex]
			 leftIndex++
		 } else if l[leftIndex] <= r[rightIndex] {
			 tmp[i] = l[leftIndex]
			 leftIndex++
		 } else {
			 tmp[i] = r[rightIndex]
			 rightIndex++
		 }
	 }
 
	 return tmp
 }