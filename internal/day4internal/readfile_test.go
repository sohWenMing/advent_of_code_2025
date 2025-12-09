package day4internal

import (
	"reflect"
	"slices"
	"testing"
)

func TestIntegration(t *testing.T) {
	numCanMove, err := ReadFileGetMovable("./testFile4.txt")
	if err != nil {
		t.Errorf("didn't expect error, got err")
		return
	}

	if numCanMove != 13 {
		t.Errorf("got %d\nwant %d\n", numCanMove, 13)
	}
}

func TestGetCellMapFromFile(t *testing.T) {
	type test struct {
		name         string
		testFilePath string
		expected     []string
	}

	tests := []test{
		{
			"basic test, all should read in",
			"./testFile1.txt",
			[]string{
				"0|0", "1|0", "2|0", "3|0",
				"0|1", "1|1", "2|1", "3|1",
				"0|2", "1|2", "2|2", "3|2",
				"0|3", "1|3", "2|3", "3|3",
			},
		},
		{
			"empty test",
			"./testFile2.txt",
			[]string{},
		},
		{
			"one cell test",
			"./testFile3.txt",
			[]string{"2|2"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := []string{}
			cellMap, _, err := GetCellsValsFromFile(test.testFilePath)
			if err != nil {
				t.Errorf("didn't expect error, got err")
			}
			for key, _ := range cellMap {
				got = append(got, key)
			}
			slices.Sort(got)
			slices.Sort(test.expected)
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("got %v\nwant %v\n", got, test.expected)
				return
			}
		})
	}
}

func TestGetAllCells(t *testing.T) {
	type test struct {
		name             string
		testFilePath     string
		expectedNumCells int
	}

	tests := []test{
		{
			"basic test, 16 cells",
			"./testFile1.txt",
			16,
		},
	}
	for _, test := range tests {
		_, got, err := GetCellsValsFromFile(test.testFilePath)
		if err != nil {
			t.Errorf("didn't expect error, got err")
		}
		if len(got) != test.expectedNumCells {
			t.Errorf("got %d\nwant %d\n", len(got), test.expectedNumCells)
		}
	}
}

func TestGetSurroundCells(t *testing.T) {
	cell := Cell{1, 1}
	surroundingCells := GetSurroundingCells(cell)
	want := []Cell{
		{0, 0},
		{1, 0},
		{2, 0},
		{0, 1},
		{2, 1},
		{0, 2},
		{1, 2},
		{2, 0},
	}
	if !reflect.DeepEqual(surroundingCells, want) {
		t.Errorf("got %v\nwant %v\n", surroundingCells, want)
	}
}
