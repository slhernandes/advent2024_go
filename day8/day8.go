package aoc

import (
	"aoc/lib"
	"fmt"
)

func PrintGrid(grid []string, anti map[lib.Coordinate]bool) {
	for x, v := range grid {
		for y, c := range v {
			_, ok := anti[lib.Coordinate{X: x, Y: y}]
			if ok {
				fmt.Printf("%c", '#')
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Println()
	}
}

func PartOne(s string) (int, error) {
	grid := lib.SplitFilterEmpty(s, "\n")
	atns := make(map[rune][]lib.Coordinate)
	anti := make(map[lib.Coordinate]bool)
	FindAntennae(grid, &atns)
	FindAntiNodes(grid, atns, &anti)
	return len(anti), nil
}

func FindAntiNodes(grid []string, atns map[rune][]lib.Coordinate, anti *map[lib.Coordinate]bool) {
	for _, v := range atns {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				diff_1 := lib.Diff(v[i], v[j])
				anti_1 := lib.Add(v[i], diff_1)
				diff_2 := lib.Diff(v[j], v[i])
				anti_2 := lib.Add(v[j], diff_2)
				if Contained(anti_1, len(grid), len(grid[0])) {
					(*anti)[anti_1] = true
				}
				if Contained(anti_2, len(grid), len(grid[0])) {
					(*anti)[anti_2] = true
				}
			}
		}
	}
}

func Contained(a lib.Coordinate, x int, y int) bool {
	return a.X >= 0 && a.Y >= 0 && a.X < x && a.Y < y
}

func FindAntennae(grid []string, atns *map[rune][]lib.Coordinate) {
	for x, v := range grid {
		for y, c := range v {
			if c != '.' {
				(*atns)[c] = append((*atns)[c], lib.Coordinate{X: x, Y: y})
			}
		}
	}
}

func PartTwo(s string) (int, error) {
	grid := lib.SplitFilterEmpty(s, "\n")
	atns := make(map[rune][]lib.Coordinate)
	anti := make(map[lib.Coordinate]bool)
	FindAntennae(grid, &atns)
	FindAntiNodesAdv(grid, atns, &anti)
	return len(anti), nil
}

func FindAntiNodesAdv(grid []string, atns map[rune][]lib.Coordinate, anti *map[lib.Coordinate]bool) {
	for _, v := range atns {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				(*anti)[v[i]] = true
				delta := lib.Diff(v[i], v[j])
				now := v[i]
				for ; Contained(lib.Add(now, delta), len(grid), len(grid[0])); now = lib.Add(now, delta) {
					(*anti)[now] = true
				}
				(*anti)[now] = true
				now = v[i]
				for ; Contained(lib.Diff(now, delta), len(grid), len(grid[0])); now = lib.Diff(now, delta) {
					(*anti)[now] = true
				}
				(*anti)[now] = true
			}
		}
	}
}
