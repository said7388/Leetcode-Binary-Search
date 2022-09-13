/*
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Follow up: Do not use any built-in library function such as sqrt.

*/

package main

func isPerfectSquare(num int) bool {

	if num == 1 {
		return true
	}

	l, r := 0, num

	for l < r {
		m := l + (r-l)/2
		if m*m < num {
			l = m + 1
		} else if m*m > num {
			r = m
		} else {
			return true
		}
	}

	return false
}
