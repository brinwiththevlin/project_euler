package p002

func Solve(limit int) uint64 {
	var sum uint64 = 0
	var f1 uint64 = 1
	var f2 uint64 = 2

	for f2 < uint64(limit) {
		if f2%2 == 0 {
			sum += f2
		}
		f1, f2 = f2, f1+f2
	}
	return sum
}
