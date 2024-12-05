package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

//go:embed test2.txt
var test2 string

var N string = "N"
var S string = "S"
var E string = "E"
var W string = "W"
var NE string = "NE"
var NW string = "NW"
var SE string = "SE"
var SW string = "SW"

func main() {
	// result_part_1 := part1()
	result_part_2 := part2()
	// fmt.Println(result_part_1)
	fmt.Println(result_part_2)
}

func part1() int {
	result := 0
	parsedInput := parseInput(input)
	for i := range parsedInput {
		runeSample := []rune(parsedInput[i])
		for j := range parsedInput[i] {
			if string(runeSample[j]) == "X" {
				count_xmas := searchXmasAllDir(parsedInput, i, j)
				result += count_xmas
			}
		}
	}
	return result
}

func part2() int {
	result := 0
	parsedInput := parseInput(input)
	for i := range parsedInput {
		runeSample := []rune(parsedInput[i])
		for j := range parsedInput[i] {
			if string(runeSample[j]) == "A" {
				count_xmas := searchXmasPart2(parsedInput, i, j)
				result += count_xmas
			}
		}
	}
	return result
}

func parseInput(input string) (parsedList []string) {
	return strings.Split(input, "\n")
}

func searchXmasAllDir(input []string, i int, j int) int {
	return searchXmasOneDir(input, i, j, "M", N) +
		searchXmasOneDir(input, i, j, "M", S) +
		searchXmasOneDir(input, i, j, "M", E) +
		searchXmasOneDir(input, i, j, "M", W) +
		searchXmasOneDir(input, i, j, "M", NE) +
		searchXmasOneDir(input, i, j, "M", NW) +
		searchXmasOneDir(input, i, j, "M", SE) +
		searchXmasOneDir(input, i, j, "M", SW)
}

func searchXmasOneDir(input []string, i int, j int, next_letter string, dir string) int {
	// This code should be much better
	switch dir {
	case N:
		if i == 0 {
			return 0
		}
		runeSample := []rune(input[i-1])
		if next_letter == "S" && string(runeSample[j]) == "S" {
			return 1
		}
		if next_letter == string(runeSample[j]) {
			return searchXmasOneDir(input, i-1, j, getNextLetter(next_letter), dir)
		} else {
			return 0
		}
	case S:
		if i == len(input)-1 {
			return 0
		}
		runeSample := []rune(input[i+1])
		if next_letter == "S" && string(runeSample[j]) == "S" {
			return 1
		}
		if next_letter == string(runeSample[j]) {
			return searchXmasOneDir(input, i+1, j, getNextLetter(next_letter), dir)
		} else {
			return 0
		}
	case E:
		runeSample := []rune(input[i])
		if j == len(runeSample)-1 {
			return 0
		}
		if next_letter == "S" && string(runeSample[j+1]) == "S" {
			return 1
		}
		if next_letter == string(runeSample[j+1]) {
			return searchXmasOneDir(input, i, j+1, getNextLetter(next_letter), dir)
		} else {
			return 0
		}
	case W:
		if j == 0 {
			return 0
		}
		runeSample := []rune(input[i])
		if next_letter == "S" && string(runeSample[j-1]) == "S" {
			return 1
		}
		if next_letter == string(runeSample[j-1]) {
			return searchXmasOneDir(input, i, j-1, getNextLetter(next_letter), dir)
		} else {
			return 0
		}
	case NE:
		if i == 0 {
			return 0
		}
		runeSample := []rune(input[i-1])
		if j == len(runeSample)-1 {
			return 0
		}
		if next_letter == "S" && string(runeSample[j+1]) == "S" {
			return 1
		}
		if next_letter == string(runeSample[j+1]) {
			return searchXmasOneDir(input, i-1, j+1, getNextLetter(next_letter), dir)
		} else {
			return 0
		}
	case NW:
		if i == 0 || j == 0 {
			return 0
		}
		runeSample := []rune(input[i-1])
		if next_letter == "S" && string(runeSample[j-1]) == "S" {
			return 1
		}
		if next_letter == string(runeSample[j-1]) {
			return searchXmasOneDir(input, i-1, j-1, getNextLetter(next_letter), dir)
		} else {
			return 0
		}
	case SE:
		if i == len(input)-1 {
			return 0
		}
		runeSample := []rune(input[i+1])
		if j == len(runeSample)-1 {
			return 0
		}
		if next_letter == "S" && string(runeSample[j+1]) == "S" {
			return 1
		}
		if next_letter == string(runeSample[j+1]) {
			return searchXmasOneDir(input, i+1, j+1, getNextLetter(next_letter), dir)
		} else {
			return 0
		}
	case SW:
		if i == len(input)-1 || j == 0 {
			return 0
		}
		runeSample := []rune(input[i+1])
		if next_letter == "S" && string(runeSample[j-1]) == "S" {
			return 1
		}
		if next_letter == string(runeSample[j-1]) {
			return searchXmasOneDir(input, i+1, j-1, getNextLetter(next_letter), dir)
		} else {
			return 0
		}
	}

	fmt.Println("SOMETHING WENT WRONG 1")
	return 0
}

func getNextLetter(letter string) string {
	switch letter {
	case "X":
		return "M"
	case "M":
		return "A"
	case "A":
		return "S"
	case "S":
		fmt.Println("SOMETHING WENT WRONG 2")
	}
	return ""
}

func searchXmasPart2(input []string, i int, j int) int {
	runeSample := []rune(input[i])
	if i == 0 || i == len(input)-1 || j == 0 || j == len(runeSample)-1 {
		return 0
	}
	runeSampleN := []rune(input[i+1])
	runeSampleS := []rune(input[i-1])

	diagon_1 := (string(runeSampleN[j-1]) == "M" && string(runeSampleS[j+1]) == "S") || (string(runeSampleN[j-1]) == "S" && string(runeSampleS[j+1]) == "M")
	diagon_2 := (string(runeSampleN[j+1]) == "M" && string(runeSampleS[j-1]) == "S") || (string(runeSampleN[j+1]) == "S" && string(runeSampleS[j-1]) == "M")
	if diagon_1 && diagon_2 {
		return 1
	}
	return 0
}
