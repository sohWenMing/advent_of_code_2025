package day3internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func GetLargestPossibleNumWithRemainingChars(input string, minChars int) (char string, err error) {
	if len(input) < minChars {
		return "", errors.New("length of input less than minChars")
	}
	// here if we are already at the number of chars we need, just return
	if len(input) == minChars {
		return input, nil
	}
	// if we only need one last character, get the largest
	if minChars == 1 {
		numInt, err := GetLargestFromString(input)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d", numInt), nil
	}
	if minChars == 0 {
		return "", nil
	}
	largest := 0
	stopIdx := 0
	for i, char := range input {
		num, err := strconv.ParseInt(string(char), 0, 10)
		if err != nil {
			return "", err
		}
		numInt := int(num)
		if numInt > largest {
			largest = numInt
			stopIdx = i
		}
		if len(input[i:]) <= minChars {
			break
		}
	}
	var b strings.Builder
	b.WriteString(string(input[stopIdx]))
	returnedChar, err := GetLargestPossibleNumWithRemainingChars(
		input[stopIdx+1:], minChars-1,
	)
	if err != nil {
		return "", err
	}
	b.WriteString(returnedChar)
	return b.String(), nil
}

func GetLargestPossibleNumFromString(input string) (returned int, err error) {
	if len(input) <= 2 {
		return returned, errors.New("input length was 0")
	}
	numToPos, err := GetFirstLargestFromInput(input)
	if err != nil {
		return 0, err
	}
	leftString := input[numToPos.Pos+1:]
	secondNum, err := GetLargestFromString(leftString)
	if err != nil {
		return 0, err
	}
	finalNumString := fmt.Sprintf("%d%d", numToPos.Num, secondNum)
	finalNum, err := strconv.ParseInt(finalNumString, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(finalNum), nil
}

type NumToPos struct {
	Num int
	Pos int
}

func PrettyJson(input any) string {
	jsonBytes, _ := json.MarshalIndent(input, "", "   ")
	return string(jsonBytes)

}
func GetFirstLargestFromInput(input string) (numtoPos NumToPos, err error) {
	returned := NumToPos{0, 0}
	if len(input) == 0 {
		return returned, errors.New("input length was 0")
	}
	last := len(input) - 1
	for i, char := range input {
		if i >= last {
			return returned, nil
		}
		num, err := convertRuneToInt(char)
		if err != nil {
			return NumToPos{}, err
		}
		if num > returned.Num {
			returned.Num = num
			returned.Pos = i
		}
	}
	return returned, nil
}

func GetLargestFromString(input string) (int, error) {
	largest := 0
	for _, char := range input {
		num, err := convertRuneToInt(char)
		if err != nil {
			return 0, err
		}
		if num > largest {
			largest = num
		}
	}
	return largest, nil
}

func convertRuneToInt(char rune) (int, error) {
	num64, err := strconv.ParseInt(string(char), 10, 0)
	if err != nil {
		return 0, fmt.Errorf("error while parsing character %s %w\n",
			string(char), err)
	}
	return int(num64), nil
}
