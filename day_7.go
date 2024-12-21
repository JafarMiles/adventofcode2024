package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent_of_code_day_7() { // part 2
	file, ferr := os.Open("day_7.in")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	runningTotal := 0
	for scanner.Scan() {
		line := scanner.Text()

		output := strings.Split(line, ": ")
		equationTotal, err := strconv.Atoi(output[0])
		if err != nil {
			panic(err)
		}
		var equationInputs []int
		for _, input := range strings.Split(output[1], " ") {
			equationInput, err2 := strconv.Atoi(input)
			if err2 != nil {
				panic(err2)
			}
			equationInputs = append(equationInputs, equationInput)
		}
		if exists_equation(equationTotal, equationInputs) {
			runningTotal += equationTotal
			fmt.Printf("equation exists for input: %s\n", line)
		} else {
			fmt.Printf("equation does not exist for input: %s\n", line)
		}
	}
	fmt.Printf("Total calibration result: %d\n", runningTotal)
}

func exists_equation(total int, inputs []int) bool {
	fmt.Printf("exists_equation total %d, inputs ", total)
	for _, val := range inputs {
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
	if total <= 0 {
		return false
	}
	if len(inputs) == 0 {
		return false
	}
	if len(inputs) == 1 {
		return total == inputs[0]
	}
	if exists_equation(total-inputs[len(inputs)-1], inputs[:len(inputs)-1]) {
		return true
	}
	if total%inputs[len(inputs)-1] == 0 && exists_equation(total/inputs[len(inputs)-1], inputs[:len(inputs)-1]) {
		return true
	}
	lastAsString := strconv.Itoa(inputs[len(inputs)-1])
	totalAsString := strconv.Itoa(total)
	if len(totalAsString) <= len(lastAsString) {
		return false
	}
	if totalAsString[len(totalAsString)-len(lastAsString):] != lastAsString {
		return false
	}

	newTotal, err := strconv.Atoi(totalAsString[:len(totalAsString)-len(lastAsString)])
	if err != nil {
		panic(err)
	}

	return exists_equation(newTotal, inputs[:len(inputs)-1])
}
