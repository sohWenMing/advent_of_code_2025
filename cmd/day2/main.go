package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sohWenMing/advent_of_code/internal/day2sequencing"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var b bytes.Buffer
	_, err = b.ReadFrom(file)
	if err != nil {
		panic(err)
	}
	stringFromBytes := strings.TrimSuffix(b.String(), "\n")
	numRanges := strings.Split(stringFromBytes, ",")

	invalids := []int64{}
	sum := int64(0)

	for _, numRange := range numRanges {
		numRangeFromString, err := day2sequencing.GetNumRangeFromString(numRange)
		if err != nil {
			log.Fatalf("unexpected error in getting numRanges: %v\n", err)
		}
		for _, num := range numRangeFromString {
			numString := fmt.Sprintf("%d", num)
			if day2sequencing.IsRepeated(numString) {
				invalids = append(invalids, int64(num))
			}
		}
	}
	for _, invalid := range invalids {
		sum += invalid
	}
	fmt.Println("sum: ", sum)
}
