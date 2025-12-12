package day6internal

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	numBytes, opBytes, err := GetSlicesFromFile("./testinput.txt", "*")
	if err != nil {
		t.Errorf("didnt' expect error, got err %v\n", err)
		return
	}
	opSlice := GetNonSpaceByteSlice(opBytes)
	numSlices, err := GetNumsFromBytes(numBytes, len(opSlice))
	if err != nil {
		t.Errorf("didnt' expect error, got err %v\n", err)
		return
	}
	fmt.Println("numBytesSlices: ", numSlices)

	operatorSlices, err := GetOperatorSlices(opSlice)
	if err != nil {
		t.Errorf("didnt' expect error, got err %v\n", err)
		return
	}
	fmt.Println("operatorSlices: ", operatorSlices)

	idxToNumToOpMap := getMap(numSlices, operatorSlices)
	fmt.Println("idxToNumsToOpMap: ", idxToNumToOpMap)

	sum := 0
	for _, numsToOp := range idxToNumToOpMap {
		toAdd := numsToOp.RunOperation()
		sum += toAdd
	}
	fmt.Println("sum: ", sum)

}
