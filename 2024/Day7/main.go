package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {
	// result_part_1 := part1()
	result_part_2 := part2()
	// fmt.Println(result_part_1)
	fmt.Println(result_part_2)
}

func part1() int {
	res := 0
	equations := parseInput(input)
	for i := range equations {
		result, equation := equations[i][0], equations[i][2:]
		exists_equation := resolveEquation(result, equation, 0)
		if exists_equation {
			res += result
		}
	}
	return res
}

func part2() int {
	res := 0
	equations := parseInput(input)
	for i := range equations {
		result, equation := equations[i][0], equations[i][2:]
		exists_equation := resolveEquationPart2(result, equation, 0)
		if exists_equation {
			res += result
		}
	}
	return res
}

func parseInput(input string) (parsed_input [][]int) {
	rows := strings.Split(input, "\n")

	for i := range rows {
		aux := strings.Split(rows[i], ":")
		parameters := strings.Split(aux[1], " ")
		result, _ := strconv.Atoi(aux[0])
		parsed_parameters := []int{result}

		for j := range parameters {
			parsed_parameter, _ := strconv.Atoi(parameters[j])
			parsed_parameters = append(parsed_parameters, parsed_parameter)
		}
		parsed_input = append(parsed_input, parsed_parameters)
	}
	return parsed_input
}

func resolveEquation(result int, equation []int, preliminary_result int) bool {
	if preliminary_result == result {
		return true
	}
	if preliminary_result > result {
		return false
	}
	if len(equation) == 0 {
		return false
	}
	var new_pre_mult_result int
	if preliminary_result == 0 {
		new_pre_mult_result = equation[0]
	} else {
		new_pre_mult_result = preliminary_result * equation[0]
	}
	new_pre_plus_result := preliminary_result + equation[0]
	return resolveEquation(result, equation[1:], new_pre_plus_result) || resolveEquation(result, equation[1:], new_pre_mult_result)
}

func resolveEquationPart2(result int, equation []int, preliminary_result int) bool {
	if preliminary_result == result {
		return true
	}
	if preliminary_result > result {
		return false
	}
	if len(equation) == 0 {
		return false
	}
	new_pre_plus_result := preliminary_result + equation[0]
	var new_pre_mult_result int
	var new_pre_conc_result int
	if preliminary_result == 0 {
		new_pre_conc_result = equation[0]
		new_pre_mult_result = equation[0]
	} else {
		pot_10_digits := int(math.Pow10(getAmountOfDigits(equation[0], 1)))
		new_pre_conc_result = (preliminary_result * (pot_10_digits)) + equation[0]
		new_pre_mult_result = preliminary_result * equation[0]
	}
	return resolveEquationPart2(result, equation[1:], new_pre_plus_result) || resolveEquationPart2(result, equation[1:], new_pre_mult_result) || resolveEquationPart2(result, equation[1:], new_pre_conc_result)
}

func getAmountOfDigits(num int, digits int) int {
	if num >= 10 {
		return getAmountOfDigits(num/10, digits+1)
	} else {
		return digits
	}
}
