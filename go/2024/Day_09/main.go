package main

import (
	"fmt"
	"os"
	"github.com/jwillisnetdev/advent_of_code2024/go/pkg/input"
)

func main() {
	i, err := os.ReadFile("input.txt")
	if err != nil {
		input.Error("Error reading input file: %s", err)
	}
	_ = i

	Part1(i)
	Part1([]byte(`2333133121414131402`))
}

func Part1(input []byte) {
	input = sanitizeInput(input)
	blocks := ParseBlocks(input)
	compacted := Compact(blocks)
	sum := GetHashSum(compacted)
	fmt.Printf("Part 1: %d\n", sum)
}

func sanitizeInput(input []byte) []byte {
	if input[len(input)-1] == '\n' {
		return input[:len(input)-1]
	}
	return input
}

type Block struct {
	id, blocks, free int
}

func (b Block) String() string { return fmt.Sprintf("Block %d: %d blocks, %d free", b.id, b.blocks, b.free) }

func ParseBlocks(input []byte) []Block {
	id := 0
	blocks := make([]Block, len(input)/2+1)
	for i := 0; i < len(input) - 1; i += 2 {
		b := Block{
			id: id,
			blocks: btoi(input[i]),
			free: btoi(input[i+1]),
		}
		blocks[id] = b
		id++
	}
	blocks[id] = Block{id: id, blocks: btoi(input[len(input)-1]), free: 0}
	return blocks
}

// Returns a list of IDs repeated by their size
func Compact(blocks []Block) ([]int) {
	compacted := []int{}
	startPtr, endPtr := 0, len(blocks)-1
	start, end := blocks[startPtr], blocks[endPtr]
	endLen := 0
	for {
		for range start.blocks {
			compacted = append(compacted, start.id)
		}

		for range start.free {
			if endLen == end.blocks {
				endPtr--
				end = blocks[endPtr]
				endLen = 0
			}
			compacted = append(compacted, end.id)
			endLen++
		}

		startPtr++
		if startPtr == endPtr {
			break
		}
		start = blocks[startPtr]
	}

	// Add the remaining of the last block
	if end.blocks - endLen > 0 {
		for range end.blocks - endLen {
			compacted = append(compacted, end.id)
		}
	}

	return compacted
}

func GetHashSum(input []int) int64 {
	acc := int64(0)
	for i, ii := range input {
		acc += int64(i) * int64(ii)
	}
	return acc
}

func btoi(b byte) int {	return int(b - '0') }