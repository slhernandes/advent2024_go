package day7

import (
	"aoc/lib"
	"fmt"
	"math"
	"strconv"
)

func MakeTarget(target int64, nums []int64) bool {
	if len(nums) == 1 {
		return target == nums[0]
	}
	var multRes, addRes bool

	candMult := append([]int64(nil), nums...)
	candMult[1] *= candMult[0]

	candAdd := append([]int64(nil), nums...)
	candAdd[1] += candAdd[0]
	
	if candMult[1] > target {
		multRes = false
	} else {
		multRes = MakeTarget(target, candMult[1:])
	}

	if candAdd[1] > target {
		addRes = false
	} else {
		addRes = MakeTarget(target, candAdd[1:])
	}

	return addRes || multRes
}

func ConcatInts(a int64, b int64) int64 {
	ret, _ := strconv.ParseInt(fmt.Sprintf("%d%d", a, b), 10, 64)
	return ret
}

func MakeTargetConcat(target int64, nums []int64) bool {
	if len(nums) == 1 {
		return target == nums[0]
	}

	var multRes, addRes, catRes bool
	candMult := append([]int64(nil), nums...)
	candMult[1] *= candMult[0]

	candAdd := append([]int64(nil), nums...)
	candAdd[1] += candAdd[0]

	Digits := func (a int64) int {
		return int(math.Ceil(math.Log10(float64(a))))
	}

	candCat := append([]int64(nil), nums...)
	if Digits(target) < Digits(candCat[0]) + Digits(candCat[1]) {
		catRes = false
	} else {
		candCat[1] = ConcatInts(candCat[0], candCat[1])
		if candCat[1] > target {
			catRes = false
		} else {
			catRes = MakeTargetConcat(target, candCat[1:])
		}
	}
	
	if candMult[1] > target {
		multRes = false
	} else {
		multRes = MakeTargetConcat(target, candMult[1:])
	}

	if candAdd[1] > target {
		addRes = false
	} else {
		addRes = MakeTargetConcat(target, candAdd[1:])
	}

	return addRes || multRes || catRes
}

func PartOne(s string) (int64,error) {
	lines := lib.SplitFilterEmpty(s, "\n")
  ret := int64(0)

	for _, v := range lines {
		tmp := lib.SplitFilterEmpty(v, ": ")
		fst, _ := strconv.Atoi(tmp[0])

		var snd []int64
		tmp_arr := lib.SplitFilterEmpty(tmp[1], " ")

		for _, w := range tmp_arr {
			i, _ := strconv.Atoi(w)
			snd = append(snd, int64(i))
		}

		if MakeTarget(int64(fst), snd) {
			ret += int64(fst)
		}
	}

	return ret, nil
}
func PartTwo(s string) (int64,error) {
	lines := lib.SplitFilterEmpty(s, "\n")
  ret := int64(0)

	for _, v := range lines {
		tmp := lib.SplitFilterEmpty(v, ": ")
		fst, _ := strconv.Atoi(tmp[0])

		var snd []int64
		tmp_arr := lib.SplitFilterEmpty(tmp[1], " ")

		for _, w := range tmp_arr {
			i, _ := strconv.Atoi(w)
			snd = append(snd, int64(i))
		}

		if MakeTargetConcat(int64(fst), snd) {
			ret += int64(fst)
		}
	}

	return ret, nil
}
