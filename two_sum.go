package main

import (
	"fmt"
	"slices"
)

func twoSum(nums []int, target int) []int {
	nums_len := len(nums)
	result := []int{0, 0}

	for i := 0; i < nums_len; i++ {
		rem := target - nums[i]
		rem_index := slices.Index(nums, rem)

		if rem_index != -1 && rem_index != i {
			result = []int{i, rem_index}
			break
		}

	}
	return result
}

func main() {
	nums := []int{3, 2, 4}
	target := 6

	fmt.Println(twoSum(nums, target))
}
