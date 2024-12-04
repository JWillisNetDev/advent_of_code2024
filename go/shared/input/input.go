package input

import (
	"fmt"
	"os"
)

func GetInput() string {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go input.txt")
	}

	fp := os.Args[1]
	bytes, err := os.ReadFile(fp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	return string(bytes)
}