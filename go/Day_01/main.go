package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

type OccurrenceMap map[int]int

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	task1(bytes)
	task2(bytes)
}

func task1(b []byte) {
	lefts, rights := getNumbers(b)
	slices.Sort(lefts)
	slices.Sort(rights)

	acc := float64(0)
	for i := len(lefts) - 1; i >= 0; i-- {
		acc += math.Abs(float64(lefts[i] - rights[i]))
	}

	fmt.Printf("Result(task1): %v\n", int(acc))
}

func task2(b []byte) {
	lefts, rights := getNumbers(b)
	occ := buildOccurrenceMap(rights)

	sim := 0
	for _, v := range lefts {
		sim = sim + getSimilarityScore(v, occ)
	}

	fmt.Printf("Result(task2): %v\n", sim)
}

func getSimilarityScore(v int, occ OccurrenceMap) int {
	if vv, ok := occ[v]; ok {
		return v * vv
	}
	return 0
}

func buildOccurrenceMap(r []int) OccurrenceMap {
	m := make(OccurrenceMap)
	for _, v := range r {
		if vv, ok := m[v]; ok {
			m[v] = vv + 1
		} else {
			m[v] = 1
		}
	}
	return m
}

func getNumbers(b []byte) ([]int, []int) {
	s := bufio.NewScanner(bytes.NewReader(b))
	var lefts, rights []int

	for s.Scan() {
		var left, right int
		line := s.Text()
		n, err := fmt.Sscanf(line, "%d %d", &left, &right)
		if err != nil {
			log.Fatalf("Failed to parse line: %v", err)
		}
		if n != 2 {
			log.Fatalf("Expected 2 numbers, got %d", n)
		}
		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	return lefts, rights
}
