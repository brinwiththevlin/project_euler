package p003

func Solve(composit int) uint64 {
	factor := uint64(composit)
	var divisor uint64 = 2

	for divisor < factor {
		if factor%divisor == 0 {
			factor = factor / divisor
		} else {
			divisor++
		}
	}
	// TODO: solve problem 3
	return factor
}
