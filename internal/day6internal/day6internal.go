package day6internal

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type accumlator struct {
	isAccumlating bool
	buf           bytes.Buffer
}

type accumulatorResult struct {
	numToWrite int
	isWrite    bool
	err        error
}

const (
	Add = iota
	Subtract
	Divide
	Multiply
)

type numsToOp struct {
	nums []int
	op   int
}

func Part2(filePath string, sep string) (int, error) {
	idxToNumToOpMap, err := getIdxtoNumtoOpMap(filePath, sep)
	if err != nil {
		return 0, err
	}

	workingVertNumtoOps := []numsToOp{}

	for _, numToOp := range idxToNumToOpMap {
		verInts, err := numToOp.mapNumsToVertInts()
		if err != nil {
			return 0, err
		}
		workingNumToOp := numsToOp{
			verInts,
			numToOp.op,
		}
		workingVertNumtoOps = append(workingVertNumtoOps, workingNumToOp)
	}
	fmt.Println("workingVertNumToOps: ", workingVertNumtoOps)

	sum := 0
	for _, workingVertNumtoOp := range workingVertNumtoOps {
		sum += workingVertNumtoOp.RunOperation()
	}
	return sum, nil
}

func (n *numsToOp) mapNumsToVertInts() (returned []int, err error) {
	numStrings := getNumStringsFromIntSlice(n.nums)

	//set current length first to the length of the longest string
	curLength := getLongestStringLength(numStrings)

	// start at the longest, keep going until curLength > 0
	returned = []int{}

	for curLength > 0 {
		var b strings.Builder
		for _, numString := range numStrings {
			curIdx := curLength - 1
			if len(numString) >= curLength {
				b.WriteByte(numString[curIdx])
			}
		}
		accumulatedString := b.String()
		if len(accumulatedString) > 0 {
			numFromString, err := strconv.ParseInt(accumulatedString, 10, 0)
			if err != nil {
				return nil, err
			} else {
				returned = append(returned, int(numFromString))
			}
		}
		curLength -= 1
	}
	return returned, nil

}

func getLongestStringLength(input []string) int {
	longest := 0
	for _, numString := range input {
		length := len(numString)
		if length > longest {
			longest = length
		}
	}
	return longest
}

func getNumStringsFromIntSlice(input []int) []string {
	returned := make([]string, len(input))
	for i, num := range input {
		returned[i] = fmt.Sprintf("%d", num)
	}
	return returned
}

func Part1(filePath string, sep string) (int, error) {
	idxToNumToOpMap, err := getIdxtoNumtoOpMap(filePath, sep)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, numsToOp := range idxToNumToOpMap {
		toAdd := numsToOp.RunOperation()
		sum += toAdd
	}
	return sum, nil
}

func getIdxtoNumtoOpMap(filePath string, sep string) (idxToNumToOpMap, error) {
	numBytes, opBytes, err := GetSlicesFromFile(filePath, sep)
	if err != nil {
		return nil, err
	}
	opSlice := GetNonSpaceByteSlice(opBytes)
	numSlices, err := GetNumsFromBytes(numBytes, len(opSlice))
	if err != nil {
		return nil, err
	}
	operatorSlices, err := GetOperatorSlices(opSlice)
	if err != nil {
		return nil, err
	}
	idxToNumToOpMap := getMap(numSlices, operatorSlices)
	return idxToNumToOpMap, nil
}

type idxToNumToOpMap map[int]numsToOp

func (nTo *numsToOp) RunOperation() int {
	sum := 0
	switch nTo.op {
	case Add:
		for _, num := range nTo.nums {
			sum += num
		}
		return sum
	case Subtract:
		for _, num := range nTo.nums {
			sum -= num
		}
		return sum
	case Divide:
		for _, num := range nTo.nums {
			if sum == 0 {
				continue
			} else {
				sum = sum / num
			}
		}
		return sum
	case Multiply:
		for i, num := range nTo.nums {
			if i == 0 {
				sum = num
				continue
			}
			sum = sum * num
		}
		return sum
	default:
		return -1
	}
}

func getMap(numSlices [][]int, opEnums []int) idxToNumToOpMap {
	returned := make(idxToNumToOpMap)
	for _, slice := range numSlices {
		for i, num := range slice {
			val, ok := returned[i]
			if !ok {
				val = numsToOp{
					nums: []int{},
					op:   i,
				}
			}
			val.nums = append(val.nums, num)
			val.op = opEnums[i]
			returned[i] = val
		}
	}
	return returned
}

func GetOperatorSlices(opByteSlice []byte) (opEnums []int, err error) {
	returned := make([]int, len(opByteSlice))
	for i, byte := range opByteSlice {
		switch string(byte) {
		case "+":
			returned[i] = Add
		case "-":
			returned[i] = Subtract
		case "/":
			returned[i] = Divide
		case "*":
			returned[i] = Multiply
		default:
			return nil, fmt.Errorf("invalid operator: %s", string(byte))
		}
	}
	return returned, nil
}

func GetNumsFromBytes(input []byte, sliceLen int) ([][]int, error) {

	returned := [][]int{}
	curSlice := []int{}

	acc := &accumlator{}
	for _, byte := range input {
		accResult := acc.accumulate(byte)
		if accResult.err != nil {
			return nil, accResult.err
		}
		if accResult.isWrite {
			curSlice = append(curSlice, accResult.numToWrite)
			if len(curSlice) == sliceLen {
				returned = append(returned, curSlice)
				curSlice = []int{}
			}
		}
	}
	if acc.isAccumlating {
		num, err := strconv.ParseInt(string(acc.buf.String()), 0, 10)
		if err != nil {
			return nil, err
		}
		curSlice := append(curSlice, int(num))
		returned = append(returned, curSlice)
	}
	for _, slice := range returned {
		if len(slice) != sliceLen {
			return nil, fmt.Errorf("slice is of wrong length. slice %v\n len %d\n", slice, sliceLen)
		}
	}
	return returned, nil
}

func (a *accumlator) accumulate(inByte byte) (result accumulatorResult) {
	if !isByteNum(inByte) {
		switch a.isAccumlating {
		case false:
			return accumulatorResult{0, false, nil}
		default:
			num, err := strconv.ParseInt(a.buf.String(), 10, 0)
			if err != nil {
				return accumulatorResult{0, false, err}
			} else {
				a.buf.Reset()
				a.isAccumlating = false
				return accumulatorResult{int(num), true, nil}
			}
		}
	} else {
		a.isAccumlating = true
		a.buf.WriteByte(inByte)
		return accumulatorResult{0, false, nil}
	}
}

func isByteNum(inByte byte) bool {
	eval := inByte >= '0' && inByte <= '9'
	return eval
}

func GetSlicesFromFile(filePath string, sep string) (numBytes []byte, opBytes []byte, err error) {
	bytes, err := readFile(filePath)
	if err != nil {
		return nil, nil, err
	}
	sepIdx := getSepIdx(bytes, sep)
	if sepIdx == -1 {
		return nil, nil, fmt.Errorf("seperator: %s not found in bytes\n", sep)
	}
	return bytes[:sepIdx], bytes[sepIdx:], nil
}

func readFile(filePath string) (bytes []byte, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err = io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func getSepIdx(inputBytes []byte, sep string) (idx int) {
	sepBytes := []byte(sep)
	idx = bytes.Index(inputBytes, sepBytes)
	return idx
}

func getBytesFromSep(inputBytes []byte, sep string) []byte {
	sepIdx := getSepIdx(inputBytes, sep)
	if sepIdx == -1 {
		return inputBytes
	}
	return inputBytes[sepIdx:]
}

func GetNonSpaceByteSlice(inputBytes []byte) []byte {
	n := 0
	for _, byte := range inputBytes {
		if strings.TrimSpace(string(byte)) != "" {
			inputBytes[n] = byte
			n++
		}
	}
	return inputBytes[:n]
}
