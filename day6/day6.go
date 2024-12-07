package day

import (
	"aoc/lib"
	"errors"
	"fmt"
	"slices"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Coordinate struct {
	x, y int
}

func PrintVis(grid []string, vis map[Coordinate]Direction) {
	for i, v := range grid {
		for j, c := range v {
			if _, ok := vis[Coordinate{i, j}]; ok {
				fmt.Printf("X")
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Println()
	}
}

func ValidCoordinate(coord Coordinate, max_row int, max_col int) bool {
	return coord.x >= 0 && coord.x < max_row && coord.y >= 0 && coord.y < max_col
}

func Simulate(grid []string, coord Coordinate, dir Direction, vis *map[Coordinate]Direction) (bool, error) {
	if len(grid) == 0 {
		return false, errors.New("0 row received")
	}
	if len(grid[0]) == 0 {
		return false, errors.New("0 column received")
	}

	dir_cand, ok := (*vis)[coord]
	if ok {
		if dir_cand == dir {
			return true, nil
		}
	}

	offset := []int{-1, 0, 1, 0, -1}
	next_candidate := Coordinate{coord.x + offset[dir], coord.y + offset[dir+1]}
	if ValidCoordinate(next_candidate, len(grid), len(grid[0])) {
		var res bool
		if grid[next_candidate.x][next_candidate.y] == '#' {
			next_dir := (dir+1)%4
			tmp, err := Simulate(grid, coord, next_dir, vis)
			if err != nil {
				return false, err
			}
			res = tmp
		} else {
			(*vis)[coord] = dir
			tmp, err := Simulate(grid, next_candidate, dir, vis)
			if err != nil {
				return false, err
			}
			res = tmp
		}
		return res, nil
	}
	(*vis)[coord] = dir
	return false, nil
}

func FindStart(grid []string) (Coordinate, Direction, error) {
	dir := []rune{'^', '>', 'v', '<'}
	for x, v := range grid {
		for y, c := range v {
			if slices.Contains(dir, c) {
				return Coordinate{x, y}, Direction(slices.Index(dir, c)), nil
			}
		}
	}
	return Coordinate{}, Up, errors.New("Grid does not have start")
}

func PartOne(s string) (int,error) {
	grid := lib.SplitFilterEmpty(s, "\n")
	vis := make(map[Coordinate]Direction)
	start, dir, err := FindStart(grid)
	if err != nil {
		return 0, err
	}

	_, err = Simulate(grid, start, dir, &vis)
	if err != nil {
		return 0, err
	}
	ret := 0
	for range vis {
		ret++
	}
	//PrintVis(grid, vis)
	return ret, nil
}

func AlterString(s string, i int) (bool, string) {
	r := []rune(s)
	dir := []rune{'^', '>', 'v', '<'}
	if r[i] != '\n' && !slices.Contains(dir, r[i]){
		r[i] = '#'
	} else {
		return true, ""
	}
	return false, string(r)
}

func PartTwo(s string) (int,error) {
	ret := 0
	fmt.Println("!!!WARNING!!!\nThis solution is O(n^4) and possibly has uncovered (yet maybe rare) edge cases.")
	for i := range s {
		cont, new_s := AlterString(s, i)
		if cont {
			continue
		}
		grid := lib.SplitFilterEmpty(new_s, "\n")
		vis := make(map[Coordinate]Direction)
		start, dir, err := FindStart(grid)
		if err != nil {
			return 0, err
		}

		res, err := Simulate(grid, start, dir, &vis)
		if err != nil {
			return 0, err
		}
		if res {
			ret++
		}
	}
	return ret, nil
}
