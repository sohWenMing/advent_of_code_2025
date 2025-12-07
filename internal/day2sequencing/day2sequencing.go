package day2sequencing

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func CheckIsAllRepeatedSubString(input string) bool {
	if len(input) <= 1 {
		return false
	}
	middle := len(input) / 2
	curIdx := 0
	var b strings.Builder
	for len(b.String()) < middle {
		b.WriteByte(input[curIdx])
		subString := b.String()
		curString := input[curIdx+1:]
		if recurseCheckSubstr(subString, curString) == true {
			return true
		}
		curIdx++
	}
	return false
}

func recurseCheckSubstr(subString, curString string) bool {
	curString = strings.TrimPrefix(curString, subString)
	if len(curString) == 0 {
		return true
	}
	switch strings.HasPrefix(curString, subString) {
	case false:
		return false
	default:
		return recurseCheckSubstr(subString, curString)
	}
}

func IsRepeated(input string) bool {
	if len(input)%2 != 0 {
		return false
	}
	middle := len(input) / 2
	start := input[:middle]
	end := input[middle:]
	if start == end {
		return true
	}
	return false
}

func GetNumRangeFromString(input string) (returned []int, err error) {
	slice := strings.Split(input, "-")
	if len(slice) != 2 {
		return nil, fmt.Errorf("invalid input: %s\n", input)
	}
	start := slice[0]
	end := slice[1]
	startNum, err := strconv.ParseInt(start, 10, 0)
	if err != nil {
		return nil, err
	}
	endNum, err := strconv.ParseInt(end, 10, 0)
	if err != nil {
		return nil, err
	}
	returned = []int{}
	curNum := startNum
	for curNum <= endNum {
		returned = append(returned, int(curNum))
		curNum += 1
	}
	return returned, nil
}

func RunSubStringsFromStringCheckRepeats(input string) bool {
	if len(input) <= 1 {
		return false
	}
	curIdx := 0
	lastIdx := len(input) - 1
	for curIdx <= lastIdx {
		checkString := input[curIdx:]
		if CheckForRepeats(checkString) {
			return true
		}
		curIdx++
	}
	return false
}

func CheckForRepeats(input string) bool {
	if len(input) <= 1 {
		return false
	}

	buf := []byte{}
	curIdx := 0
	for curIdx < len(input)/2 {
		buf = append(buf, input[curIdx])
		checkAgainst := []byte(input[curIdx+1:])
		if bytes.Index(checkAgainst, buf) == 0 {
			return true
		}
		curIdx++
	}
	return false
}

func CheckHasRepeatSubset(input string) bool {
	if len(input) <= 1 {
		return false
	}
	subsets := getSubSetStrings(input)
	fmt.Println("subsets: ", subsets)
	for _, subset := range subsets {
		fmt.Println("checking for subset: ", subset)
		splits := strings.Split(input, subset)
		if len(splits) >= 2 {
			return true
		}
	}
	return false
}

func splitStringOnSubset(wholeString, subset string) []string {
	return strings.Split(wholeString, subset)
}

func getSubSetStrings(input string) []string {
	returned := []string{}
	byteSubsets := getSubsets([]byte(input))
	for _, slice := range byteSubsets {
		if len(slice) > 0 {
			returned = append(returned, string(slice))
		}
	}
	return returned
}

func getSubsets(input []byte) [][]byte {
	returned := [][]byte{}
	if len(input) == 0 {
		returned = append(returned, []byte{})
		return returned
	}
	current := input[0]
	minusFront := getSliceMinusFront(input)
	returnedFromGetSubsets := getSubsets(minusFront)
	for _, slice := range returnedFromGetSubsets {
		returned = append(returned, slice)
		currentSlice := []byte{current}
		concat := slices.Concat(currentSlice, slice)
		returned = append(returned, concat)
	}
	return returned
}

func getSliceMinusFront(input []byte) []byte {
	if len(input) == 1 {
		return []byte{}
	}
	return input[1:]
}
