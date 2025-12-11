package main

import (
	"fmt"

	"github.com/sohWenMing/advent_of_code/internal/day6internal"
)

func main() {
	sum, err := day6internal.Part1("./input.txt", "+")
	if err != nil {
		panic(err)
	}
	fmt.Println("sum: ", sum)

}
