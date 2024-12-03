package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	
	// fuck you, regex time.
	regex := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)

	scanner := bufio.NewScanner(file)
	acc := 0
	shouldDo := true
	for scanner.Scan() {
		line := scanner.Text()
		for _, match := range regex.FindAllString(line, -1) {
			var lhs, rhs int
			if n, _ := fmt.Sscanf(match, "mul(%d,%d)", &lhs, &rhs); n > 0 && shouldDo {
				acc = acc + (lhs * rhs)
			} else if match == "do()" {
				shouldDo = true
			} else if match == "don't()" {
				shouldDo = false
			}
		}
	}

	fmt.Printf("Result: %d\n", acc)
}