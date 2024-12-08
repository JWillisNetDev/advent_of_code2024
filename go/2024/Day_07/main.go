package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Error(s string, e ...any) {
	fmt.Fprintf(os.Stderr, s, e...)
	os.Exit(1)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		Error("Error reading input: %s", err)
	}

	Part1And2(input)
}

func Part1And2(input []byte) {
	c := ParseCalibrations(input)
	acc := int64(0)
	for _, cc := range c {
		acc += cc.GetIsValidValue()
	}

	fmt.Printf("Part 1: %d\n", acc)
}

type Calibration struct {
	result int64
	root   *Node
}

type Node struct {
	v          int64
	nAdd, nMul *Node
	nCat       *Node
}

func ParseCalibrations(input []byte) []Calibration {
	buf := bufio.NewScanner(bytes.NewBuffer(input))
	arr := []Calibration{}
	for buf.Scan() {
		l := buf.Bytes()
		c := ParseLine(l)
		arr = append(arr, c)
	}
	return arr
}

func GetTips(r *Node) []*Node {
	// There should never be a state where both add and mul nodes are nil
	if r.nAdd == nil && r.nMul == nil {
		return []*Node{r}
	}
	return append(GetTips(r.nAdd), append(GetTips(r.nMul), GetTips(r.nCat)...)...)
}

func AddNode(r *Node, v int64) {
	tips := GetTips(r)
	for _, n := range tips {
		n.nAdd = &Node{v: n.v + v}
		n.nMul = &Node{v: n.v * v}
		n.nCat = &Node{v: ConcatInt(n.v, v)}
	}
}

func ConcatInt(l, r int64) int64 {
	s := fmt.Sprintf("%d%d", l, r)
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		Error("Failed to concatenate numbers %d, %d (%s) into single number: %s", l, r, s, err)
	}
	return n
}

func (c Calibration) GetIsValidValue() int64 {
	for _, n := range GetTips(c.root) {
		if n.v == c.result {
			return c.result
		}
	}
	return 0
}

var CalibrationRegex = regexp.MustCompile(`^([0-9]+): `)

func ParseLine(line []byte) Calibration {
	matches := CalibrationRegex.FindSubmatch(line)

	res, err := strconv.ParseInt(string(matches[1]), 10, 64)
	if err != nil {
		Error("Error converting the result part %s of the following line to a number: %s - err %s", matches[1], line, err)
	}

	rest := line[len(matches[0]):]
	exploded := bytes.Split(rest, []byte{' '})
	p, err := strconv.ParseInt(string(exploded[0]), 10, 64)
	if err != nil {
		Error("Failed to parse %s following to an int64: %s", exploded[0], err)
	}
	root := &Node{v: p}

	for _, l := range exploded[1:] {
		p, err := strconv.ParseInt(string(l), 10, 64)
		if err != nil {
			Error("Error converting the parameter %s of the line %s to a number: %s", l, line, err)
		}
		AddNode(root, p)
	}

	c := Calibration{res, root}
	return c
}
