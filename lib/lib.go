package lib

import (
	"slices"
	"strings"
)

type Coordinate struct {
	X, Y int
}

func Diff(a Coordinate, b Coordinate) Coordinate {
	return Coordinate{a.X - b.X, a.Y - b.Y}
}

func Add(a Coordinate, b Coordinate) Coordinate {
	return Coordinate{a.X + b.X, a.Y + b.Y}
}

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
