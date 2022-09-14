/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

*/

package main

func searchRange(nums []int, target int) []int {
	min := -1
	max := -1
	l, r := 0, len(nums)-1

	for l <= r {
		if min > -1 && max > -1 {
			break
		}
		if nums[l] == target {
			min = l
		} else {
			l += 1
		}
		if nums[r] == target {
			max = r
		} else {
			r -= 1
		}
	}

	return []int{min, max}
}
