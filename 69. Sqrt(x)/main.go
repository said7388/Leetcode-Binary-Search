/*
Given a non-negative integer x, compute and return the square root of x.

Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.

Note: You are not allowed to use any built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.

*/

package main

func mySqrt(x int) int {
	l, r := 0, x
	result := 0

	for l <= r {
		m := (l + r) / 2
		if m*m > x {
			r = m - 1
		} else {
			result = m
			l = m + 1
		}
	}

	return result
}
