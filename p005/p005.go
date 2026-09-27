package p005

func Solve(limit int) uint64 {
	// this question is the same as LCM(1,...,limit)
	// which is the same as LCM(2,..., limit)
	result := 1
	for i := 2; i <= limit; i++ {
		result = lcm(result, i)
	}
	return uint64(result)
}

// Greatest Common Divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Least Common Multiple (LCM) via GCD
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// brute force
// func Solve(limit int) uint64 {
// 	// TODO: solve problem 5
// outer:
// 	for i := limit; ; i++ {
// 		for x := 2; x <= limit; x++ {
// 			if i%x != 0 {
// 				continue outer
// 			}
// 		}
// 		return uint64(i)
// 	}
// }
