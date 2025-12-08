package day3internal

import (
	"reflect"
	"testing"
)

func TestGetLargestPossibleNumWithRemaining(t *testing.T) {
	type test struct {
		name          string
		input         string
		minChars      int
		expected      string
		isErrExpected bool
	}
	tests := []test{
		{
			"basic test should pass",
			"987654321111111",
			12,
			"987654321111",
			false,
		},
		{
			"basic test 2 should pass",
			"811111111111119",
			12,
			"811111111119",
			false,
		},
		{
			"basic test 2 should pass",
			"234234234234278",
			12,
			"434234234278",
			false,
		},
		{
			"basic test 2 should pass",
			"818181911112111",
			12,
			"888911112111",
			false,
		},
	}
	for _, test := range tests {
		got, err := GetLargestPossibleNumWithRemainingChars(test.input, test.minChars)
		switch test.isErrExpected {
		case true:
			if err == nil {
				t.Errorf("expected error, didn't get one")
				return
			}
		default:
			if err != nil {
				t.Errorf("didn't expect error, got %v\n", err)
			}
			if got != test.expected {
				t.Errorf("got %s\nwant %s\n", got, test.expected)
			}
		}
	}
}

func TestGetLargestPossible(t *testing.T) {
	type test struct {
		name          string
		input         string
		expected      int
		isErrExpected bool
	}

	tests := []test{
		{
			"basic test should pass",
			"1234",
			34,
			false,
		},
		{
			"testing ends",
			"65555556",
			66,
			false,
		},
		{
			"empty string should error",
			"",
			0,
			true,
		},
		{
			"one char, should error",
			"1",
			0,
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetLargestPossibleNumFromString(test.input)
			switch test.isErrExpected {
			case true:
				if err == nil {
					t.Errorf("expected error, didn't get one")
					return
				}
			default:
				if err != nil {
					t.Errorf("didn't expect error, got %v\n", err)
				}
				if got != test.expected {
					t.Errorf("got %d\nwant %d\n", got, test.expected)
				}
			}
		})
	}
}

func TestGetLargestFromString(t *testing.T) {
	type test struct {
		name          string
		input         string
		expected      int
		isErrExpected bool
	}
	tests := []test{
		{
			"basic test should pass",
			"1234",
			4,
			false,
		},
		{
			"reverse test ",
			"4321",
			4,
			false,
		},
	}
	for _, test := range tests {
		got, err := GetLargestFromString(test.input)
		switch test.isErrExpected {
		case true:
			if err == nil {
				t.Errorf("expected error, didn't get one")
				return
			}
		default:
			if got != test.expected {
				t.Errorf("got %d\nwant %d\n", got, test.expected)
			}
		}
	}
}

func TestGetFirstLargestFromInput(t *testing.T) {
	type test struct {
		name          string
		input         string
		expected      NumToPos
		isErrExpected bool
	}
	tests := []test{
		{
			"basic test should pass",
			"1234",
			NumToPos{3, 2},
			false,
		},
		{
			"empty string should fail",
			"",
			NumToPos{},
			true,
		},
		{
			"test multiple repeat largest should get first",
			"33333321",
			NumToPos{3, 0},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetFirstLargestFromInput(test.input)
			switch test.isErrExpected {
			case true:
				if err == nil {
					t.Errorf("expected error, didn't get one")
					return
				}
			default:
				if !reflect.DeepEqual(got, test.expected) {
					t.Errorf("got %s\nwant %s\n",
						PrettyJson(got), PrettyJson(test.expected))
				}
			}
		})
	}
}
