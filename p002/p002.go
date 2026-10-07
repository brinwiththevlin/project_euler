package p002

func Solve(limit int) uint64 {
	var sum uint64 = 2
	var e1 uint64 = 2
	var e2 uint64 = 8

	for e2 < uint64(limit) {
		sum += e2
		e1, e2 = e2, e1+4*e2
	}
	return sum
}
