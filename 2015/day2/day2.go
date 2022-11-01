package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	input := loadInput()
	totalRibbonLen := 0
	for _, dims := range input {
		pos, smallest1 := minVal(dims...)
	}
}

func Part1() int {
	input := loadInput()
	totalWrapPaper := 0
	for _, dims := range input {
		areaSide1 := dims[0] * dims[1]
		areaSide2 := dims[1] * dims[2]
		areaSide3 := dims[0] * dims[2]
		wrapPaper := 2*areaSide1 + 2*areaSide2 + 2*areaSide3
		_, slack := minVal(areaSide1, areaSide2, areaSide3)
		totalWrapPaper += wrapPaper + slack
	}
	return totalWrapPaper
}

func loadInput() [][]int {
	var input [][]int
	// Path of file is relative to main.go as that is where this function is used
	file, err := os.Open("day2/input.txt")
	errCheck(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strDims := strings.Split(line, "x")
		var intDims []int
		for _, dim := range strDims {
			intDim, err := strconv.Atoi(dim)
			errCheck(err)
			intDims = append(intDims, intDim)
		}
		input = append(input, intDims)
	}
	return input
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func minVal(input ...int) (int, int) {
	min := input[0]
	minPos := 0
	for pos, value := range input {
		if value < min {
			min = value
			minPos = pos
		}
	}
	return minPos, min
}
