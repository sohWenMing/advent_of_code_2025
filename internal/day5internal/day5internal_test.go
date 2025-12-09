package day5internal

import (
	"reflect"
	"testing"
)

func TestGetNumsAndStartEndsFromFile(t *testing.T) {
	nums, startEnds, err := GetNumsAndStartEndsFromFile("./testinput.txt")
	if err != nil {
		t.Errorf("didn't expect error, got %v\n", err)
		return
	}
	wantNums := []int64{1, 5, 8, 11, 17, 32}
	if !reflect.DeepEqual(nums, wantNums) {
		t.Errorf("got %v\nwant %v\n", nums, wantNums)
	}
	wantStartEnds := []StartEnd{
		{3, 5},
		{10, 14},
		{16, 20},
		{12, 18},
	}
	if !reflect.DeepEqual(startEnds, wantStartEnds) {
		t.Errorf("got %v\nwant %v\n", startEnds, wantStartEnds)
	}
}

func TestReadFileGetLines(t *testing.T) {
	freshLines, availableLines, err := ReadFileGetLines("./testinput.txt")
	if err != nil {
		t.Errorf("didn't expect error, got %v\n", err)
		return
	}
	freshLinesWant := 4
	if len(freshLines) != freshLinesWant {
		t.Errorf("got %d\nwant %d\n", len(freshLines), freshLinesWant)
	}
	availableLinesWant := 6
	if len(availableLines) != availableLinesWant {
		t.Errorf("got %d\nwant %d\n", len(availableLines), availableLinesWant)
	}
}

func TestParseStringToStartEnd(t *testing.T) {
	type test struct {
		name          string
		input         string
		expected      StartEnd
		isErrExpected bool
	}

	tests := []test{
		{
			"basic test should pass",
			"1-2",
			StartEnd{1, 2},
			false,
		},
		{
			"no dash, should fail",
			"12",
			StartEnd{},
			true,
		},
		{
			"empty string should fail",
			"",
			StartEnd{},
			true,
		},
		{
			"wrong range should fail",
			"2-1",
			StartEnd{},
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ParseStringToStartEnd(test.input)
			switch test.isErrExpected {
			case true:
				if err == nil {
					t.Errorf("expected err, didn't get one")
					return
				}
			default:
				if !reflect.DeepEqual(got, test.expected) {
					t.Errorf("got %v\nwant %v\n", got, test.expected)
					return
				}
			}
		})
	}
}
