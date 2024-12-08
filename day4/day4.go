package aoc

import (
	"aoc/lib"
	"errors"
	"fmt"
	"strings"
)

func Verticals(ss []string) ([]string, error) {
	rows := len(ss)
	if rows == 0 {
		err := errors.New("the input has 0 length")
		return []string{}, err
	}
	cols := len(ss[0])
	var ret []string

	for i := 0;i < cols; i++ {
		var tmp strings.Builder
		for j := 0; j < rows; j++ {
			tmp.WriteByte(ss[j][i])
		}
		ret = append(ret, tmp.String())
	}
	return ret, nil
}

func Diagonals(ss []string) ([]string, error) {
	rows := len(ss)
	if rows == 0 {
		err := errors.New("the input has 0 length")
		return []string{}, err
	}
	cols := len(ss[0])
	var ret []string

	for i := 0; i < rows-1; i++ {
		var tmp strings.Builder
		for j := 0; j <= min(i, cols); j++ {
			err := tmp.WriteByte(ss[i-j][j])
			if err != nil {
				return []string{}, err
			}
		}
		ret = append(ret, tmp.String())
	}

	for i := 0; i < cols; i++ {
		var tmp strings.Builder
		for j := 0; j < cols-i; j++ {
			err := tmp.WriteByte(ss[rows-j-1][i+j])
			if err != nil {
				return []string{}, err
			}
		}
		ret = append(ret, tmp.String())
	}

	return ret, nil
}

func XMAS(s []string) (int, error) {
	if len(s) < 3 {
		return 0, errors.New("slice is not long enough")
	}

	if len(s[0]) < 3 {
		return 0, errors.New("string length is not long enough")
	}

	rows := len(s) - 2
	cols := len(s[0]) - 2
	offset := []int{0, 2, 2, 0, 0}
	ret := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var str strings.Builder
			if s[i+1][j+1] != 'A' {
				continue
			}
			for k := 0; k < 4; k++ {
				err := str.WriteByte(s[i + offset[k]][j + offset[k + 1]])

				if err != nil {
					return 0, err
				}
			}
			tmp := fmt.Sprintf("%v%v", str.String(), str.String())
			if strings.Contains(tmp, "MM") && strings.Contains(tmp, "SS") {
				ret++
			}
		}
	}
	
	return ret, nil
}

func PartOne(s string) (int,error) {
	grid := lib.SplitFilterEmpty(s, "\n")
	diag, err := Diagonals(grid)
	if err != nil {
		return 0, err
	}

	diagRev, err := Diagonals(lib.ReverseSlice(grid))
	if err != nil {
		return 0, err
	}

	vert, err := Verticals(grid)
	if err != nil {
		return 0, err
	}


	ret := 0

	for _, val := range grid {
		ret += strings.Count(val, "XMAS")
		ret += strings.Count(val, "SAMX")
	}

	for _, val := range diag {
		ret += strings.Count(val, "XMAS")
		ret += strings.Count(val, "SAMX")
	}

	for _, val := range diagRev {
		ret += strings.Count(val, "XMAS")
		ret += strings.Count(val, "SAMX")
	}

	for _, val := range vert {
		ret += strings.Count(val, "XMAS")
		ret += strings.Count(val, "SAMX")
	}

	return ret, nil
}

func PartTwo(s string) (int,error) {
	grid := lib.SplitFilterEmpty(s, "\n")
  return XMAS(grid)
}
