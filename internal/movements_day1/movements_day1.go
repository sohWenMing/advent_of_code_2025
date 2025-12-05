package movementsday1

import (
	"encoding/json"
	"fmt"
)

type MovementState struct {
	ZeroCounter int
	CurrentPos  int
	MaxPos      int
}

func InitMovementState(startingPos, maxPos int) *MovementState {
	var zeroCounter int
	if startingPos == 0 {
		zeroCounter = 1
	} else {
		zeroCounter = 0
	}

	return &MovementState{
		zeroCounter, startingPos, maxPos,
	}
}

const right = "RIGHT"
const left = "LEFT"

type DirectionAndCountInterface interface {
	GetDirection() string
	GetCount() int
}

func (ms *MovementState) Move(d DirectionAndCountInterface) error {
	direction := d.GetDirection()
	switch direction {
	case right:
		ms.MoveRight(d.GetCount())
		return nil
	case left:
		ms.MoveLeft(d.GetCount())
		return nil
	default:
		return fmt.Errorf("direction not valid. Direction: %s\n", direction)
	}
}

func (ms *MovementState) PrettyJSON() string {
	jsonBytes, _ := json.MarshalIndent(ms, "", "    ")
	return string(jsonBytes)
}

func (ms *MovementState) GetPos() int {
	return ms.CurrentPos
}
func (ms *MovementState) GetZeroCount() int {
	return ms.ZeroCounter
}

func (ms *MovementState) MoveRight(count int) {
	numTilMax := ms.MaxPos - ms.CurrentPos
	if numTilMax >= count {
		ms.CurrentPos += count
	} else {
		ms.CurrentPos = 0 + (count - numTilMax - 1)
	}
	if ms.CurrentPos == 0 {
		ms.ZeroCounter++
	}
}
func (ms *MovementState) MoveLeft(count int) {
	numTilZero := ms.CurrentPos - 0
	if numTilZero >= count {
		ms.CurrentPos -= count
	} else {
		ms.CurrentPos = ms.MaxPos - (count - numTilZero - 1)
	}
	if ms.CurrentPos == 0 {
		ms.ZeroCounter++
	}
}
