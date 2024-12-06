package day5

import (
	"aoc/lib"
	"errors"
	"math/big"
	"slices"
	"strconv"
)

func ParseRules(s string) map[int](*big.Int) {
	ret := make(map[int](*big.Int))
	ruleList := lib.SplitFilterEmpty(s, "\n")
	for _, v := range ruleList {
		tmp := lib.SplitFilterEmpty(v, "|")
		fst, _ := strconv.Atoi(tmp[0])
		snd, _ := strconv.Atoi(tmp[1])
		_, ok := ret[snd]
		if !ok {
			ret[snd] = big.NewInt(0)
		}
		ret[snd].SetBit(ret[snd], fst, 1)
	}
	return ret
}

func ParsePages(s string) [][]int {
	sArr := lib.SplitFilterEmpty(s, "\n")
	var ret [][]int
	for _, str := range sArr {
		arr := lib.SplitFilterEmpty(str, ",")
		var arrInt []int
		for _, v := range arr {
			tmp, _ := strconv.Atoi(v)
			arrInt = append(arrInt, tmp)
		}
		ret = append(ret, arrInt)
	}
	return ret
}

func CheckPages(pages []int, rules map[int](*big.Int)) int {
	mid := len(pages)/2
	now := big.NewInt(0)
	for _, v := range slices.Backward(pages) {
		check := big.NewInt(0)
		_, ok := rules[v]
		if !ok {
			now.SetBit(now, v, 1)
			continue
		}
		check.And(now, rules[v])
		if check.Cmp(big.NewInt(0)) != 0 {
			return 0
		}
		now.SetBit(now, v, 1)
	}
	return pages[mid]
}

func SortedMid(pages []int, rules map[int](*big.Int)) (int, error) {
	mid := len(pages)/2
	now := big.NewInt(0)
	var reIndex []int
	for i := len(pages) - 1; i >= 0; i-- {
		check := big.NewInt(0)
		_, ok := rules[pages[i]]
		if !ok {
			now.SetBit(now, pages[i], 1)
			continue
		}
		check.And(now, rules[pages[i]])
		if check.Cmp(big.NewInt(0)) != 0 {
			reIndex = append(reIndex, pages[i])
			pages = append(pages[:i], pages[i+1:]...)
		} else {
			now.SetBit(now, pages[i], 1)
		}
	}

	for _, val := range reIndex {
		nowTmp := big.NewInt(0)
		nowTmp.Add(now, nowTmp)
		rule, ok := rules[val]
		if !ok {
			err := errors.New("The rule has nil value")
			return 0, err
		}
		var i int
		for i = range pages {
			check := big.NewInt(0)
			check.And(nowTmp, rule)
			if check.Cmp(big.NewInt(0)) != 0 {
				nowTmp.SetBit(nowTmp, pages[i], 0)
			} else {
				i--
				break
			}
		}
		if i < len(pages) {
			pages = append(pages, 0)
			copy(pages[i+2:], pages[i+1:])
			pages[i+1] = val
		} else {
			pages = append(pages, val)
		}
		now.SetBit(now, val, 1)
	}
	return pages[mid], nil
}

func PartOne(s string) (int,error) {
	sl := lib.SplitFilterEmpty(s, "\n\n")
	rules := sl[0]
	pages := sl[1]
	ruleMap := ParseRules(rules)
	pageArr := ParsePages(pages)

	ret := 0
	for _, val := range pageArr {
		ret += CheckPages(val, ruleMap)
	}
	return ret, nil
}

func PartTwo(s string) (int,error) {
	sl := lib.SplitFilterEmpty(s, "\n\n")
	rules := sl[0]
	pages := sl[1]
	ruleMap := ParseRules(rules)
	pageArr := ParsePages(pages)
	var unsorted []int

	for i, val := range pageArr {
		if CheckPages(val, ruleMap) == 0 {
			unsorted = append(unsorted, i)
		}
	}

	ret := 0
	for _, val := range unsorted {
		 tmp, err := SortedMid(pageArr[val], ruleMap)
		 if err != nil {
		 	return 0, err
		 }
		 ret += tmp
	}
	return ret, nil
}
