package main

import (
	"fmt"
)

// Solution for: https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	total := 0
	for i, j := range nums {
		for k, l := range nums {
			total = l + j
			if total == target {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

func main() {

	a := []int{2, 7, 11, 15}
	fmt.Println(twoSum(a, 9))
}
