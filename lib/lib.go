package lib

import (
	"slices"
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

func ReverseSlice[K comparable](sl []K) []K {
	var ret []K

	for _, val := range slices.Backward(sl) {
		ret = append(ret, val)
	}
	return ret
}
