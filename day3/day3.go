package aoc

import (
	"regexp"
	"strconv"
	"strings"
)

func SplitFilterEmpty(s string, sep string) []string {
	tmp := strings.Split(s, sep)
	var ret []string
	for _, val := range tmp {
		if len(val) != 0 {
			ret = append(ret, val)
		}
	}
	return ret
}

func CalculateInput(s string) (int, error) {
	tmp := SplitFilterEmpty(s, "do()")
	ret := 0

	re_dnt, err := regexp.Compile(`don't\(\)`)
	if err != nil {
		return 0, err
	}

	re, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	if err != nil {
		return 0, err
	}

	for _, val := range tmp {
		to_calc := re_dnt.Split(val, 2)[0]
		found := re.FindAllStringSubmatch(to_calc, -1)
		for _, nums := range found {
			fst, _ := strconv.Atoi(nums[1])
			snd, _ := strconv.Atoi(nums[2])
			ret += fst * snd
		}
	}
	return ret, nil
}

func PartOne(s string) (int, error) {
	re, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	if err != nil {
		return 0, err
	}

	results := re.FindAllStringSubmatch(s, -1)
	ret := 0

	for _, val := range results {
		fst, _ := strconv.Atoi(val[1])
		snd, _ := strconv.Atoi(val[2])
		ret += fst * snd
	}

	return ret, nil
}

func PartTwo(s string) (int, error) {
	ret, err := CalculateInput(s)
	if err != nil {
		return 0, err
	}

	return ret, nil
}
