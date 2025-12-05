package readfile

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	RIGHT = iota
	LEFT
)

type DirectionAndCount struct {
	Direction Direction
	Count     int
}

func (d DirectionAndCount) PrettyJSON() string {
	jsonBytes, _ := json.MarshalIndent(d, "", "    ")
	return string(jsonBytes)
}

func ReadFile(filePath string, numLines int) ([]DirectionAndCount, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	returnedDirectionsAndCount, err := ReadLinesFromFile(file, numLines)
	return returnedDirectionsAndCount, nil
}

func ReadLinesFromFile(r io.Reader, numLines int) ([]DirectionAndCount, error) {
	returnedDirectionsAndCount := []DirectionAndCount{}
	linesRead := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if linesRead >= numLines {
			break
		}
		line := scanner.Text()
		directionAndCount, err := processLine(line)
		if err != nil {
			return nil, err
		}
		returnedDirectionsAndCount = append(returnedDirectionsAndCount, directionAndCount)
		linesRead++
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
	count, err := strconv.Atoi(countString)

	if err != nil {
		return returnedDirectionAndCount, err
	}
	switch prefix {
	case "R":
		returnedDirectionAndCount.Direction = RIGHT
	default:
		returnedDirectionAndCount.Direction = LEFT
	}
	returnedDirectionAndCount.Count = count
	return returnedDirectionAndCount, nil
}
