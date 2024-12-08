package input

import (
	"fmt"
	"os"
)

func Error(s string, e ...any) {
	fmt.Fprintf(os.Stderr, "\033[31m"+s+"\033[0m", e...)
	os.Exit(1)
}

func Success(s string, e ...any) {
	fmt.Fprintf(os.Stdout, "\033[32m"+s+"\033[0m", e...)
}

func GetInput() []byte {
	if len(os.Args) == 1 {
		Error("A filepath is required to run this day.")
	}
	if len(os.Args) < 2 {
		Error("Too many arguments provided. Please provide a filepath only.")
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		Error("Failed to read input file: %s", err)
	}

	return bytes
}
