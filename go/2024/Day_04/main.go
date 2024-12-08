package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	Task1(input)
	Task2(input)
}

const (
	X byte = 'X'
	M byte = 'M'
	A byte = 'A'
	S byte = 'S'
)

func Task1(b []byte) {
	m := NewMap(b)
	acc := 0

	for _, row := range m.points {
		for _, p := range row {
			for _, x := range []int{-1, 0, 1} {
				for _, y := range []int{-1, 0, 1} {
					if x == 0 && y == 0 {
						continue
					}
					if m.TestR(&p, x, y, X) {
						acc++
					}
				}
			}
		}
	}
	fmt.Printf("Task 1: %d\n", acc)
}

func Task2(b []byte) {
	m := NewMap(b)
	acc := 0

	for y := range m.height {
		for x := range m.width {
			if m.TestXMas(x, y) {
				acc++
			}
		}
	}
	fmt.Printf("Task 2: %d\n", acc)
}

type Point struct {
	x, y int
	b    byte
}

type Map struct {
	width  int
	height int
	points [][]Point
}

func NewMap(input []byte) *Map {
	width := bytes.IndexByte(input, '\n') + 1
	height := bytes.LastIndexByte(input, '\n')/width + 1

	points := make([][]Point, height)
	for i := range points {
		points[i] = make([]Point, width)
	}

	for i, b := range []byte(input) {
		if b == '\n' {
			continue
		}

		points[i/width][i%width] = Point{
			x: i % width,
			y: i / width,
			b: b,
		}
	}

	return &Map{
		width:  width,
		height: height,
		points: points,
	}
}

var XMAS = []byte{X, M, A, S}

func getNextByte(b byte) byte {
	if i := bytes.IndexByte(XMAS, b); i >= 0 && i < 3 {
		return XMAS[i+1]
	} else {
		return 0
	}
}

func (m *Map) TestR(p *Point, offsetX, offsetY int, b byte) bool {
	if b == 0 {
		return true
	}
	if p == nil {
		return false
	}
	if p.b == b {
		return m.TestR(m.Get(p.x+offsetX, p.y+offsetY), offsetX, offsetY, getNextByte(b))
	}
	return false
}

func (m *Map) Get(x, y int) *Point {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return nil
	}
	return &m.points[y][x]
}

type Corners struct {
	topLeft, topRight, bottomLeft, bottomRight Point
}

func (m *Map) GetCorners(x, y int) (Corners, bool) {
	if x < 1 || x >= m.width-1 || y < 1 || y >= m.height-1 {
		return Corners{}, false
	}

	return Corners{
		topLeft:     *m.Get(x-1, y-1),
		topRight:    *m.Get(x+1, y-1),
		bottomLeft:  *m.Get(x-1, y+1),
		bottomRight: *m.Get(x+1, y+1),
	}, true
}

func (m *Map) TestXMas(x, y int) bool {
	if p := m.Get(x, y); p == nil || p.b != A {
		return false
	}

	//|M S|S M|M M|S S|
	//| A | A | A | A |
	//|M S|S M|S S|M M|

	if c, ok := m.GetCorners(x, y); ok {
		return (c.topLeft.b == M && c.topRight.b == S && c.bottomLeft.b == M && c.bottomRight.b == S) ||
			(c.topLeft.b == S && c.topRight.b == M && c.bottomLeft.b == S && c.bottomRight.b == M) ||
			(c.topLeft.b == M && c.topRight.b == M && c.bottomLeft.b == S && c.bottomRight.b == S) ||
			(c.topLeft.b == S && c.topRight.b == S && c.bottomLeft.b == M && c.bottomRight.b == M)

	}
	return false
}
