package readfile

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	Right = iota
	Left
)

type DirectionAndCount struct {
	Direction Direction
	Count     int
}

func (d DirectionAndCount) PrettyJSON() string {
	jsonBytes, _ := json.MarshalIndent(d, "", "    ")
	return string(jsonBytes)
}

func ReadFile(filePath string) ([]DirectionAndCount, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	returnedDirectionsAndCount := []DirectionAndCount{}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		directionAndCount, err := processLine(line)
		if err != nil {
			return nil, err
		}
		returnedDirectionsAndCount = append(returnedDirectionsAndCount, directionAndCount)
	}
	return returnedDirectionsAndCount, nil
}

func processLine(line string) (DirectionAndCount, error) {
	returnedDirectionAndCount := DirectionAndCount{}
	rightPrefix := "R"
	leftPrefix := "L"
	if !strings.HasPrefix(line, rightPrefix) && !strings.HasPrefix(line, leftPrefix) {
		return returnedDirectionAndCount, errors.New("line does not have requiredPrefix")
	}
	if strings.HasPrefix(line, rightPrefix) {
		returnedDirectionAndCount, err := getDirectionAndCountFromLine(line, rightPrefix)
		return returnedDirectionAndCount, err
	} else {
		returnedDirectionAndCount, err := getDirectionAndCountFromLine(line, leftPrefix)
		return returnedDirectionAndCount, err
	}
}
func getDirectionAndCountFromLine(line, prefix string) (DirectionAndCount, error) {
	returnedDirectionAndCount := DirectionAndCount{}
	countString := strings.TrimPrefix(line, prefix)
	fmt.Println("countString", countString)
	count, err := strconv.Atoi(countString)
	fmt.Println("count: ", count)

	if err != nil {
		return returnedDirectionAndCount, err
	}
	switch prefix {
	case "R":
		returnedDirectionAndCount.Direction = Right
	default:
		returnedDirectionAndCount.Direction = Left
	}
	returnedDirectionAndCount.Count = count
	return returnedDirectionAndCount, nil
}
