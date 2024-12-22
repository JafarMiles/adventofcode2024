package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Puzzle struct {
	ButtonA, ButtonB, Prize Coordinate
}

func advent_of_code_day_13() { // part 1
	file, ferr := os.Open("day_13.in")
	if ferr != nil {
		panic(ferr)
	}

	var puzzles []Puzzle
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		firstLineOfPuzzle := scanner.Text()
		buttonA := parse_puzzle_input_line(firstLineOfPuzzle, false)

		scanner.Scan()
		secondLineOfPuzzle := scanner.Text()
		buttonB := parse_puzzle_input_line(secondLineOfPuzzle, false)

		scanner.Scan()
		thirdLineOfPuzzle := scanner.Text()
		prize := parse_puzzle_input_line(thirdLineOfPuzzle, true)

		puzzles = append(puzzles, Puzzle{buttonA, buttonB, prize})

		scanner.Scan()
		scanner.Text() //Empty line
	}

	total := 0
	for _, puzzle := range puzzles {
		solution, ok := find_solution(puzzle)
		if !ok {
			continue // no solution found
		}
		total += solution.Row*3 + solution.Column
	}
	fmt.Printf("the total cost of getting the most prizes is %d\n", total)

	// parse input to get each equation - check
	// process each equation to find out if it is solvable
	// if it is solvable, with the restriction A,B <=100
	// find the solution, and compute the cost
	// sum up all the costs together
}

func find_solution(puzzle Puzzle) (Coordinate, bool) {
	// find a solution to the equation
	// 94*A + 22*B = 8400 -> A = (8400 - 22*B) / 94
	// 34*A + 67*B = 5400 -> 34 * (8400 - 22*B) / 94 +67*B = 5400
	// -> 34 * (8400 - 22*B) + 67*B*94 = 5400*94
	// -> B*(94*67 - 22*34) = 5400*94 -8400*34
	// -> B = (5400*94-3) / (94*67-22*34)
	// (507600-285600) / (6298-748)
	// 222000 / 5550 -> 22200 / 555 -> 200 / 5 -> B=40
	// => A = (8400 - 22*40) / 94 = (8400 - 880) / 94 = 7520 / 94 = 80
	// the solution is:
	// A=80
	// B=40

	// x1*a + y1*b = p1
	// x2*a + y2*b = p2
	// a = (p1 - y1*b) / x1
	// x2*(p1 - y1*b) / x1 + y2*b  = p2
	// x2*p1 - x2*y1*b + y2*x1*b = p2*x1
	// b = (p2*x1-x2*p1) / (y2*x1 - y1*x2)

	// a = (p1 - y1* (p2*x1-x2*p1) / (y2*x1 - y1*x2)) / x1
	// a = (p1 * (y2*x1 - y1*x2) - y1*(p2*x1-x2*p1)) / (x1 * (y2*x1 - y1*x2))
	// a = (p1*y2*x1 - p1*y1*x2 -y1*p2*x1 + y1*x2*p1) / (x1*(y2*x1-y1*x2))
	// a = (p1*y2*x1 - y1*p2*x1) / (x1*(y2*x1-y1*x2))

	// a = (p1*y2-y1*p2)/(y2*x1-y1*x2)

	// there is a unique solution if the denominator is not zero!
	p1 := puzzle.Prize.Row
	x1 := puzzle.ButtonA.Row
	y1 := puzzle.ButtonB.Row
	p2 := puzzle.Prize.Column
	x2 := puzzle.ButtonA.Column
	y2 := puzzle.ButtonB.Column
	denominator := y2*x1 - y1*x2
	if denominator != 0 {
		numeratorForA := p1*y2 - y1*p2
		numeratorForB := p2*x1 - x2*p1
		if numeratorForA%denominator == 0 && numeratorForB%denominator == 0 {
			// we have an integer solution!
			return Coordinate{numeratorForA / denominator, numeratorForB / denominator}, true
		}
		return Coordinate{}, false
	}
	// if denominator is 0, then buttons A and B will move the arm in the same direction
	// first check if the direction will reach the prize
	if p2*y1-p1*y2 != 0 {
		return Coordinate{}, false // not on the same line
	}
	if p2%y2 == 0 && p1%y1 == 0 {
		return Coordinate{0, p1 / y1}, true
	}
	if p2%x2 == 0 && p1%x1 == 0 {
		return Coordinate{p1 / x1, 0}, true
	}
	// if highest_common_factor of x1 and y1 divides p1, then there's an integer solution
	lcd := highest_common_factor(x1, y1)
	if p1%lcd == 0 {
		a, b, success := try_find_solution(x1, y1, p1)
		if success {
			return Coordinate{a, b}, true
		}
	}
	return Coordinate{}, false
}

func try_find_solution(x int, y int, p int) (int, int, bool) {
	if p < 0 {
		return 0, 0, false
	}
	if y%p == 0 {
		return 0, y / p, true
	}
	a, b, success := try_find_solution(x, y, p-x)
	if !success {
		return a, b, success
	}
	return a + 1, b, true
}

func highest_common_factor(x int, y int) int {
	if x == y {
		return x
	}
	if y > x {
		return highest_common_factor(y, x)
	}
	return highest_common_factor(x-y, y)
}

func parse_puzzle_input_line(input string, isPrize bool) Coordinate {
	var operation string
	if isPrize {
		operation = "="
	} else {
		operation = "+"
	}
	inputs := strings.Split(strings.Split(input, ":")[1], ",")
	axValue, err := strconv.Atoi(strings.Split(inputs[0], operation)[1])
	if err != nil {
		panic(err)
	}
	ayValue, err2 := strconv.Atoi(strings.Split(inputs[1], operation)[1])
	if err2 != nil {
		panic(err2)
	}
	if isPrize {
		return Coordinate{axValue + 10000000000000, ayValue + 10000000000000}
	}
	return Coordinate{axValue, ayValue}
}
