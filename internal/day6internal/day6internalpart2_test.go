package day6internal

import (
	"testing"
)

func TestPart2(t *testing.T) {
	bytes, err := part2ReadFile("testinput.txt")
	if err != nil {
		t.Errorf("didn't expect error, got %v\n", err)
		return
	}
	lines := getLines(bytes, []byte("\n"))
	records := getStructs(lines)
	result := 0
	for _, record := range records {
		operationResult, err := record.operate()
		if err != nil {
			t.Errorf("didn't expect error, got %v\n", err)
			return
		}
		result += operationResult
	}
	if result != 3263827 {
		t.Errorf("didn't get expected")
	}
}

func TestOperate(t *testing.T) {
	byteSlice := [][]byte{}
	byteSlice = append(byteSlice, []byte("64 "))
	byteSlice = append(byteSlice, []byte("23 "))
	byteSlice = append(byteSlice, []byte("314"))

	testInput := numBytesToOp{
		byteSlice, "+",
	}

	result, err := testInput.operate()
	if err != nil {
		t.Errorf("didn't expect error, got %v\n", err)
		return
	}
	if result != 1058 {
		t.Errorf("didn't get expected")
		return
	}
}
