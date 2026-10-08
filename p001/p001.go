package p001

func Solve(limit int) uint64 {
	if limit < 1 {
		return 0
	}
	ulimit := uint64(limit)
	return s(3, ulimit) + s(5, ulimit) - s(15, ulimit)
}

func s(d, n uint64) uint64 {
	upper := (n - 1) / d
	t := upper * (upper + 1) / 2
	return d * t
}
