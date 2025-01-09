package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {
	result_part_1 := part1()
	// result_part_2 := part2()
	fmt.Println(result_part_1)
	// fmt.Println(result_part_2)

}

func part1() int {
	var result int
	parsed_input := parseInput(input)
	for i := 0; i < 14; i++ {
		parsed_input = blink(parsed_input)
		fmt.Println(parsed_input, len(parsed_input))

	}
	result = len(parsed_input)
	return result
}

func part2() int {
	res := 0

	return res
}

func parseInput(input string) (parsed_input []int) {
	rows := strings.Split(input, " ")
	for i := range rows {
		element, _ := strconv.Atoi(rows[i])
		parsed_input = append(parsed_input, element)
	}
	return parsed_input
}

func blink(input []int) []int {
	i := 0
	for i < len(input) {
		rule := get_rule(input[i])
		// fmt.Println(i, rule, input)
		switch rule {
		case 1:
			input[i] = 1
		case 2:
			digits := getAmountOfDigits(input[i], 1)
			mid := digits / 2
			str_element := fmt.Sprint(input[i])
			rune_element := []rune(str_element)
			str_left_element := string(rune_element[:mid])
			str_right_element := string(rune_element[mid:])
			left_element, _ := strconv.Atoi(str_left_element)
			right_element, _ := strconv.Atoi(str_right_element)

			right_side := make([]int, len(input[i+1:]))
			copy(right_side, input[i+1:])
			input = append(input[:i], left_element)
			input = append(input, right_element)
			input = append(input, right_side...)
			i++
		case 3:
			input[i] *= 2024
		}
		i++
	}
	return input
}

func get_rule(element int) int {
	if element == 0 {
		return 1
	}
	if getAmountOfDigits(element, 1)%2 == 0 {
		return 2
	}
	return 3
}

func getAmountOfDigits(num int, digits int) int {
	if num >= 10 {
		return getAmountOfDigits(num/10, digits+1)
	} else {
		return digits
	}
}
