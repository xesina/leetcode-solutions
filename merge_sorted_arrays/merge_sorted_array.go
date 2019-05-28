package mergesortedarray


// https://leetcode.com/problems/merge-sorted-array/

// we maintain 3 pointers :
// * i for track nums1 array items
// * j for track nums2 array items
// * k for track space available at the end of nums1 to insert to

func merge(nums1 []int, m int, nums2 []int, n int) {

    i := m - 1
    j := n - 1
    k := m + n - 1
    
    for i >= 0 && j >= 0 {
        if nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
        k--
    }

    for j >= 0 {
        nums1[k] = nums2[j]
        k--
        j--
    }
}
