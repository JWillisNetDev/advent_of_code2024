package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Encountered an error: %s\n", err)
		os.Exit(1)
	}
	Task1(input)
	Task2(input)
}

func Task1(input []byte) {
	buf := bufio.NewScanner(bytes.NewBuffer(input))
	rules := ParseRules(buf)
	updates := ParseUpdates(buf)
	acc := 0
	for _, u := range updates {
		acc += u.GetSumValidForRuleSet(rules)
	}
	fmt.Printf("Task 1: %d\n", acc)
}

func Task2(input []byte) {
	buf := bufio.NewScanner(bytes.NewBuffer(input))
	rules := ParseRules(buf)
	updates := ParseUpdates(buf)
	acc := 0
	for _, u := range updates {
		acc += u.GetSumInvalidForRuleSet(rules)
	}
	fmt.Printf("Task 2: %d\n", acc)
}

type Rule struct {
	mustComeBefore []int
}

type RuleSet map[int]*Rule

func ParseRules(s *bufio.Scanner) RuleSet {
	rs := make(RuleSet)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 || line == "\n" {
			break
		}

		var l, r int
		n, err := fmt.Sscanf(line, "%d|%d", &l, &r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse a line: %s\n", line)
			os.Exit(1)
		}
		if n != 2 {
			fmt.Fprintf(os.Stderr, "Failed to parse 2 arguments from line: %s\n", line)
			os.Exit(1)
		}

		if pg, ok := rs[l]; ok {
			pg.mustComeBefore = append(pg.mustComeBefore, r)
		} else {
			rs[l] = NewRule(l, []int{r})
		}
	}

	return rs
}

func NewRule(v int, mstCmBfr []int) *Rule {
	return &Rule{
		mustComeBefore: mstCmBfr,
	}
}

type Update struct {
	pages []int
}

func ParseUpdates(s *bufio.Scanner) []*Update {
	updates := []*Update{}
	for s.Scan() {
		line := s.Bytes()
		nums := bytes.Split(line, []byte(","))
		u := &Update{pages: make([]int, len(nums))}
		for i, n := range nums {
			nn, err := strconv.Atoi(string(n))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to parse update: %s at %s\n", line, n)
				os.Exit(1)
			}

			u.pages[i] = nn
		}

		updates = append(updates, u)
	}

	return updates
}

func (u *Update) IsValidRuleSet(rs RuleSet) bool {
	before := make(map[int]struct{})
	for _, n := range u.pages {
		before[n] = struct{}{}
		pg, ok := rs[n]
		if !ok {
			continue
		}

		for _, pp := range pg.mustComeBefore {
			if _, ok := before[pp]; ok {
				return false
			}
		}
	}
	return true
}

func (u *Update) GetSumValidForRuleSet(rs map[int]*Rule) int {
	if !u.IsValidRuleSet(rs) {
		return 0
	}
	r := u.pages[(len(u.pages)-1)/2]
	return r
}

func (u *Update) GetSumInvalidForRuleSet(rs RuleSet) int {
	if u.IsValidRuleSet(rs) {
		return 0
	}

	sorted := append([]int{}, u.pages...)
	sort.Slice(sorted, func(i, j int) bool {
		r1, ok := rs[sorted[i]]
		if !ok {
			return false
		}
		return slices.Contains(r1.mustComeBefore, sorted[j])
	})
	return sorted[(len(sorted)-1)/2]
}
