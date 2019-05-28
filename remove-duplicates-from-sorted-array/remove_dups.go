package removedups

// remove dups from sorted array
func removeDuplicates(nums []int) int {
    length := len(nums)

    if length == 0 {
        return 0
    }
    
    // last not duplicate index
    i := 0
    
    for j := 1; j < length; j++ {
        if nums[j] != nums[i] {
            i++
            nums[i] = nums[j]
        }
    }

	return i + 1
    
}