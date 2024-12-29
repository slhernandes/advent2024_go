package aoc

import (
	"aoc/lib"
	"slices"
	"strconv"
)

type pii struct {
	fst, snd int
}

func Sum(a []int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func Partition(nums []int, spaces []int, length int) ([]int, []int) {
	var rest []int
	acc := 0
	i, j := 0, 0
	for {
		if i < len(nums) {
			acc += nums[i]
			i++
		}
		if acc >= length {
			rest = lib.ReverseSlice(nums[i-1:])
			nums = nums[:i]
			nums[i-1] = nums[i-1] - acc + length
			rest[len(rest)-1] = acc - length
			break
		}
		if j < len(spaces) {
			acc += spaces[j]
			j++
		}
		if acc >= length {
			break
		}
	}
	return nums, rest
}

func Coeff(start int, count int) int {
	return (start+count)*(start+count-1)/2 - start*(start-1)/2
}

func PartOne(s string) (int, error) {
	nums := []int{}
	spaces := []int{}
	for i, v := range s {
		if v == '\n' {
			break
		}
		tmp, err := strconv.Atoi(string(v))
		if err != nil {
			return 0, err
		}
		if i%2 == 0 {
			nums = append(nums, tmp)
		} else {
			spaces = append(spaces, tmp)
		}
	}
	sum_nums := Sum(nums)
	length_nums := len(nums)
	nums, rest := Partition(nums, spaces, sum_nums)
	ret := 0
	now_r, now_l, idx := length_nums-1, 0, 0
	for even_done, odd_done, parity := false, false, false; even_done == false || odd_done == false; {
		if parity {
			if len(spaces) == 0 || len(rest) == 0 {
				odd_done = true
				parity = !parity
				continue
			}
			if spaces[0] > rest[0] {
				ret += Coeff(idx, rest[0]) * now_r
				idx += rest[0]
				now_r--
				spaces[0] = spaces[0] - rest[0]
				rest = rest[1:]
			} else if spaces[0] < rest[0] {
				ret += Coeff(idx, spaces[0]) * now_r
				idx += spaces[0]
				rest[0] = rest[0] - spaces[0]
				spaces = spaces[1:]
				parity = !parity
			} else {
				ret += Coeff(idx, rest[0]) * now_r
				idx += rest[0]
				now_r--
				spaces = spaces[1:]
				rest = rest[1:]
				parity = !parity
			}
		} else {
			if len(nums) == 0 {
				even_done = true
				continue
			}
			ret += Coeff(idx, nums[0]) * now_l
			idx += nums[0]
			now_l++
			nums = nums[1:]
			parity = !parity
		}
	}
	return ret, nil
}

func PartTwo(s string) (int, error) {
	nums := []pii{}
	spaces := []pii{}
	acc := 0
	for i, v := range s {
		if v == '\n' {
			break
		}
		tmp, err := strconv.Atoi(string(v))
		if err != nil {
			return 0, err
		}
		if i%2 == 0 {
			nums = append(nums, pii{acc, tmp})
			acc += tmp
		} else {
			spaces = append(spaces, pii{acc, tmp})
			acc += tmp
		}
	}

	slices.Reverse(nums)
	ret := 0

	for i, v := range nums {
		for j, sp := range spaces {
			if sp.fst >= v.fst {
				break
			}
			if v.snd <= sp.snd {
				nums[i].fst = sp.fst
				spaces[j].fst += v.snd
				spaces[j].snd -= v.snd
				break
			}
		}
	}

	slices.Reverse(nums)
	for i, v := range nums {
		ret += Coeff(v.fst, v.snd) * i
	}

	return ret, nil
}
