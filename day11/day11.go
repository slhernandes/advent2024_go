package aoc

import (
	"aoc/lib"
	"math"
	"strconv"
)

type ident struct {
	num int64
	cnt int
}

func ProcessInputMap(nums []int64, count int) int {
	dict := make(map[ident]int)
	ret := 0

	var helper func(num int64, count int) int
	helper = func(num int64, count int) int {
		if count == 0 {
			return 1
		}

		ret, ok := dict[ident{num: num, cnt: count}]
		if ok {
			return ret
		}

		num_len := int(math.Log10(float64(num))) + 1
		if num == 0 {
			temp := helper(1, count-1)
			dict[ident{num: 0, cnt: count}] = temp
			return temp
		} else if num_len%2 == 0 {
			div := int64(math.Pow10(num_len / 2))
			left := helper(num/div, count-1)
			right := helper(num%div, count-1)
			dict[ident{num: num / div, cnt: count - 1}] = left
			dict[ident{num: num % div, cnt: count - 1}] = right
			return left + right
		}
		temp := helper(num*2024, count-1)
		dict[ident{num: num, cnt: count}] = temp
		return temp
	}

	for _, v := range nums {
		ret += helper(v, count)
	}
	return ret
}

func PartOne(s string) (int, error) {
	s_new := lib.SplitFilterEmpty(s, "\n")[0]
	temp := lib.SplitFilterEmpty(s_new, " ")
	nums := make([]int64, 0)

	for _, v := range temp {
		temp_num, _ := strconv.Atoi(v)
		nums = append(nums, int64(temp_num))
	}

	ret := ProcessInputMap(nums, 25)

	return ret, nil
}

func PartTwo(s string) (int, error) {
	s_new := lib.SplitFilterEmpty(s, "\n")[0]
	temp := lib.SplitFilterEmpty(s_new, " ")
	nums := make([]int64, 0)

	for _, v := range temp {
		temp_num, _ := strconv.Atoi(v)
		nums = append(nums, int64(temp_num))
	}

	ret := ProcessInputMap(nums, 75)

	return ret, nil
}
