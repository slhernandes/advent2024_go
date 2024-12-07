package day

import (
	"aoc/lib"
	"errors"
	"fmt"
	"slices"
)

type Direction int

const (
	Up Direction = 1<<iota
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

func DirToIdx(dir Direction) int {
	ret := 0
	cur := 1
	for ; cur & int(dir) == 0; cur = cur<<1 {
		ret++
	}
	return ret
}

func Simulate(grid []string, coord Coordinate, dir Direction, vis *map[Coordinate]Direction, uturn int) (bool, error) {
	if len(grid) == 0 {
		return false, errors.New("0 row received")
	}
	if len(grid[0]) == 0 {
		return false, errors.New("0 column received")
	}

	if uturn > 1 {
		return true, nil
	}

	dir_cand, ok := (*vis)[coord]
	if ok {
		if dir_cand & dir != 0 {
			return true, nil
		}
	}

	offset := []int{-1, 0, 1, 0, -1}
	dir_idx := DirToIdx(dir)
	next_candidate := Coordinate{coord.x + offset[dir_idx], coord.y + offset[dir_idx+1]}
	if ValidCoordinate(next_candidate, len(grid), len(grid[0])) {
		var res bool
		if grid[next_candidate.x][next_candidate.y] == '#' {
			next_dir := (dir_idx+1)%4
			tmp, err := Simulate(grid, coord, Direction(1<<next_dir), vis, uturn)
			if err != nil {
				return false, err
			}
			res = tmp
		} else {
			(*vis)[coord] |= dir
			uturn_now := uturn
			if UTurned((*vis)[coord]) {
				uturn_now++
			}
			tmp, err := Simulate(grid, next_candidate, dir, vis, uturn_now)
			if err != nil {
				return false, err
			}
			res = tmp
		}
		return res, nil
	}
	(*vis)[coord] |= dir
	return false, nil
}

func FindStart(grid []string) (Coordinate, Direction, error) {
	dir := []rune{'^', '>', 'v', '<'}
	for x, v := range grid {
		for y, c := range v {
			if slices.Contains(dir, c) {
				return Coordinate{x, y}, Direction(1<<slices.Index(dir, c)), nil
			}
		}
	}
	return Coordinate{}, Up, errors.New("Grid does not have start")
}

func UTurned(dir Direction) bool {
	for i := 0; i < 4; i++ {
		if int(dir) & (1<<i) == 15 {
			return true
		}
	}
	return false
}

func PartOne(s string) (int,error) {
	grid := lib.SplitFilterEmpty(s, "\n")
	vis := make(map[Coordinate]Direction)
	start, dir, err := FindStart(grid)
	if err != nil {
		return 0, err
	}

	_, err = Simulate(grid, start, dir, &vis, 0)
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
	if r[i] != '\n' && !slices.Contains(dir, r[i]) && r[i] != '#' {
		r[i] = '#'
	} else {
		return true, ""
	}
	return false, string(r)
}

func PartTwo(s string) (int,error) {

	grid := lib.SplitFilterEmpty(s, "\n")
	width := len(grid[0])
	vis := make(map[Coordinate]Direction)
	start, dir, err := FindStart(grid)
	if err != nil {
		return 0, err
	}

	_, err = Simulate(grid, start, dir, &vis, 0)
	if err != nil {
		return 0, err
	}

	ret := 0
	for i := range vis {
		cont, new_s := AlterString(s, i.y + i.x * (width+1))
		if cont {
			continue
		}
		grid := lib.SplitFilterEmpty(new_s, "\n")
		vis := make(map[Coordinate]Direction)
		start, dir, err := FindStart(grid)
		if err != nil {
			return 0, err
		}

		res, err := Simulate(grid, start, dir, &vis, 0)
		if err != nil {
			return 0, err
		}
		if res {
			ret++
		}
	}
	return ret, nil
}
