package readfile

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestProcessLine(t *testing.T) {
	type test struct {
		name          string
		line          string
		expected      DirectionAndCount
		isErrExpected bool
	}
	tests := []test{
		{
			"basic test right should pass",
			"R1",
			DirectionAndCount{RIGHT, 1},
			false,
		},
		{
			"basic test left should pass",
			"R1",
			DirectionAndCount{RIGHT, 1},
			false,
		},
		{
			"test prefix number should fail",
			"11",
			DirectionAndCount{RIGHT, 1},
			true,
		},
		{
			"test no prefix should fail",
			"1",
			DirectionAndCount{RIGHT, 1},
			true,
		},
		{
			"test no number should fail",
			"R",
			DirectionAndCount{RIGHT, 1},
			true,
		},
		{
			"test number has characters should fail",
			"RRRRRR",
			DirectionAndCount{RIGHT, 1},
			false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := processLine(test.line)
			switch test.isErrExpected {
			case true:
				if err == nil {
					t.Errorf("didn't expect error, got %v\n", err)
					return
				}
			default:
				if err != nil {
					t.Errorf("didn't expect error, got error %v\n", err)
					return
				}
				if !reflect.DeepEqual(got, test.expected) {
					t.Errorf("got %s\n want %s\n", got.PrettyJSON(), test.expected.PrettyJSON())
				}
			}
		})
	}
}

func TestReadLines(t *testing.T) {
	// R29
	// R6
	// L43
	// L6
	// R28
	expected := []DirectionAndCount{
		{RIGHT, 29},
		{RIGHT, 6},
		{LEFT, 43},
		{LEFT, 6},
		{RIGHT, 28},
	}
	got, err := ReadFile("../../cmd/day1/input.txt", 5)
	if err != nil {
		t.Errorf("didn't expect error, got error %v\n", err)
		return
	}

	if !reflect.DeepEqual(got, expected) {
		var expectedBuilder strings.Builder
		var gotBuilder strings.Builder

		for _, record := range expected {
			expectedBuilder.WriteString(
				fmt.Sprintf("%s\n", record.PrettyJSON()),
			)
		}
		for _, record := range got {
			gotBuilder.WriteString(
				fmt.Sprintf("%s\n", record.PrettyJSON()),
			)
		}
		t.Errorf("got %s\nwant%s\n", gotBuilder.String(), expectedBuilder.String())
	}
}
