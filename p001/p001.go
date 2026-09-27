package p001

func Solve(limit int) uint64 {
	ulimit := uint64(limit)
	seen := make(map[uint64]struct{})

	multiples := map[uint64]uint64{3: 3, 5: 5}

	var sum uint64 = 0

	for anyBelowLimit(multiples, ulimit) {
		for k, v := range multiples {
			if _, ok := seen[v]; !ok && v < ulimit {
				sum += v
				seen[v] = struct{}{}
			}
			if v < ulimit {
				multiples[k] += k
			}
		}
	}
	return sum
}

func anyBelowLimit(m map[uint64]uint64, limit uint64) bool {
	for _, val := range m {
		if val < limit {
			return true
		}
	}
	return false
}
