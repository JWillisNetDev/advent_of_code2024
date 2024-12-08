package chart

import (
	"testing"
)

func testChartDimensions(t *testing.T, c *Chart, w, h, sz int) {
	t.Helper()
	if c.Width() != w {
		t.Errorf("Width expected %d, got %d", w, c.Width())
	}
	if c.Height() != h {
		t.Errorf("Height expected %d, got %d", h, c.Height())
	}
	if c.Size() != sz {
		t.Errorf("Size expected %d, got %d", h, c.Size())
	}
}

func TestNewChart(t *testing.T) {
	c, _ := NewChart([]byte(`...
...
...`))

	testChartDimensions(t, c, 3, 3, 9)

	c, _ = NewChart([]byte(`...
...
...
`))

	testChartDimensions(t, c, 3, 3, 9)
}

func TestWhere(t *testing.T) {
	c, _ := NewChart([]byte(`abc
abc
abc`))

	tests := []struct {
		b        byte
		expected []Coord
	}{
		{'a', []Coord{{0, 0}, {0, 1}, {0, 2}}},
		{'b', []Coord{{1, 0}, {1, 1}, {1, 2}}},
		{'c', []Coord{{2, 0}, {2, 1}, {2, 2}}},
	}
	for _, tt := range tests {
		for i, actual := range c.Where(tt.b) {
			if actual != tt.expected[i] {
				t.Errorf("Coordinate expected %s, got %s", actual, tt.expected[i])
			}
		}
	}
}
