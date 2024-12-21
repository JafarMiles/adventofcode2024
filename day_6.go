package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type void struct{}

var member void

func advent_of_code_day_6() { // part 2
	file, ferr := os.Open("day_6.in")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)
	var mapWithObstacles []string

	guard_row := 0
	guard_column := 0
	guard_orientation := 0 //0 up, 1, right, 2 down, 3 left
	row_input := 0
	for scanner.Scan() {
		line := scanner.Text()

		if strings.ContainsRune(line, '^') {
			guard_row = row_input
			guard_column = strings.IndexRune(line, '^')
		}
		mapWithObstacles = append(mapWithObstacles, line)
		row_input++
	}

	runningTotal := 0

	previously_visited_squares := find_all_travelled_squares(mapWithObstacles, guard_row, guard_column, guard_orientation)

	for i := 0; i < len(mapWithObstacles); i++ {
		for j := 0; j < len(mapWithObstacles[0]); j++ {
			_, ok := previously_visited_squares[visited_str(i, j)]
			if !ok {
				continue // can skip if initial walk never encounters square
			}

			if mapWithObstacles[i][j] == '.' {
				// modify loaded map object to avoid additional memory allocations
				mapWithObstacles[i] = mapWithObstacles[i][:j] + "#" + mapWithObstacles[i][j+1:]

				if does_guard_loop(mapWithObstacles, guard_row, guard_column, guard_orientation) {
					fmt.Printf("X")
					runningTotal++
				} else {
					fmt.Printf("x")
				}
				// reset modifications on map
				mapWithObstacles[i] = mapWithObstacles[i][:j] + "." + mapWithObstacles[i][j+1:]
			}
		}
	}

	fmt.Printf("\nthere are %d squares which will cause a loop\n", runningTotal)
}

func make_copy_with_modification(initialMap []string, row int, column int) []string {
	var modifiedMapWithObstacles []string
	for k := 0; k < len(initialMap); k++ {
		modifiedMapWithObstacles = append(modifiedMapWithObstacles, initialMap[k])
	}
	newLine := modifiedMapWithObstacles[row][:column] + "#" + modifiedMapWithObstacles[row][column+1:]
	modifiedMapWithObstacles[row] = newLine
	fmt.Printf("modified map (%d, %d):\n", row, column)
	for _, line := range modifiedMapWithObstacles {
		fmt.Println(line)
	}

	return modifiedMapWithObstacles
}

func does_guard_loop(mapWithObstacles []string, guard_row int, guard_column int, guard_orientation int) bool {
	previously_visited_squares := make(map[string]void)
	previously_visited_squares[visited_with_direction_str(guard_row, guard_column, guard_orientation)] = member

	for {
		next_row, next_column := get_next_guard_position(guard_row, guard_column, guard_orientation)

		if !position_on_map(mapWithObstacles, next_row, next_column) {
			return false // end the guard has left the map!
		}
		if !position_is_blocked(mapWithObstacles, next_row, next_column) {
			guard_row = next_row
			guard_column = next_column
		} else {
			guard_orientation = (guard_orientation + 1) % 4
		}

		_, ok := previously_visited_squares[visited_with_direction_str(guard_row, guard_column, guard_orientation)]
		if ok {
			return true // we have a loop
		} else {
			previously_visited_squares[visited_with_direction_str(guard_row, guard_column, guard_orientation)] = member
		}
	}
}

func advent_of_code_day_6_part1() {
	file, ferr := os.Open("day_6.in")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)
	var mapWithObstacles []string

	guard_row := 0
	guard_column := 0
	guard_orientation := 0 //0 up, 1, right, 2 down, 3 left
	row_input := 0
	for scanner.Scan() {
		line := scanner.Text()

		if strings.ContainsRune(line, '^') {
			guard_row = row_input
			guard_column = strings.IndexRune(line, '^')
		}
		mapWithObstacles = append(mapWithObstacles, line)
		row_input++
	}
	previously_visited_squares := find_all_travelled_squares(mapWithObstacles, guard_row, guard_column, guard_orientation)

	fmt.Printf("the guard has travelled %d squares\n", len(previously_visited_squares))
}

func find_all_travelled_squares(mapWithObstacles []string, guard_row int, guard_column int, guard_orientation int) map[string]void {
	previously_visited_squares := make(map[string]void)
	previously_visited_squares[visited_str(guard_row, guard_column)] = member

	for {
		next_row, next_column := get_next_guard_position(guard_row, guard_column, guard_orientation)

		if !position_on_map(mapWithObstacles, next_row, next_column) {
			break // end the guard has left the map!
		}
		if !position_is_blocked(mapWithObstacles, next_row, next_column) {
			guard_row = next_row
			guard_column = next_column
			previously_visited_squares[visited_str(guard_row, guard_column)] = member
		} else {
			guard_orientation = (guard_orientation + 1) % 4
		}
	}
	return previously_visited_squares
}

func get_next_guard_position(row int, column int, orientation int) (int, int) {
	if orientation == 0 { //up
		return row - 1, column
	}
	if orientation == 1 { //right
		return row, column + 1
	}
	if orientation == 2 { //down
		return row + 1, column
	}
	return row, column - 1 //left
}

func position_is_blocked(mapWithObstacles []string, row int, column int) bool {
	return mapWithObstacles[row][column] == '#'
}

func position_on_map(mapWithObstacles []string, row int, column int) bool {
	if row < 0 {
		return false
	}
	if row >= len(mapWithObstacles) {
		return false
	}
	if column < 0 {
		return false
	}
	if column >= len(mapWithObstacles[0]) {
		return false
	}
	return true
}

func visited_str(row int, column int) string {
	return strconv.Itoa(row) + "," + strconv.Itoa(column)
}

func visited_with_direction_str(row int, column int, oritentation int) string {
	return strconv.Itoa(row) + "," + strconv.Itoa(column) + "," + strconv.Itoa(oritentation)
}
