package p006

// this is a reduced form of the expression, see proof at p006_proof.tex
func Solve(n int) uint64 {
	// TODO: solve problem 6
	return uint64((n*n - 1) * n * (3*n + 2) / 12)
}
