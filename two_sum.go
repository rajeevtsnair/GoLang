/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	length := len(nums)
	var result []int
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if target == (nums[i] + nums[j]) {
				result = []int{i, j}
				break
			}

		}
	}
	return result
}

func main() {

	nums := []int{2, 11, 7, 15}
	target := 9
	var result []int
	result = twoSum(nums, target)
	fmt.Println(result)
}
