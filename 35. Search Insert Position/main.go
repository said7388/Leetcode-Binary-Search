/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/

package main

func searchInsert(nums []int, target int) int {
	aux := 0

	for i, n := range nums {

		if n >= target {
			return i
		}

		aux = i
	}

	return aux + 1
}
