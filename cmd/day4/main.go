package main

import (
	"fmt"

	"github.com/sohWenMing/advent_of_code/internal/day4internal"
)

func main() {
	numCanMove, err := day4internal.ReadFileGetMovable("./input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("numCanMove: ", numCanMove)

}
