/*
Given a characters array letters that is sorted in non-decreasing order and a character target, return the smallest character in the array that is larger than target.

Note that the letters wrap around.

For example, if target == 'z' and letters == ['a', 'b'], the answer is 'a'.

*/

package main

func nextGreatestLetter(letters []byte, target byte) byte {
	result := letters[0]
	l, r := 0, (len(letters) - 1)

	for l <= r {
		m := (l + r) / 2
		if letters[m] > target {
			result = letters[m]
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return result
}
