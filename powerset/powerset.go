package main

import (
	"math"
)

// https://leetcode.com/problems/subsets

// Time Complexity: O(n2^n)
func powerset(nums []int) [][]int {
	n := len(nums)
	pSize := int(math.Pow(2, float64(n)))
	result := make([][]int, 0, pSize)

	for counter := 0; counter < pSize; counter++ {
		var subset []int
		for j := 0; j < n; j++ {
			if counter&(1<<uint(j)) == 0 {
				subset = append(subset, nums[j])
			}
		}
		result = append(result, subset)
	}

	return result
}

func powerset2(arr []string) [][]string {
	n := len(arr)
	pSize := int(math.Pow(2, float64(n)))
	powerset := make([][]string, 0, pSize)

	for counter := 0; counter < pSize; counter++ {
		var subset []string
		for j := 0; j < n; j++ {
			if counter&(1<<uint(j)) > 0 {
				subset = append(subset, arr[j])
			}
		}

		powerset = append(powerset, subset)
	}

	return powerset

}

/*

Consider follwoing example

arr = [1, 2, 3]
n=3
pSzie=2^3=8 (bitmask-len = n => 000)
result = [][]int
result = [[], [1], [2], [1,2], [3], [1,2], [2,3], [1,2,3]]

  counter = 0, j = 0, arr[j] = 1 {000 & (1<<0)} => 000 > 0  subset = []
  counter = 0, j = 1, arr[j] = 2 {000 & (1<<1)} => 000 > 0  subset = []
  counter = 0, j = 2, arr[j] = 3 {000 & (1<<2)} => 000 > 0  subset = []
  counter = 0, j = 3 j < 3 so exit j loop
append the new subset to powerset => result = [[]]


  counter = 1, j = 0, arr[j] = 1 {001 & (1<<0)} => 001 > 0  subset = [1]
  counter = 1, j = 1, arr[j] = 2 {001 & (1<<1)} => 000 > 0  subset = [1]
  counter = 1, j = 2, arr[j] = 3 {001 & (1<<2)} => 000 > 0  subset = [1]
  counter = 1, j = 3 j < 3 so exit j loop
append the new subset to powerset => result = [[], [1]]


counter = 2, j = 0, arr[j] = 1 {010 & (1<<0)} => 000 > 0  subset = []
counter = 2, j = 1, arr[j] = 2 {010 & (1<<1)} => 010 > 0  subset = [2]
counter = 2, j = 2, arr[j] = 3 {010 & (1<<2)} => 000 > 0  subset = [2]
counter = 2, j = 3 j < 3 so exit j loop
append the new subset to powerset => result = [[], [1], [2]]


counter = 3, j = 0, arr[j] = 1 {011 & (1<<0)} => 001 > 0  subset = [1]
counter = 3, j = 1, arr[j] = 2 {011 & (1<<1)} => 010 > 0  subset = [1,2]
counter = 3, j = 2, arr[j] = 3 {011 & (1<<1)} => 000 > 0  subset = [1,2]
counter = 3, j = 3 j < 3 so exit j loop
append the new subset to powerset => result = [[], [1], [2], [1,2]]

counter = 4, j = 0, arr[j] = 1 {100 & (1<<0)} => 000 > 0  subset = []
counter = 4, j = 1, arr[j] = 2 {100 & (1<<1)} => 000 > 0  subset = []
counter = 4, j = 2, arr[j] = 3 {100 & (1<<2)} => 100 > 0  subset = [3]
counter = 4, j = 3 j < 3 so exit j loop
append the new subset to powerset => result = [[], [1], [2], [1,2], [3]]

counter = 5, j = 0, arr[j] = 1 {101 & (1<<0)} => 001 > 0  subset = [1]
counter = 5, j = 1, arr[j] = 2 {101 & (1<<1)} => 000 > 0  subset = [1]
counter = 5, j = 2, arr[j] = 3 {101 & (1<<2)} => 100 > 0  subset = [1,3]
counter = 5, j = 3 j < 3 so exit j loop
append the new subset to powerset => result = [[], [1], [2], [1,2], [3], [1,3]]

counter = 6, j = 0, arr[j] = 1 {110 & (1<<0)} => 000 > 0  subset = []
counter = 6, j = 1, arr[j] = 2 {110 & (1<<1)} => 010 > 0  subset = [2]
counter = 6, j = 2, arr[j] = 3 {110 & (1<<2)} => 100 > 0  subset = [2,3]
counter = 6, j = 3 j < 3 so exit j loop
append the new subset to powerset => result = [[], [1], [2], [1,2], [3], [1,3], [2,3]]

counter = 7, j = 0, arr[j] = 1 {111 & (1<<0)} => 001 > 0  subset = [1]
counter = 7, j = 1, arr[j] = 2 {111 & (1<<1)} => 010 > 0  subset = [1,2]
counter = 7, j = 2, arr[j] = 3 {111 & (1<<2)} => 100 > 0  subset = [1,2,3]
counter = 7, j = 3 j < 3 so exit j loop
append the new subset to powerset => result = [[], [1], [2], [1,2], [3], [1,3], [2,3], [1,2,3]]

counter = 8 counter < 8 so exit the counter loop

return the result array => result = [[], [1], [2], [1,2], [3], [1,3], [2,3], [1,2,3]]
*/
