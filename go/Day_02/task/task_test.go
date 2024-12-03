package task

import "testing"

func TestSample1(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	actual := Task1(input)
	if actual != 2 {
		t.Errorf("Expected 2, but got %d", actual)
	}
}

func TestSample2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	actual := Task2(input)
	if actual != 4 {
		t.Errorf("Expected 4, but got %d", actual)
	}
}