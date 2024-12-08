package aoc

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)
	
func filterEmpty(s []string) []string {
	var ret []string
	for _, val := range s {
		if len(val) != 0 {
			ret = append(ret, val)
		}
	}
	return ret
}

func parseInput(s string) ([]int, []int, error) {
	var fst []int
	var snd []int

	s_split := filterEmpty(strings.Split(s, "\n"))
	for i, val := range s_split {
		val_split := filterEmpty(strings.Split(val, " "))
		if len(val_split) != 2 {
			err_str := fmt.Sprintf("Cannot split line: %v", i)
			return nil, nil, errors.New(err_str)
		}
		fst_val, err := strconv.Atoi(val_split[0])
		if err != nil {
			return nil, nil, err
		}
		snd_val, err := strconv.Atoi(val_split[1])
		if err != nil {
			return nil, nil, err
		}
		fst = append(fst, fst_val)
		snd = append(snd, snd_val)
	}
	if len(fst) != len(snd) {
		err_str := fmt.Sprintf("Length of fst and snd are different:\nlen(fst) = %v,\nlen(snd) = %v", len(fst), len(snd))
		return nil, nil, errors.New(err_str)
	}
	return fst, snd, nil
}

func PartOne(s string) (int,error) {
	fst, snd, err := parseInput(s)

	if err != nil {
		return 0, err
	}

	sort.Ints(fst)
	sort.Ints(snd)

	ret := 0

	for i := 0; i < len(fst); i++ {
		tmp := fst[i] - snd[i]
		if tmp < 0 {
			tmp = -tmp
		}
		ret += tmp
	}

	return ret, nil
}
func PartTwo(s string) (int,error) {
	dict := make(map[int]int)
	fst, snd, err := parseInput(s)

	if err != nil {
		return 0, err
	}

	for _, key := range snd {
		dict[key]++
	}

	var ret int

	for _, key := range fst {
		ret += key * dict[key]
	}

  return ret, nil
}
