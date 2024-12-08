package main

import (
	"fmt"
	"aoc/day2/task"
	"os"
)

func main() {
	input, err:= os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	result1 := task.Task1(string(input))
	fmt.Printf("Task 1: %d\n", result1)
	result2 := task.Task2(string(input))
	fmt.Printf("Task 2: %d\n", result2)
}