package main

import (
	"fmt"

	movementsday1 "github.com/sohWenMing/advent_of_code/internal/movements_day1"
	readFile "github.com/sohWenMing/advent_of_code/internal/readfile_day1"
)

func main() {
	directionsAndCount, err := readFile.ReadFile("./testinput.txt", 5000)
	if err != nil {
		panic(err)
	}

	movementState := movementsday1.InitMovementState(50, 99)

	for i, record := range directionsAndCount {
		if i == len(directionsAndCount)-1 {
			fmt.Println("movement state before: ", movementState.PrettyJSON())
		}

		// zeroCountBeforeMove := movementState.GetZeroCount()
		err := movementState.Move(record)
		if err != nil {
			panic(err)
		}
		if i == len(directionsAndCount)-1 {
			fmt.Println("record: ", record.PrettyJSON())
			fmt.Println("movement state after: ", movementState.PrettyJSON())
		}

		// zeroCounterAftermove := movementState.GetZeroCount()

	}
	fmt.Println("Zero Count: ", movementState.GetZeroCount())

}
