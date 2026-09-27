package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/p062"
	// Import your problem packages here
)

// Runner defines a standard signature for any Euler solution.
// It takes a raw string input and returns a string answer or an error.
type Runner func(string) (string, error)

// registry maps problem numbers to their respective execution functions.
var registry = map[int]Runner{
	62: func(input string) (string, error) {
		target, err := strconv.Atoi(input)
		if err != nil {
			return "", fmt.Errorf("invalid permutation count: %w", err)
		}
		ans := p062.Solve(target)
		return strconv.FormatUint(ans, 10), nil
	},
}

func main() {
	probNum := flag.Int("p", 0, "The Project Euler problem number to run")
	input := flag.String("i", "", "The problem-specific custom input data")
	flag.Parse()

	if *probNum == 0 {
		fmt.Println("Error: Please specify a problem number using the -p flag.")
		os.Exit(1)
	}

	runFn, exists := registry[*probNum]
	if !exists {
		fmt.Printf("Error: Problem %d has not been registered yet.\n", *probNum)
		os.Exit(1)
	}

	fmt.Printf("--- Running Project Euler Problem %d ---\n", *probNum)
	answer, err := runFn(*input)
	if err != nil {
		fmt.Printf("Execution failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Answer: %s\n", answer)
}
