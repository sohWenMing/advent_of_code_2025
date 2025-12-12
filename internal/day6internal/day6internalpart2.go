package day6internal

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type numBytesToOp struct {
	numBytes [][]byte
	op       string
}

func (n *numBytesToOp) Operate() (result int, err error) {
	nums := []int{}
	curIdx := len(n.numBytes[0]) - 1
	var b bytes.Buffer
	for curIdx >= 0 {
		for _, slice := range n.numBytes {
			curByte := slice[curIdx]
			if curByte != ' ' {
				b.WriteByte(curByte)
			}
		}
		if len(b.Bytes()) > 0 {
			numToAppend, err := strconv.ParseInt(b.String(), 10, 0)
			if err != nil {
				return 0, fmt.Errorf("error parsing %s to int\n", b.String())
			} else {
				nums = append(nums, int(numToAppend))
			}
		}
		b.Reset()
		curIdx--
	}
	switch strings.TrimSpace(n.op) {
	case "+":
		initNum := 0
		for _, num := range nums {
			initNum += num
		}
		return initNum, nil
	default:
		initNum := 1
		for _, num := range nums {
			initNum = initNum * num
		}
		return initNum, nil
	}
}

func (n *numBytesToOp) repr() {
	fmt.Println("struct start")
	fmt.Println("")
	fmt.Println("======numBytes=========")
	for i, line := range n.numBytes {
		fmt.Printf("num %d in numBytes: %q\n", i, string(line))
	}
	fmt.Println("===============")
	fmt.Println("op: ", n.op)
	fmt.Println("struct end")
	fmt.Println("")
}

func Part2ReadFile(filePath string) (bytes []byte, err error) {
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

func GetLines(inputBytes []byte, sep []byte) [][]byte {
	slices := bytes.Split(inputBytes, sep)
	returnedLines := [][]byte{}
	for _, slice := range slices {
		if len(slice) > 0 {
			returnedLines = append(returnedLines, slice)
		}
	}
	return returnedLines
}

func checkIsBlank(lines [][]byte, idx int, lineLength int) bool {
	for _, line := range lines {
		if line[idx] != ' ' {
			return false
		} else {
			continue
		}
	}
	return true
}

func createNumBytesToOpFromLines(lines [][]byte, lastLineIdx int, beginSliceIdx int, curIdx int) numBytesToOp {
	initStruct := numBytesToOp{}
	initStruct.numBytes = [][]byte{}
	for j, line := range lines {
		if j == lastLineIdx {
			initStruct.op = strings.TrimSpace(string(line[beginSliceIdx:curIdx]))
		} else {
			initStruct.numBytes = append(initStruct.numBytes, line[beginSliceIdx:curIdx])
		}
	}
	return initStruct
}

func GetStructs(lines [][]byte) []numBytesToOp {
	returned := []numBytesToOp{}

	lastLineIdx := len(lines) - 1
	beginSliceIdx := 0

	lineLength := len(lines[0])
	for i := 0; i < lineLength; i++ {

		isBlank := checkIsBlank(lines, i, lineLength)

		if isBlank {
			numBytesToOp := createNumBytesToOpFromLines(lines, lastLineIdx, beginSliceIdx, i)
			returned = append(returned, numBytesToOp)
			beginSliceIdx = i + 1
			continue
		} else if i == lineLength-1 {
			initStruct := numBytesToOp{}
			initStruct.numBytes = [][]byte{}
			for j, line := range lines {
				if j == lastLineIdx {
					initStruct.op = strings.TrimSpace(string(line[beginSliceIdx:]))
				} else {
					initStruct.numBytes = append(initStruct.numBytes, line[beginSliceIdx:])
				}
			}
			returned = append(returned, initStruct)
		}
	}
	return returned
}

func getOperatorLine(inputBytes []byte, seperator string) ([]byte, error) {
	sep := []byte(seperator)
	sepIdx := bytes.Index(inputBytes, sep)
	if sepIdx == -1 {
		return nil, fmt.Errorf("sep %s cannot be found.", seperator)
	}
	return inputBytes[sepIdx:], nil
}

type accumulator struct {
	buf bytes.Buffer
}

func getOpSegments(inputBytes []byte) [][]byte {
	returned := [][]byte{}
	reader := bytes.NewReader(inputBytes)
	acc := &accumulator{}
	for {
		fmt.Printf("buf string: %q\n", acc.buf.String())
		readByte, err := reader.ReadByte()
		if err != nil {
			returned = handleEOFErr(acc, returned)
			printReturned(returned)
			break
		}
		if readByte == '\n' {
			continue
		}
		if isByteNotSpace(readByte) {
			returned = acc.WriteNonSpace(readByte, returned)
			printReturned(returned)
			continue
		} else {
			acc.buf.WriteByte(readByte)
			continue
		}
	}
	return returned
}

func printReturned(input [][]byte) {
	for i, slice := range input {
		fmt.Printf("slice %d in printReturned %q\n", i, string(slice))
	}
}

func (a *accumulator) WriteNonSpace(readByte byte, curReturned [][]byte) (returned [][]byte) {
	isBufFilled := len(a.buf.Bytes()) != 0
	if isBufFilled {
		sliceCopy := append([]byte(nil), a.buf.Bytes()...)
		curReturned = append(curReturned, sliceCopy)
		printReturned(curReturned)
		a.buf.Reset()
	}
	a.buf.WriteByte(readByte)
	return curReturned
}

func handleEOFErr(acc *accumulator, curReturned [][]byte) (returned [][]byte) {
	if len(acc.buf.Bytes()) > 0 {
		appendBytes := acc.buf.Bytes()
		curReturned = append(curReturned, appendBytes)
	}
	return curReturned
}
func isByteNotSpace(inputByte byte) bool {
	isByteNotSpace := inputByte != ' '
	return isByteNotSpace
}
