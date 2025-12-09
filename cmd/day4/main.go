package main

import (
	"fmt"

	"github.com/sohWenMing/advent_of_code/internal/day4internal"
)

func main() {
	part1()
	part2()
}

func part2() {
	cellMap, allCells, err := day4internal.GetCellsValsFromFile("./input.txt")
	if err != nil {
		panic(err)
	}
	finalNumCanMove := 0
	for {
		numCanMove := getNumCanMoveAndRemove(cellMap, allCells)
		if numCanMove == 0 {
			break
		} else {
			finalNumCanMove += numCanMove
		}
	}
	fmt.Println("finalNumCanMove: ", finalNumCanMove)
}

func part1() {
	numCanMove, err := day4internal.ReadFileGetMovable("./input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("numCanMove: ", numCanMove)
}

func getNumCanMoveAndRemove(
	cellMap day4internal.FilledCellMap,
	allCells []day4internal.Cell,
) (numCanMove int) {
	numCanMove = 0
	cellsToRemove := []day4internal.Cell{}
	for _, cell := range allCells {
		if !day4internal.CheckFilledCellExist(cellMap, cell) {
			continue
		}
		if day4internal.CheckCanMove(cell, cellMap) {
			numCanMove++
			cellsToRemove = append(cellsToRemove, cell)
		}
	}
	for _, cellToRemove := range cellsToRemove {
		removeCellFromMap(cellToRemove, cellMap)
	}
	return numCanMove
}

func removeCellFromMap(cell day4internal.Cell, filledCellmap day4internal.FilledCellMap) {
	key := fmt.Sprintf("%d|%d", cell.XIdx, cell.Yidx)
	delete(filledCellmap, key)
}
