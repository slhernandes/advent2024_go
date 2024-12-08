package day

import (
	"aoc/lib"
	"fmt"
)

type Coordinate struct {
	x, y int
}

func Diff(a Coordinate, b Coordinate) Coordinate {
	return Coordinate{a.x - b.x, a.y - b.y}
}

func Add(a Coordinate, b Coordinate) Coordinate {
	return Coordinate{a.x + b.x, a.y + b.y}
}

func PrintGrid(grid []string, anti map[Coordinate]bool) {
	for x, v := range grid {
		for y, c := range v {
			_, ok := anti[Coordinate{x, y}]
			if ok {
				fmt.Printf("%c", '#')
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Println()
	}
}

func PartOne(s string) (int,error) {
	grid := lib.SplitFilterEmpty(s, "\n")
	atns := make(map[rune][]Coordinate)
	anti := make(map[Coordinate]bool)
	FindAntennae(grid, &atns)
	FindAntiNodes(grid, atns, &anti)
	return len(anti), nil
}

func FindAntiNodes(grid []string, atns map[rune][]Coordinate, anti *map[Coordinate]bool) {
	for _, v := range atns {
		for i := 0; i < len(v)-1; i++ {
			for j := i+1; j < len(v); j++ {
				diff_1 := Diff(v[i], v[j])
				anti_1 := Add(v[i], diff_1)
				diff_2 := Diff(v[j], v[i])
				anti_2 := Add(v[j], diff_2)
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

func Contained(a Coordinate, x int, y int) bool {
	return  a.x >= 0 && a.y >= 0 && a.x < x && a.y < y;
}

func FindAntennae(grid []string, atns *map[rune][]Coordinate) {
	for x, v := range grid {
		for y, c := range v {
			if c != '.' {
				(*atns)[c] = append((*atns)[c], Coordinate{x, y})
			}
		}
	}
}

func PartTwo(s string) (int,error) {
	grid := lib.SplitFilterEmpty(s, "\n")
	atns := make(map[rune][]Coordinate)
	anti := make(map[Coordinate]bool)
	FindAntennae(grid, &atns)
	FindAntiNodesAdv(grid, atns, &anti)
  return len(anti), nil
}

func FindAntiNodesAdv(grid []string, atns map[rune][]Coordinate, anti *map[Coordinate]bool) {
	for _, v := range atns {
		for i := 0; i < len(v)-1; i++ {
			for j := i+1; j < len(v); j++ {
				(*anti)[v[i]] = true
				delta := Diff(v[i], v[j])
				now := v[i]
				for  ; Contained(Add(now, delta), len(grid), len(grid[0])); now = Add(now, delta) {
					(*anti)[now] = true
				}
				(*anti)[now] = true
				now = v[i]
				for  ; Contained(Diff(now, delta), len(grid), len(grid[0])); now = Diff(now, delta) {
					(*anti)[now] = true
				}
				(*anti)[now] = true
			}
		}
	}
}
