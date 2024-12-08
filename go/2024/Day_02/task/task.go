package task

import (
	"bufio"
	"bytes"
	"log"
	"math"
	"strconv"
	"strings"
)

func Task1(input string) int {
	scanner := bufio.NewScanner(bytes.NewBufferString(input))
	acc := 0
	for scanner.Scan() {
		line := scanner.Text()

		numbers := explodeNumbers(line)

		if isValidSeq(numbers) {
			acc++
		}
	}

	return acc
}

func Task2(input string) int {
	scanner := bufio.NewScanner(bytes.NewBufferString(input))
	acc := 0
	for scanner.Scan() {
		line := scanner.Text()

		numbers := explodeNumbers(line)

		if isValidSeq2(numbers) {
			acc++
		}
	}

	return acc
}

func isValidSeq2(numbers []int) bool {
	if isValidSeq(numbers) {
		return true
	}

	for i := 0; i < len(numbers)-1; i++ {
		slice := append(numbers[:i], numbers[i+2:]...)
		if isValidSeq(slice) {
			return true
		}
	}

	return false
}

func isValidSeq(numbers []int) bool {
	isIncreasing, isDecreasing := true, true
	if len(numbers) == 1 {
		return true
	}

	if len(numbers) == 2 {
		diff := math.Abs(float64(numbers[1] - numbers[0]))
		return diff <= 3 && diff >= 1
	}

	for i := 0; i < len(numbers)-1; i++ {
		diff := float64(numbers[i+1] - numbers[i])

		if math.Abs(diff) < 1 || math.Abs(diff) > 3 {
			return false
		} else if diff > 0 {
			isDecreasing = false
		} else if diff < 0 {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func explodeNumbers(s string) []int {
	var numbers []int
	exploded := strings.Split(s, " ")
	for _, v := range exploded {
		if num, err := strconv.Atoi(v); err == nil {
			numbers = append(numbers, num)
		} else {
			log.Fatalf("Failed to convert %s to int", v)
			break
		}
	}
	return numbers
}
