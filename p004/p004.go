package p004

import (
	"math"
)

/* Again no clever tricks here, just using a reverse funciton instead of a string conversion
 *  If you feel like it, leverege teh fact that even digit palindromes must be divisible by 11
 */

func Solve(digitCount int) uint64 {
	// TODO: solve problem 4
	var bestPalindrome uint64 = 0
	upper := uint64(math.Pow10(digitCount))
	lower := uint64(math.Pow10(digitCount - 1))

	for i := upper - 1; i >= lower; i-- {
		for j := i; j >= lower; j-- {
			num := i * j

			if num < bestPalindrome {
				// the current palindrome is longer than this number, don't bother checking
				break
			}

			rev_num := reverse(num)
			if num == rev_num && num > bestPalindrome {
				bestPalindrome = num
			}
		}
	}
	return bestPalindrome
}

func reverse(num uint64) uint64 {
	var rev uint64 = 0
	for num > 0 {
		rev = rev*10 + num%10
		num = num / 10
	}
	return rev
}
