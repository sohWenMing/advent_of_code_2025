package day5internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReadFromInput(t *testing.T) {
	numIngredients, err := GetPart2("./testinput.txt")
	if err != nil {
		t.Errorf("didn't expect error, got %v\n", err)
	}
	fmt.Println("numingredients: ", numIngredients)

}

func TestRecursiveAppendStartEnd(t *testing.T) {
	type test struct {
		name       string
		startEnd   StartEnd
		inputSlice []StartEnd
		expected   []StartEnd
	}
	tests := []test{
		{
			"append to empty slice",
			StartEnd{1, 2},
			[]StartEnd{},
			[]StartEnd{
				{1, 2},
			},
		},
		{
			"add to 1 record slice",
			StartEnd{1, 5},
			[]StartEnd{
				{2, 3},
			},
			[]StartEnd{
				{1, 1}, {2, 3}, {4, 5},
			},
		},
		{
			"accounted for in first index",
			StartEnd{2, 3},
			[]StartEnd{
				{2, 3}, {4, 5},
			},
			[]StartEnd{
				{2, 3}, {4, 5},
			},
		},
		{
			"extend left in current",
			StartEnd{1, 3},
			[]StartEnd{
				{2, 3}, {4, 5},
			},
			[]StartEnd{
				{1, 1}, {2, 3}, {4, 5},
			},
		},
		{
			"should insert in middle",
			StartEnd{1, 6},
			[]StartEnd{
				{3, 5}, {6, 8},
			},
			[]StartEnd{
				{1, 2}, {3, 5}, {6, 8},
			},
		},
		{
			"should insert in middle, start is greater",
			StartEnd{2, 6},
			[]StartEnd{
				{3, 5}, {7, 8},
			},
			[]StartEnd{
				{2, 2}, {3, 5}, {6, 6}, {7, 8},
			},
		},
		{
			"extends over",
			StartEnd{6, 7},
			[]StartEnd{
				{2, 5}, {6, 9},
			},
			[]StartEnd{
				{2, 5}, {6, 9},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RecursiveAppendStartEnd(test.startEnd, test.inputSlice)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("got %v\nwant %v\n", got, test.expected)
			}
		})
	}

}
