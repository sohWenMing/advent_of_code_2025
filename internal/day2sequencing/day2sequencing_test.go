package day2sequencing

import (
	"reflect"
	"testing"
)

type test struct {
	name           string
	input          string
	isExpectRepeat bool
}

func TestIsRepeated(t *testing.T) {
	tests := []test{
		{
			"basic test should be repeated",
			"11",
			true,
		},
		{
			"longer sequence should pass",
			"1212",
			true,
		},
		{
			"longer sequence should fail",
			"1234",
			false,
		},
		{
			"odd characters should fail",
			"123",
			false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsRepeated(test.input)
			if got != test.isExpectRepeat {
				t.Errorf("got %v\nwant %v\n", got, test.isExpectRepeat)
			}
		})
	}
}

func TestGetNumRangeFromString(t *testing.T) {
	type getNumRangeTest struct {
		name          string
		input         string
		expected      []int
		isErrExpected bool
	}

	tests := []getNumRangeTest{
		{
			"basic test should pass",
			"1-2",
			[]int{1, 2},
			false,
		},
		{
			"no dash should fail",
			"11",
			[]int{},
			true,
		},
		{
			"empty string should fail",
			"",
			[]int{},
			true,
		},
		{
			"error input should fail 1",
			"11-",
			[]int{},
			true,
		},
		{
			"error input should fail 2",
			"-11",
			[]int{},
			true,
		},
		{
			"final test",
			"2-10",
			[]int{2, 3, 4, 5, 6, 7, 8, 9, 10},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetNumRangeFromString(test.input)
			switch test.isErrExpected {
			case true:
				if err == nil {
					t.Errorf("expected error, didn't get one")
					return
				}
			default:
				if err != nil {
					t.Errorf("didn't expect error, got %v\n", err)
					return
				}
				if !reflect.DeepEqual(got, test.expected) {
					t.Errorf("got %v\nwant %v\n", got, test.expected)
				}
			}
		})
	}
}

func TestRunSubStrings(t *testing.T) {
	tests := []test{
		{
			"basic test should find",
			"11",
			true,
		},
		{
			"basic test should not find",
			"12",
			false,
		},
		{
			"should find substring repeated in end",
			"122",
			true,
		},
		{
			"should find substring multiple chars front",
			"1212",
			true,
		},
		{
			"should find substring multiple chars middle",
			"121213",
			true,
		},
		{
			"should find substring multiple chars end",
			"12323",
			true,
		},
		{
			"check characters repeated not in sequence",
			"121",
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := RunSubStringsFromStringCheckRepeats(test.input)
			if got != test.isExpectRepeat {
				t.Errorf("got %v\nwant %v\n", got, test.isExpectRepeat)
			}
		})
	}
}

func TestCheckForRepeats(t *testing.T) {
	tests := []test{
		{
			"basic test should find",
			"11",
			true,
		},
		{
			"empty string should not find",
			"",
			false,
		},
		{
			"one character should not find",
			"1",
			false,
		},
		{
			"test 3 characters, repeat at front",
			"112",
			true,
		},
		{
			"test 4 characters, has repeat ",
			"1212",
			true,
		},
		{
			"test 4 characters, has no repeat ",
			"1234",
			false,
		},
		{
			"test 5 characters, has repeat ",
			"12122",
			true,
		},
		{
			"test 5 characters, has no repeat ",
			"12345",
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CheckForRepeats(test.input)
			if got != test.isExpectRepeat {
				t.Errorf("got %v\nwant %v\n", got, test.isExpectRepeat)
			}
		})
	}

}
