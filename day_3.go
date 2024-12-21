package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func advent_of_code_day_3() {
	file, ferr := os.Open("day_3_input.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	text := ""
	for scanner.Scan() {
		text += scanner.Text()
	}

	textChunks := strings.Split(text, "don't()")

	runningTotal := parseAllMulCommands(textChunks[0])
	for i := 1; i < len(textChunks); i++ {
		textChunk := textChunks[i]
		chunks := strings.Split(textChunk, "do()")
		for j := 1; j < len(chunks); j++ {
			runningTotal += parseAllMulCommands(chunks[j])
		}
	}

	fmt.Printf("Total of results: %d\n", runningTotal)
}

func parseAllMulCommands(line string) int {
	statementRegex, err := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	if err != nil {
		panic(err)
	}

	intRegex, intErr := regexp.Compile("[0-9]{1,3}")
	if intErr != nil {
		panic(intErr)
	}

	statements := statementRegex.FindAllString(line, -1)
	innerLoopTotal := 0
	for i := 0; i < len(statements); i++ {
		numbers := intRegex.FindAllString(statements[i], -1)
		firstNum, err1 := strconv.Atoi(numbers[0])
		if err1 != nil {
			panic(err1)
		}
		secondNum, err2 := strconv.Atoi(numbers[1])
		if err2 != nil {
			panic(err2)
		}

		result := firstNum * secondNum
		fmt.Printf("found uncorrupted statement %s parsed %d * %d = %d and running total %d\n", statements[i], firstNum, secondNum, result, innerLoopTotal)
		innerLoopTotal += result
	}
	return innerLoopTotal
}
