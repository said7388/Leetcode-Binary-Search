/*
Given two integer arrays arr1 and arr2, and the integer d, return the distance value between the two arrays.

The distance value is defined as the number of elements arr1[i] such that there is not any element arr2[j] where |arr1[i]-arr2[j]| <= d.

*/

package main

import "sort"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr1)
	sort.Ints(arr2)

	var i, j, res int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j]+d {
			if arr1[i]+d < arr2[j] {
				res++
			}
			i++
		} else {
			j++
		}
	}
	return res + (len(arr1) - i)
}
