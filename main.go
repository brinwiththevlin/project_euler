package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/brinwiththevlin/project_euler/registry"
)

func main() {
	probNum := flag.Int("p", 0, "The Project Euler problem number to run")
	input := flag.String("i", "", "The problem-specific custom input data")
	flag.Parse()

	if *probNum == 0 {
		fmt.Println("Error: Please specify a problem number using the -p flag.")
		os.Exit(1)
	}

	runFn, exists := registry.Get(*probNum)
	if !exists {
		fmt.Printf("Error: Problem %d has not been registered yet.\n", *probNum)
		os.Exit(1)
	}

	fmt.Printf("--- Running Project Euler Problem %d ---\n", *probNum)
	start := time.Now()
	answer, err := runFn(*input)
	if err != nil {
		fmt.Printf("Execution failed: %v\n", err)
		os.Exit(1)
	}
	end := time.Now()

	fmt.Printf("Answer: %s\n", answer)
	fmt.Printf("calculation time: %v\n", end.Sub(start).Seconds())
}
