package duplicates

import "sort"

// https://leetcode.com/problems/contains-duplicate/solution/
// containsDuplicate1 uses map to solve the problem
func containsDuplicate1(nums []int) bool {
    m := make(map[int]int)
    
    for _, i:= range nums {
        if _, ok := m[i]; !ok {
            m[i] = 1
            continue
        }
        
        return true
        
    }
    
    return false
}

// containsDuplicate2 uses sort. After sorting, we can sweep the sorted array to find if there are any two consecutive duplicate elements.
func containsDuplicate2(nums []int) bool {
    sort.Ints(nums)
    n := len(nums)
    
    for i:=0; i<n-1;i++ {
        if nums[i] == nums[i+1] {
            return true
        }
    }
    
    return false
}

