package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {
	// result_part_1 := part1()
	result_part_2 := part2()
	// fmt.Println(result_part_1)
	fmt.Print(result_part_2)
}

func part1() int {
	result := 0
	r, _ := regexp.Compile("mul\\([0-9]{0,3},[0-9]{0,3}\\)")
	matches := r.FindAllString(input, -1)
	for i := range matches {
		result += multiply(matches[i])
	}
	return result
}

func part2() int {
	result := 0
	r, _ := regexp.Compile(`(mul\([0-9]{0,3},[0-9]{0,3}\))|(don\'t\(\))|(do\(\))`)
	matches := r.FindAllString(input, -1)
	is_disabled := false
	for i := range matches {
		if matches[i] == `don't()` {
			is_disabled = true
		} else if matches[i] == `do()` {
			is_disabled = false
		} else if !is_disabled {
			result += multiply(matches[i])
		}
	}
	return result
}

func multiply(input string) int {
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	matches := r.FindStringSubmatch(input)
	element_1, _ := strconv.Atoi(matches[1])
	element_2, _ := strconv.Atoi(matches[2])
	return element_1 * element_2
}
