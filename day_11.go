package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent_of_code_day_11() { // part 1
	file, ferr := os.Open("day_11.in")
	if ferr != nil {
		panic(ferr)
	}

	var stones []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()

	stonesAsStrings := strings.Split(text, " ")
	for _, stoneAsString := range stonesAsStrings {
		fmt.Printf("initial stoneAsStrings %s\n", stoneAsString)
		stone, err := strconv.Atoi(stoneAsString)
		if err != nil {
			panic(err)
		}
		stones = append(stones, stone)
	}
	blinks := 75
	//totalStones := part_1(stones, blinks)

	totalStones := part_2(stones, blinks)

	fmt.Printf("number of stones after %d blinks is %d\n", blinks, totalStones)
}

func part_2(stones []int, blinks int) int {
	var cache map[Coordinate]int = make(map[Coordinate]int)
	totalStones := 0
	for i := 0; i < len(stones); i++ {
		totalStones += number_of_stones(stones[i], blinks, cache)
	}
	return totalStones
}

func part_1(stones []int, blinks int) int {
	for i := 0; i < blinks; i++ {
		fmt.Printf("stones after %d blinks is ", i+1)
		var stonesAfterBlink []int
		for j := 0; j < len(stones); j++ {
			stones_after_rule := apply_rules_to_stone(stones[j])
			for k := 0; k < len(stones_after_rule); k++ {
				stonesAfterBlink = append(stonesAfterBlink, stones_after_rule[k])
				fmt.Printf("%d ", stones_after_rule[k])
			}
		}
		fmt.Printf("\n")
		stones = stonesAfterBlink
	}
	return len(stones)
}

func number_of_stones(stone int, blinks int, cache map[Coordinate]int) int {
	//fmt.Printf("calling number_of_stones with %d, %d\n", stone, blinks)
	if blinks == 0 {
		return 1
	}
	cacheKey := Coordinate{stone, blinks}
	item, ok := cache[cacheKey]
	if ok {
		return item
	}

	if stone == 0 {
		result := number_of_stones(1, blinks-1, cache)
		cache[cacheKey] = result
		return result
	}
	number_of_digits := number_of_digits(stone)
	if number_of_digits%2 == 0 {
		first_half, second_half := split_number(stone, number_of_digits)
		result := number_of_stones(first_half, blinks-1, cache) + number_of_stones(second_half, blinks-1, cache)
		cache[cacheKey] = result
		return result
	}
	result2 := number_of_stones(stone*2024, blinks-1, cache)
	cache[cacheKey] = result2
	return result2
}

func split_number(number int, number_of_digits int) (int, int) {
	half_number_of_digits := number_of_digits / 2
	first_half := number
	for i := 0; i < half_number_of_digits; i++ {
		first_half = first_half / 10
	}
	diff := first_half
	for i := 0; i < half_number_of_digits; i++ {
		diff = diff * 10
	}
	//fmt.Printf("Rule 2: %d -> %d %d\n", number, first_half, number-diff)
	return first_half, number - diff
}

func number_of_digits(number int) int {
	if number < 0 {
		panic(errors.New("expected number >= 0"))
	}
	if number < 10 {
		return 1
	}
	return 1 + number_of_digits(number/10)
}

func apply_rules_to_stone(stone int) []int {
	var newStones []int

	if stone == 0 { //rule 1
		newStones = append(newStones, 1)
		return newStones
	}
	stoneAsString := strconv.Itoa(stone)
	if len(stoneAsString)%2 == 0 {
		length := len(stoneAsString) / 2
		firstNewStoneAsString := stoneAsString[:length]
		secondNewStoneAsString := stoneAsString[length:]
		//fmt.Printf("applying rule 2: %s -> %s, %s\n", stoneAsString, firstNewStoneAsString, secondNewStoneAsString)

		firstNewStone, err := strconv.Atoi(firstNewStoneAsString)
		if err != nil {
			panic(err)
		}
		newStones = append(newStones, firstNewStone)
		secondNewStone, err2 := strconv.Atoi(secondNewStoneAsString)
		if err2 != nil {
			panic(err2)
		}
		newStones = append(newStones, secondNewStone) //rule 2
		return newStones
	}
	newStones = append(newStones, 2024*stone) //rule 3
	return newStones
}
