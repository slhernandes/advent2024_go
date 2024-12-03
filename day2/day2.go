package day2

import (
	"strconv"
	"strings"
)

func filterEmpty(sl []string) []string {
	var ret []string
	for _, val := range sl {
		if len(val) != 0 {
			ret = append(ret, val);
		}
	}
	return ret
}

func parseInput(s string) [][]string {
	var ret [][]string
	rows := filterEmpty(strings.Split(s, "\n"))
	for _, val := range rows {
		val_split := filterEmpty(strings.Split(val, " "))
		ret = append(ret, val_split)
	}
	return ret
}

func isSafeDecr(sl []string) (bool, error) {
	if len(sl) <= 1 {
		return true, nil
	}
	fst, err := strconv.Atoi(sl[0])
	if err != nil {
		return false, err
	}
	snd, err := strconv.Atoi(sl[1])
	if err != nil {
		return false, err
	}
	if fst - snd >= 1 && fst - snd <= 3 {
		return isSafeDecr(sl[1:])
	}
	return false, nil
}

func isSafeIncr(sl []string) (bool, error) {
	if len(sl) <= 1 {
		return true, nil
	}
	fst, err := strconv.Atoi(sl[0])
	if err != nil {
		return false, err
	}
	snd, err := strconv.Atoi(sl[1])
	if err != nil {
		return false, err
	}
	if snd - fst >= 1 && snd - fst <= 3 {
		return isSafeIncr(sl[1:])
	}
	return false, nil
}

func isSafeDecrDampener(sl []string) (bool, error) {
	if len(sl) <= 2 {
		return true, nil
	}
	fst, err := strconv.Atoi(sl[0])
	if err != nil {
		return false, err
	}
	snd, err := strconv.Atoi(sl[1])
	if err != nil {
		return false, err
	}
	thd, err := strconv.Atoi(sl[2])
	if err != nil {
		return false, err
	}
	goodFstPair := fst - snd >= 1 && fst - snd <= 3
	goodSndPair := fst - thd >= 1 && fst - thd <= 3
	if goodFstPair && goodSndPair {
		fstRes, err := isSafeDecrDampener(sl[1:]) 
		if err != nil {
			return false, err
		}
		sndRes, err := isSafeDecr(sl[2:]) 
		if err != nil {
			return false, err
		}
		return fstRes || sndRes, nil
	}
	if goodFstPair {
		return isSafeDecrDampener(sl[1:])
	}
	if goodSndPair {
		return isSafeDecr(sl[2:])
	}
	return false, nil
}

func isSafeIncrDampener(sl []string) (bool, error) {
	if len(sl) <= 2 {
		return true, nil
	}
	fst, err := strconv.Atoi(sl[0])
	if err != nil {
		return false, err
	}
	snd, err := strconv.Atoi(sl[1])
	if err != nil {
		return false, err
	}
	thd, err := strconv.Atoi(sl[2])
	if err != nil {
		return false, err
	}
	goodFstPair := snd - fst >= 1 && snd - fst <= 3
	goodSndPair := thd - fst >= 1 && thd - fst <= 3
	if goodFstPair && goodSndPair {
		fstRes, err := isSafeIncrDampener(sl[1:]) 
		if err != nil {
			return false, err
		}
		sndRes, err := isSafeIncr(sl[2:]) 
		if err != nil {
			return false, err
		}
		return fstRes || sndRes, nil
	}
	if goodFstPair {
		return isSafeIncrDampener(sl[1:])
	}
	if goodSndPair {
		return isSafeIncr(sl[2:])
	}
	return false, nil
}

func isSafe(sl []string) (bool, error) {
	decrSafe, err := isSafeDecr(sl)
	if err != nil {
		return false, err
	}
	incrSafe, err := isSafeIncr(sl)
	if err != nil {
		return false, err
	}
	return incrSafe || decrSafe, nil
}

func isSafeDampener(sl []string) (bool, error) {
	decrSafeDmp, err := isSafeDecrDampener(sl)
	if err != nil {
		return false, err
	}

	decrSafe, err := isSafeDecr(sl[1:])
	if err != nil {
		return false, err
	}

	incrSafeDmp, err := isSafeIncrDampener(sl)
	if err != nil {
		return false, err
	}

	incrSafe, err := isSafeIncr(sl[1:])
	if err != nil {
		return false, err
	}

	return incrSafeDmp || decrSafeDmp || decrSafe || incrSafe, nil
}

func PartOne(s string) (int,error) {
	var ret int
	input := parseInput(s)
	for _, v := range input {
		res, err := isSafe(v)
		if err != nil {
			return 0, err
		}
		if res {
			ret++
		}
	}
	return ret, nil
}

func PartTwo(s string) (int,error) {
	var ret int
	input := parseInput(s)
	for _, v := range input {
		res, err := isSafeDampener(v)
		if err != nil {
			return 0, err
		}
		if res {
			ret++
		}
	}
	return ret, nil
}
