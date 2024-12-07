package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from input: %s", err)
		os.Exit(1)
	}
	_ = input
	Part1(input)
	Part2(input)
}

type Unit struct{}

func Part1(input []byte) {
	m := NewMap(input)
	s := m.MapState(DirUp)
	visited := map[Point]Unit{s.guard: {}}
	acc := 1
	for {
		ss, isEnd := s.GetNextState(m)
		if isEnd {
			break
		}

		if _, ok := visited[ss.guard]; !ok {
			visited[ss.guard] = Unit{}
			acc++
		}

		s = ss
	}

	fmt.Printf("Part 1: %d\n", acc)
}

func Part2(input []byte) {
	m := NewMap(input)
	ms := m.MapState(DirUp)
	acc := 0
	for _, p := range ms.GetGuardPath(m) {
		if p == ms.guard {
			continue
		}

		mm := NewMap(input)
		mm.SetByte(p, OBSTACLE)
		ms := mm.MapState(DirUp)
		states := map[MapState]struct{}{ms: {}}
		cur := ms

		for {
			next, isEnd := cur.GetNextState(mm)
			if isEnd {
				break
			}
			if _, ok := states[next]; ok {
				acc++
				break
			}
			cur = next
			states[cur] = struct{}{}
		}
	}

	fmt.Printf("Part 2: %d\n", acc)
}

func Error(s string, e ...any) {
	fmt.Fprintf(os.Stderr, s, e...)
	os.Exit(1)
}

type Map struct {
	width, height int
	input         []byte
}

func NewMap(input []byte) *Map {
	return &Map{
		width:  bytes.Index(input, []byte("\n")),
		height: bytes.Count(input, []byte("\n")),
		input:  bytes.ReplaceAll(input, []byte("\n"), []byte{}),
	}
}

func (m *Map) String() string {
	var sb strings.Builder
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			sb.WriteByte(m.input[y*m.width+x])
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

type Point struct {
	x, y int
}

func (p Point) String() string           { return fmt.Sprintf("(X: %d, Y: %d)", p.x, p.y) }
func (p Point) Translate(pp Point) Point { return Point{x: p.x + pp.x, y: p.y + pp.y} }

func (m *Map) Points() []Point {
	p := []Point{}
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			p = append(p, Point{x: x, y: y})
		}
	}
	return p
}

func (m *Map) IsOutOfBounds(p Point) bool {
	return p.x < 0 || p.x >= m.width || p.y < 0 || p.y >= m.height
}

func (m *Map) GetByte(p Point) byte {
	if p.x < 0 || p.x >= m.width {
		Error("Value x was out of bounds for map: %d\n", p.x)
	}

	if p.y < 0 || p.y >= m.height {
		Error("Value x was out of bounds for map: %d\n", p.x)
	}

	return m.input[p.y*m.width+p.x]
}

func (m *Map) SetByte(p Point, b byte) {
	if p.x < 0 || p.x >= m.width {
		Error("Value x was out of bounds for map: %d\n", p.x)
	}

	if p.y < 0 || p.y >= m.height {
		Error("Value x was out of bounds for map: %d\n", p.x)
	}

	m.input[p.y*m.width+p.x] = b
}

const (
	GUARD    = '^'
	EMPTY    = '.'
	OBSTACLE = '#'
)

func (m *Map) GetGuardPoint() Point {
	for _, p := range m.Points() {
		if m.GetByte(p) == GUARD {
			return p
		}
	}

	Error("A guard point could not be located")
	return Point{}
}

var (
	DirUp    = Point{y: -1}
	DirRight = Point{x: 1}
	DirDown  = Point{y: 1}
	DirLeft  = Point{x: -1}
)

type MapState struct {
	guard Point
	dir   Point
}

func (m *Map) MapState(dir Point) MapState {
	return MapState{
		guard: m.GetGuardPoint(),
		dir:   dir,
	}
}

func GetNextDir(dir Point) Point {
	switch dir {
	case DirUp:
		return DirRight
	case DirRight:
		return DirDown
	case DirDown:
		return DirLeft
	case DirLeft:
		return DirUp
	default:
		Error("Unexpected value for next dir: %s", dir)
		return Point{}
	}
}

func (ms MapState) GetNextState(m *Map) (MapState, bool) {
	t := ms.guard.Translate(ms.dir)

	if m.IsOutOfBounds(t) {
		return MapState{}, true
	}

	switch m.GetByte(t) {
	case EMPTY, GUARD:
		ms.guard = t
		return ms, false
	case OBSTACLE:
		ms.dir = GetNextDir(ms.dir)
		return ms, false
		// return ms.GetNextState(m)
	default:
		Error("Unanticipated byte encountered: %c at %s", m.GetByte(t), t)
		return MapState{}, false
	}
}

func (ms MapState) GetGuardPath(m *Map) []Point {
	points := map[Point]struct{}{}

	for {
		next, end := ms.GetNextState(m)
		if end {
			break
		}
		points[next.guard] = struct{}{}
		ms = next
	}

	parr := []Point{}
	for p := range points {
		parr = append(parr, p)
	}
	return parr
}
