package main

import (
	"fmt"

	"github.com/sohWenMing/advent_of_code/internal/day6internal"
)

func main() {
	Part1()
	Part2()
}

func Part1() {
	sum, err := day6internal.Part1("./input.txt", "+")
	if err != nil {
		panic(err)
	}
	fmt.Println("sum: ", sum)
}

func Part2() {
	bytes, err := day6internal.Part2ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := day6internal.GetLines(bytes, []byte("\n"))
	records := day6internal.GetStructs(lines)
	result := 0
	for _, record := range records {
		operationResult, err := record.Operate()
		if err != nil {
			panic(err)
		}
		result += operationResult
	}
	fmt.Println("part2 result: ", result)
}
