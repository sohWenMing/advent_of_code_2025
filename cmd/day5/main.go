package main

import (
	"fmt"

	"github.com/sohWenMing/advent_of_code/internal/day5internal"
)

func main() {
	part2()
}

func part1() {
	numAvailable, err := day5internal.GetNumAvailable("./input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("numAvailable: ", numAvailable)
}

func part2() {
	fmt.Println("part2 started")
	numIngredients, err := day5internal.GetPart2("./input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("numIngredients part 2:", numIngredients)

}
