package movementsday1

import (
	"testing"
)

type testDirectionAndCountInterface struct {
	direction string
	count     int
}

func (t testDirectionAndCountInterface) GetDirection() string {
	return t.direction
}

func (t testDirectionAndCountInterface) GetCount() int {
	return t.count
}

func checkPosition(t *testing.T, expectedPos int, movementState *MovementState) {
	gotPos := movementState.GetPos()
	if gotPos != expectedPos {
		t.Errorf("got %d\nwant %d", gotPos, expectedPos)
	}
}

func TestMoveTwoRight(t *testing.T) {
	movementState := InitMovementState(98, 99)
	testInterfaces := []testDirectionAndCountInterface{
		{
			"RIGHT", 2,
		},
		{
			"RIGHT", 1,
		},
	}
	for _, testInterface := range testInterfaces {
		err := movementState.Move(testInterface)
		if err != nil {
			t.Errorf("didn't expect error, got err %v\n", err)
			return
		}
	}
	expectedPos := 101
	checkPosition(t, expectedPos, movementState)
	expectedCounter := 1
	checkCounter(t, movementState, expectedCounter)
}
func TestMoveTwoLeft(t *testing.T) {
	movementState := InitMovementState(1, 99)
	testInterfaces := []testDirectionAndCountInterface{
		{
			"LEFT", 1,
		},
		{
			"LEFT", 2,
		},
	}
	for _, testInterface := range testInterfaces {
		err := movementState.Move(testInterface)
		if err != nil {
			t.Errorf("didn't expect error, got err %v\n", err)
			return
		}
	}
	expectedPos := -2
	checkPosition(t, expectedPos, movementState)
	expectedCounter := 1
	checkCounter(t, movementState, expectedCounter)
}

func checkCounter(t *testing.T, movementState *MovementState, expectedZeroCounter int) {
	gotZeroCounter := movementState.GetZeroCount()
	if gotZeroCounter != expectedZeroCounter {
		t.Errorf("got %d\nwant %d", gotZeroCounter, expectedZeroCounter)
	}
}
