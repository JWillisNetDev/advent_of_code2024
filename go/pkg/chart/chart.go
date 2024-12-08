package chart

import (
	"bufio"
	"bytes"
	"fmt"
)

type Chart struct {
	input []byte
	w, h  int
}

func NewChart(b []byte) (*Chart, error) {
	c := &Chart{
		input: bytes.ReplaceAll(b, []byte{'\n'}, []byte{}),
		h:     1,
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(b))
	scanner.Scan()
	if bb := scanner.Bytes(); len(bb) > 0 {
		c.w = len(bb)
	}

	for scanner.Scan() {
		bb := scanner.Bytes()
		l := len(bb)
		if l == 0 {
			break
		}
		c.h++

		if l != c.w {
			return nil, fmt.Errorf("ill formatted chart encountered around line: %d - width of %d does not match first line width of %d", c.h, l, c.w)
		}
	}

	n := bytes.LastIndexByte(b, '\n')
	if n > 0 && len(b) == n {
		c.h++
	}

	return c, nil
}

func (c *Chart) Size() int   { return c.w * c.h }
func (c *Chart) Width() int  { return c.w }
func (c *Chart) Height() int { return c.h }

func (c *Chart) GetIndex(i int) (byte, bool) {
	if i < 0 || i >= len(c.input) {
		return 0, false
	}
	return c.input[i], true
}

type Coord struct {
	X, Y int
}

func (c Coord) String() string { return fmt.Sprintf("(%d, %d)", c.X, c.Y) }

func (c *Chart) normalizeCoord(co Coord) int {
	return co.Y*c.w + co.X%c.w
}

func (c *Chart) denormalize(i int) Coord {
	return Coord{i % c.w, i / c.h}
}

func (c *Chart) IsWithinBounds(co Coord) bool {
	return co.X < 0 || co.X >= c.w || co.Y < 0 || co.Y >= c.h
}

func (c *Chart) Get(co Coord) (byte, bool) {
	if !c.IsWithinBounds(co) {
		return 0, false
	}
	return c.input[c.normalizeCoord(co)], true
}

func (c *Chart) First(b byte) (Coord, bool) {
	for i, bb := range c.input {
		if b == bb {
			return c.denormalize(i), true
		}
	}
	return Coord{}, false
}

func (c *Chart) FirstN(b byte, n int) (Coord, int) {
	input := c.input
	if n > 0 {
		input = input[n:]
	}
	i := bytes.Index(input, []byte{b})
	if i >= 0 {
		return c.denormalize(i + n), i + n
	}
	return Coord{}, -1
}

func (c *Chart) Where(b byte) []Coord {
	co := []Coord{}
	i := 0
	for {
		if ii := bytes.IndexByte(c.input[i:], b); ii > 0 {
			co = append(co, c.denormalize(i+ii))
			i += ii + 1
		} else {
			break
		}
	}
	return co
}
