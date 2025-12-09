package day4internal

import (
	"bufio"
	"fmt"
	"os"
)

// key should be x and y coord concat with |
// cell should only exist in mpa if it is actually filled
type FilledCellMap map[string]struct{}

type Cell struct {
	XIdx int
	Yidx int
}
type surroundingCellDiff struct {
	xDiff int
	yDiff int
}

var surroundingCellDiffs = []surroundingCellDiff{
	{-1, -1}, // top left
	{0, -1},  //  top
	{1, -1},  // top right
	{-1, 0},  // left
	{1, 0},   // right
	{-1, 1},  // bottom left
	{0, 1},   // bottom
	{1, 1},   // bottom right
}

func ReadFileGetMovable(filePath string) (numCanMove int, err error) {
	filledCellmap, allCells, err := GetCellsValsFromFile(filePath)
	if err != nil {
		return 0, err
	}
	numCanMove = 0
	for _, cell := range allCells {
		if !CheckFilledCellExist(filledCellmap, cell) {
			continue
		}
		if CheckCanMove(cell, filledCellmap) {
			numCanMove++
		}
	}
	return numCanMove, nil
}

func CheckCanMove(cell Cell, filledCellmap FilledCellMap) bool {
	numFilled := getNumSurroundingFilled(cell, filledCellmap)
	if numFilled < 4 {
		return true
	}
	return false
}

func getNumSurroundingFilled(cell Cell, filledCellmap FilledCellMap) int {
	numFilled := 0
	surroundingCells := GetSurroundingCells(cell)
	filledSurroundingCells := []Cell{}
	for _, surroundingCell := range surroundingCells {
		if CheckFilledCellExist(filledCellmap,
			surroundingCell) {
			numFilled++
			filledSurroundingCells = append(filledSurroundingCells, surroundingCell)
		}
	}
	return numFilled
}

func CheckFilledCellExist(cellmap FilledCellMap, cell Cell) bool {
	cellKey := fmt.Sprintf("%d|%d", cell.XIdx, cell.Yidx)
	_, ok := cellmap[cellKey]
	return ok
}

func GetSurroundingCells(cell Cell) (surroundingCells []Cell) {
	surroundingCells = []Cell{}
	for _, diff := range surroundingCellDiffs {
		appendCell := Cell{
			cell.XIdx + diff.xDiff,
			cell.Yidx + diff.yDiff,
		}
		surroundingCells = append(surroundingCells, appendCell)
	}
	return surroundingCells
}

func GetCellsValsFromFile(filePath string) (cellmap FilledCellMap, allCells []Cell, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	cellmap = make(FilledCellMap)
	allCells = []Cell{}

	scanner := bufio.NewScanner(file)
	yIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		readToFilledCellMap(cellmap, line, yIdx)
		readToAllCells(&allCells, line, yIdx)
		yIdx++
	}
	return cellmap, allCells, nil
}

func readToFilledCellMap(cellMap FilledCellMap, line string, yIdx int) {
	for i, cellVal := range line {
		if string(cellVal) == "@" {
			cellKey := fmt.Sprintf("%d|%d", i, yIdx)
			cellMap[cellKey] = struct{}{}
		}
	}
}

func readToAllCells(allCells *[]Cell, line string, yIdx int) {
	for i, _ := range line {
		cell := Cell{
			i, yIdx,
		}
		*allCells = append(*allCells, cell)
	}
}
