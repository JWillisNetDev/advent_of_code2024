package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		Error("Problem reading input: %s", err)
	}
	_ = input
	Part1(input)
	Part2(input)
}

func Error(s string, e ...any) {
	fmt.Fprintf(os.Stderr, s, e...)
	os.Exit(1)
}

func Part1(input []byte) {
	m := NewMap(input)
	fmt.Printf("Map of size %d, %d\n", m.w, m.h)

	antinodes := map[Vec2]struct{}{}
	acc := 0

	for i, b := range m.input {
		if b == '.' || b == 0 {
			continue
		}

		n := i
		for {
			next, nn := m.GetNextAntenna(b, n)
			if nn < 0 {
				break
			}
			v, ok := m.ToCoord(i)
			if !ok {
				Error("Attempted to create an invalidate coordinate from pos %d in input", i)
				break
			}

			// Get offset from origin antenna to next antenna
			offset := GetOffset(v, next)

			// Test top-left quadrant
			q1 := v.Translate(offset)
			if m.IsValidForAntinode(q1) {
				if _, ok := antinodes[q1]; !ok {
					acc++
					antinodes[q1] = struct{}{}
				}
			}

			// Test bot-right quadrant
			q4 := next.Translate(offset.Mult(-1))
			if m.IsValidForAntinode(q4) {
				if _, ok := antinodes[q4]; !ok {
					acc++
					antinodes[q4] = struct{}{}
				}
			}

			n = nn
		}
	}

	fmt.Printf("Part 1: %d", acc)
}

func Part2(input []byte) {
	m := NewMap(input)
	fmt.Printf("Map of size %d, %d\n", m.w, m.h)

	antinodes := map[Vec2]struct{}{}
	acc := 0

	for i, b := range m.input {
		if b == '.' || b == 0 {
			continue
		}

		n := i
		for {
			next, nn := m.GetNextAntenna(b, n)
			if nn < 0 {
				break
			}
			v, ok := m.ToCoord(i)
			if !ok {
				Error("Attempted to create an invalidate coordinate from pos %d in input", i)
				break
			}

			if _, ok := antinodes[v]; !ok {
				acc++
				antinodes[v] = struct{}{}
			}

			if _, ok := antinodes[next]; !ok {
				acc++
				antinodes[next] = struct{}{}
			}
			// Get offset from origin antenna to next antenna
			offset := GetOffset(v, next)

			// Test top-left quadrant
			acc += m.TestOffset(b, v, offset, antinodes)

			// Test bot-right quadrant
			acc += m.TestOffset(b, next, offset.Mult(-1), antinodes)

			n = nn
		}
	}

	for y := 0; y < m.h; y++ {
		for x := 0; x < m.w; x++ {
			v := Vec2{x, y}
			if _, ok := antinodes[v]; !ok {
				b, _ := m.GetByte(v)
				fmt.Printf("%c", b)
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}

	fmt.Printf("Part 2: %d", acc)
}

type Vec2 struct {
	x, y int
}

func (v Vec2) String() string         { return fmt.Sprintf("(%d, %d)", v.x, v.y) }
func (v Vec2) Translate(vv Vec2) Vec2 { return Vec2{v.x + vv.x, v.y + vv.y} }
func (v Vec2) Mult(i int) Vec2        { return Vec2{v.x * i, v.y * i} }

type Antenna struct {
	pos Vec2
}

type Map struct {
	input []byte
	w, h  int
}

func NewMap(input []byte) *Map {
	m := &Map{
		input: bytes.ReplaceAll(input, []byte("\n"), []byte{}),
		w:     bytes.Index(input, []byte("\n")),
		h:     bytes.Count(input, []byte("\n")),
	}
	return m
}

func (m *Map) ToCoord(i int) (Vec2, bool) {
	if i >= m.w*m.h || i < 0 {
		return Vec2{}, false
	}
	return Vec2{i % m.w, i / m.w}, true
}

func (m *Map) FromCoord(v Vec2) int {
	if v.x >= m.w || v.x < 0 || v.y >= m.h || v.y < 0 {
		return -1
	}
	return v.y*m.w + v.x%m.w
}

func (m *Map) GetByte(v Vec2) (byte, bool) {
	if i := m.FromCoord(v); i >= 0 {
		return m.input[i], true
	}
	return 0, false
}

func GetOffset(a, b Vec2) Vec2 {
	return Vec2{a.x - b.x, a.y - b.y}
}

func (m *Map) GetNextAntenna(b byte, n int) (Vec2, int) {
	if n < 0 || n >= len(m.input) {
		return Vec2{}, -1
	}

	if i := bytes.Index(m.input[n+1:], []byte{b}); i >= 0 {
		coord, _ := m.ToCoord(i + n + 1)
		return coord, i + n + 1
	}
	return Vec2{}, -1
}

type Antinode struct {
	v Vec2
	b byte
}

func (m *Map) IsValidForAntinode(v Vec2) bool {
	if _, ok := m.GetByte(v); ok {
		return true
	}
	return false
}

func (m *Map) TestOffset(b byte, origin, offset Vec2, antinodes map[Vec2]struct{}) int {
	o := origin.Translate(offset)
	acc := 0
	for {
		if m.IsValidForAntinode(o) {
			if _, ok := antinodes[o]; !ok {
				acc++
				antinodes[o] = struct{}{}
			}
			o = o.Translate(offset)
		} else {
			break
		}
	}
	return acc
}
