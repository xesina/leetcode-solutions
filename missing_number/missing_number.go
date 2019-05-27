package missingnumber

import (
	"sort"
)
// https://leetcode.com/problems/missing-number/

// solution using sort
func missingNumber(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    i := 0
    for i=0; i<n; i++ {
        if nums[i] == i {
            continue
        }
         return i     
    }
    return i
}

// using formula n*(n+1)/2
func missingNumber2(nums []int) int { 
    var sum int
    n := len(nums)
    
    for _, i := range nums {
        sum += i
    }
    
    return (n*(n+1)) / 2 - sum 
    
}
