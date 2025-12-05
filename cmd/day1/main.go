package main

import (
	"fmt"

	movementsday1 "github.com/sohWenMing/advent_of_code/internal/movements_day1"
	readFile "github.com/sohWenMing/advent_of_code/internal/readfile_day1"
)

func main() {
	directionsAndCount, err := readFile.ReadFile("./input.txt", 5000)
	if err != nil {
		panic(err)
	}

	movementState := movementsday1.InitMovementState(50, 99)
	fmt.Println("start movement state")
	fmt.Println("movement state: ", movementState.PrettyJSON())
	fmt.Println("")
	fmt.Println("")

	for _, record := range directionsAndCount {

		movementStateBeforeMoveJSON := movementState.PrettyJSON()
		zeroCountBeforeMove := movementState.GetZeroCount()
		err := movementState.Move(record)
		if err != nil {
			panic(err)
		}
		zeroCounterAftermove := movementState.GetZeroCount()
		movementStateAfterMoveJSON := movementState.PrettyJSON()
		if zeroCounterAftermove > zeroCountBeforeMove {
			fmt.Println("movement state before move")
			fmt.Println(movementStateBeforeMoveJSON)
			fmt.Println("=========================")
			fmt.Println("")
			fmt.Println("")
			fmt.Println("record")
			fmt.Println(record.PrettyJSON())
			fmt.Println("=========================")
			fmt.Println("")
			fmt.Println("")
			fmt.Println("movement state after move")
			fmt.Println(movementStateAfterMoveJSON)
			fmt.Println("=========================")
			fmt.Println("")
			fmt.Println("")
		}

	}
	fmt.Println("Zero Count: ", movementState.GetZeroCount())

}
