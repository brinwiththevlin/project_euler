package p004

import (
	"math"
	"strconv"
)

func Solve(digitCount int) uint64 {
	// TODO: solve problem 4
	currentDigits := 0
	var bestPalindrome uint64 = 0

	for i := uint64(math.Pow10(digitCount)) - 1; i >= uint64(math.Pow10(digitCount-1)); i-- {
		for j := i; j >= uint64(math.Pow10(digitCount-1)); j-- {
			num := i * j
			num_str := strconv.FormatUint(num, 10)

			if len(num_str) < currentDigits {
				// the current palindrome is longer than this number, don't bother checking
				continue
			}

			if isPalindrome(num_str) && num > bestPalindrome {
				bestPalindrome = num
				currentDigits = len(num_str)
			}

		}
	}
	return bestPalindrome
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
