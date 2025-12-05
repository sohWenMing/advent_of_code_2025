package readfile

import (
	"reflect"
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
			DirectionAndCount{Right, 1},
			false,
		},
		{
			"basic test left should pass",
			"R1",
			DirectionAndCount{Right, 1},
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
