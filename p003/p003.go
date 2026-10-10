package p003

import "math"

/* no full proof for this one, nothing interesting is happening here
*  prime: divisors are tried in increasing order, so any divisor that divides factor is prime
*  (its own prime factors were already divided out). the loop stops once divisor > sqrt(factor),
*  so the remaining factor has no divisor <= its square root, which makes it prime
*  largest: we only divide while factor >= divisor^2, so factor/divisor >= divisor. the remaining
*  factor never drops below a prime already divided out, so it is the largest prime factor
 */
func Solve(composite int) uint64 {
	factor := uint64(composite)
	var divisor uint64 = 2

	for divisor <= uint64(math.Floor(math.Sqrt(float64(factor)))) {
		if factor%divisor == 0 {
			factor = factor / divisor
		} else {
			divisor++
		}
	}
	return factor
}
