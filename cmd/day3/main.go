package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sohWenMing/advent_of_code/internal/day3internal"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := day3internal.GetLargestPossibleNumFromString(input)
		// fmt.Println("num: ", num)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Println("sum: ", sum)
}
