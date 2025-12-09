package day5internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StartEnd struct {
	Start int64
	End   int64
}

func GetNumAvailable(filepath string) (numAvailable int, err error) {
	nums, startEnds, err := GetNumsAndStartEndsFromFile(filepath)
	if err != nil {
		return 0, err
	}
	numAvailable = 0
	for _, num := range nums {
		for _, startEnd := range startEnds {
			if num >= startEnd.Start && num <= startEnd.End {
				numAvailable++
				break
			}
		}
	}
	return numAvailable, nil
}

func GetNumsAndStartEndsFromFile(filepath string) (
	nums []int64, startEnds []StartEnd, err error,
) {
	freshLines, availableLines, err := ReadFileGetLines(filepath)
	if err != nil {
		return nil, nil, err
	}
	nums, err = getNumsFromAvailableLines(availableLines)
	if err != nil {
		return nil, nil, err
	}
	startEnds, err = getStartEndsFromFreshLines(freshLines)
	if err != nil {
		return nil, nil, err
	}
	return nums, startEnds, nil
}

func ReadFileGetLines(filePath string) (freshLines []string, availableLines []string, err error) {
	isFreshLinesOver := false
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	//init the slices
	freshLines = []string{}
	availableLines = []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			isFreshLinesOver = true
			continue
		}
		switch isFreshLinesOver {
		case false:
			freshLines = append(freshLines, line)
			continue
		default:
			availableLines = append(availableLines, line)
			continue
		}
	}
	return freshLines, availableLines, nil
}

func getNumsFromAvailableLines(availableLines []string) (nums []int64, err error) {
	nums = []int64{}
	for _, line := range availableLines {
		num, err := ParseStringToNum(line)
		if err != nil {
			return nil, fmt.Errorf("num could not be parsed: %s\n", line)
		}
		nums = append(nums, num)
	}
	return nums, nil
}
func getStartEndsFromFreshLines(freshLines []string) (startEnds []StartEnd, err error) {
	startEnds = []StartEnd{}
	for _, line := range freshLines {
		startEnd, err := ParseStringToStartEnd(line)
		if err != nil {
			return nil, fmt.Errorf("line could not be parsed to startend: %s\n", line)
		}
		startEnds = append(startEnds, startEnd)
	}
	return startEnds, nil
}

func ParseStringToNum(input string) (num int64, err error) {
	num, parseErr := strconv.ParseInt(input, 10, 0)
	if parseErr != nil {
		return 0, fmt.Errorf("invalid input: %s", input)
	}
	return num, nil
}

func ParseStringToStartEnd(input string) (startEnd StartEnd, err error) {
	startEnd = StartEnd{}
	stringSlice := strings.Split(input, "-")
	if len(stringSlice) != 2 {
		return startEnd, fmt.Errorf("invalid input: %s", input)
	}
	start, parseErr := strconv.ParseInt(stringSlice[0], 10, 0)
	if parseErr != nil {
		return startEnd, fmt.Errorf("invalid input: %s", input)
	}
	end, parseErr := strconv.ParseInt(stringSlice[1], 10, 0)
	if parseErr != nil {
		return startEnd, fmt.Errorf("invalid input: %s", input)
	}
	if end < start {
		return startEnd, fmt.Errorf("invalid input: %s", input)
	}
	startEnd.Start = start
	startEnd.End = end
	return startEnd, nil
}
