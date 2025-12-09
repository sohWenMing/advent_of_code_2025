package main

import (
	"fmt"

	"github.com/sohWenMing/advent_of_code/internal/day5internal"
)

func main() {
	numAvailable, err := day5internal.GetNumAvailable("./input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("numAvailable: ", numAvailable)
}
