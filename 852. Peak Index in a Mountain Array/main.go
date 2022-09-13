/*
An array arr a mountain if the following properties hold:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
Given a mountain array arr, return the index i such that arr[0] < arr[1] < ... < arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].

You must solve it in O(log(arr.length)) time complexity.

*/

package main

func peakIndexInMountainArray(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if (nums[mid] > nums[mid-1]) && (nums[mid] > nums[mid+1]) {
			return mid
		} else if nums[mid] > nums[mid-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return 0
}
