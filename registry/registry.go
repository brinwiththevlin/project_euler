package registry

import "fmt"

// Runner defines a standard signature for any Euler solution.
// It takes a raw string input and returns a string answer or an error.
type Runner func(string) (string, error)

// runners maps problem numbers to their respective execution functions.
var runners = map[int]Runner{}

// Register adds a problem's Runner. Call it from the problem package's init().
func Register(n int, r Runner) {
	if _, dup := runners[n]; dup {
		panic(fmt.Sprintf("problem %d registered twice", n))
	}
	runners[n] = r
}

// Get returns the Runner for problem n, if one has been registered.
func Get(n int) (Runner, bool) {
	r, ok := runners[n]
	return r, ok
}
