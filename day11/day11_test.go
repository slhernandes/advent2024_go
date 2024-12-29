package aoc

import (
	"fmt"
	"os"
	"testing"
)

func TestSmallPartOne(t *testing.T) {
	input, err := os.ReadFile("input_small")
	if err != nil {
		t.Error("Input file not found.")
	}
	ans, err := PartOne(string(input))
	if err != nil {
		t.Error(err)
	}
	exp := 55312
	if ans != exp {
		t.Errorf("Part 1 answer wrong, expected: %v, got: %v", exp, ans)
	}
}

func TestSmallPartTwo(t *testing.T) {
	input, err := os.ReadFile("input_small")
	if err != nil {
		t.Error("Input file not found.")
	}
	ans, err := PartTwo(string(input))
	if err != nil {
		t.Error(err)
	}
	if ans <= 0 {
		t.Error("Part 2 answer wrong.")
	}
}

func TestPartOne(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Error("Input file not found.")
	}
	ans, err := PartOne(string(input))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Part One: %v\n", ans)
}

func TestPartTwo(t *testing.T) {
	input, err := os.ReadFile("input")
	if err != nil {
		t.Error("Input file not found.")
	}
	ans, err := PartTwo(string(input))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Part Two: %v\n", ans)
}