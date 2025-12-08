package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/sohWenMing/advent_of_code/internal/day3internal"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sum := 0
	sum2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		//for part 1 of day 3
		num, err := day3internal.GetLargestPossibleNumFromString(input)
		if err != nil {
			panic(err)
		}
		sum += num

		//for part 2 of day 3
		numString, err := day3internal.GetLargestPossibleNumWithRemainingChars(input, 12)
		numInt, err := strconv.ParseInt(numString, 10, 0)
		if err != nil {
			panic(err)
		}
		sum2 += int(numInt)
	}
	fmt.Println("sum: ", sum)
	fmt.Println("sum2: ", sum2)
}
