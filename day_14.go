package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	Position, Velocity Coordinate
}

func advent_of_code_day_14() { // part 1
	file, ferr := os.Open("day_14.in")
	if ferr != nil {
		panic(ferr)
	}

	var robots []Robot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs := strings.Split(line, " ")
		position := parse_coordinate(inputs[0])
		velocity := parse_coordinate(inputs[1])
		robots = append(robots, Robot{position, velocity})
	}
	maxCoordinate := Coordinate{103, 101}
	advent_of_code_day_14_part_1(robots, maxCoordinate)
	advent_of_code_day_14_part_2(robots, maxCoordinate)
}

func advent_of_code_day_14_part_2(robots []Robot, maxCoordinate Coordinate) {
	for i := 4000; i < 8000; i++ {
		mapOfLocations := make(map[Coordinate]int)
		for _, robot := range robots {
			location := predict_future_location(robot, maxCoordinate, i)
			item, ok := mapOfLocations[location]
			if !ok {
				mapOfLocations[location] = 1
			} else {
				mapOfLocations[location] = 1 + item
			}
		}

		count_of_isolated_robots := 0

		for location := range mapOfLocations {
			_, up := mapOfLocations[Coordinate{location.Row - 1, location.Column}]
			_, down := mapOfLocations[Coordinate{location.Row + 1, location.Column}]
			_, left := mapOfLocations[Coordinate{location.Row, location.Column + 1}]
			_, right := mapOfLocations[Coordinate{location.Row, location.Column - 1}]
			if !up && !down && !left && !right {
				count_of_isolated_robots++
			}
		}

		fmt.Printf("------- Iteration %d, has %d isolated robots------- ", i, count_of_isolated_robots)
		fmt.Printf("\n")
		if count_of_isolated_robots < 270 {
			print_robot_locations(mapOfLocations, maxCoordinate)
			fmt.Printf("------- Iteration end -------\n")
		}
	}
}

func print_robot_locations(mapOfLocations map[Coordinate]int, maxCoordinate Coordinate) {
	for i := 0; i < maxCoordinate.Row; i++ {
		for j := 0; j < maxCoordinate.Column; j++ {
			item, ok := mapOfLocations[Coordinate{i, j}]
			if ok {
				fmt.Printf("%d", item)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func advent_of_code_day_14_part_1(robots []Robot, maxCoordinate Coordinate) {
	midCoordinate := Coordinate{51, 50}

	tl := 0
	tr := 0
	bl := 0
	br := 0
	for _, robot := range robots {
		future_location := predict_future_location(robot, maxCoordinate, 100)
		if future_location.Row < midCoordinate.Row && future_location.Column < midCoordinate.Column {
			tl++
		}
		if future_location.Row > midCoordinate.Row && future_location.Column < midCoordinate.Column {
			tr++
		}
		if future_location.Row < midCoordinate.Row && future_location.Column > midCoordinate.Column {
			bl++
		}
		if future_location.Row > midCoordinate.Row && future_location.Column > midCoordinate.Column {
			br++
		}
	}
	fmt.Printf("the total safety score is %d * %d * %d * %d = %d\n", tl, tr, bl, br, tl*tr*bl*br)
}

func predict_future_location(robot Robot, maxCoordinate Coordinate, time int) Coordinate {
	future_row := modulo(robot.Position.Row+time*robot.Velocity.Row, maxCoordinate.Row)
	future_col := modulo(robot.Position.Column+time*robot.Velocity.Column, maxCoordinate.Column)
	return Coordinate{future_row, future_col}
}

func modulo(value int, base int) int {
	if value < 0 {
		return modulo(value+base, base)
	}
	if value >= base {
		return modulo(value-base, base)
	}
	return value
}

func parse_coordinate(input string) Coordinate {
	coordinates := strings.Split(strings.Split(input, "=")[1], ",")
	x, err := strconv.Atoi(coordinates[0])
	if err != nil {
		panic(err)
	}
	y, err2 := strconv.Atoi(coordinates[1])
	if err2 != nil {
		panic(err2)
	}
	return Coordinate{y, x}
}
