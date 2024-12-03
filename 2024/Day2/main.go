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
	parsedList := parseInput(input)

	// result_part_1 := part1(parsedList)
	result_part_2 := part2(parsedList)
	// fmt.Print(result_part_1)
	fmt.Print(result_part_2)
}

func part1(parsedList [][]string) int {
	result := 0

	for i := range parsedList {
		is_safe := is_report_safe(parsedList[i])
		if is_safe {
			result += 1
		}
	}
	return result
}

func part2_old(parsedList [][]string) int {
	result := 0
	STARTING := 0
	INCREASING := 1
	DECREASING := 2
	report_state := STARTING

	for i := range parsedList {
		is_safe := true
		report_state = STARTING
		has_failed := false
		for j := range parsedList[i] {
			item, _ := strconv.Atoi(parsedList[i][j])
			if j < (len(parsedList[i]) - 1) {
				next_item, _ := strconv.Atoi(parsedList[i][j+1])
				difference := next_item - item
				switch report_state {
				case STARTING:
					if (difference >= -3) && (difference <= 3) && !(difference == 0) {
						if difference > 0 {
							report_state = INCREASING
						} else {
							report_state = DECREASING
						}
					} else {
						has_failed = true
					}
				case INCREASING:
					if difference <= 0 || difference > 3 {
						if has_failed {
							is_safe = false
							break
						} else {
							if j == 1 {
								next_item, _ = strconv.Atoi(parsedList[i][j+1])
								previous_item, _ := strconv.Atoi(parsedList[i][j-1])

								difference = next_item - previous_item
								if difference < 0 {
									report_state = DECREASING
								}
							}
							has_failed = true
							if j < (len(parsedList[i]) - 2) {
								next_item, _ = strconv.Atoi(parsedList[i][j+2])
								difference = next_item - item
								if difference <= 0 || difference > 3 {
									is_safe = false
									break
								}
							}
						}
					}
				case DECREASING:
					difference = -difference
					if difference <= 0 || difference > 3 {
						if has_failed {
							is_safe = false
							break
						} else {
							if j == 1 {
								next_item, _ = strconv.Atoi(parsedList[i][j+1])
								previous_item, _ := strconv.Atoi(parsedList[i][j-1])

								difference = next_item - previous_item
								if difference > 0 {
									report_state = INCREASING
								}
							}
							has_failed = true
							if j < (len(parsedList[i]) - 2) {
								next_item, _ = strconv.Atoi(parsedList[i][j+2])
								difference = item - next_item
								if difference <= 0 || difference > 3 {
									is_safe = false
									break
								}
							}
						}
					}
				}
			}
		}
		if is_safe {
			result += 1
		}
	}
	return result
}

func part2(parsedList [][]string) int {
	result := 0

	for i := range parsedList {
		is_safe := true
		for j := range parsedList[i] {
			destination := make([]string, len(parsedList[i]))
			copy(destination, parsedList[i])
			is_safe = is_report_safe(append(destination[:j], destination[j+1:]...))
			if is_safe {
				result += 1
				break
			}
		}
	}
	return result
}

func is_report_safe(parsedList []string) bool {
	STARTING := 0
	INCREASING := 1
	DECREASING := 2
	report_state := STARTING
	is_safe := true
	for j := range parsedList {
		item, _ := strconv.Atoi(parsedList[j])
		if j < (len(parsedList) - 1) {
			next_item, _ := strconv.Atoi(parsedList[j+1])
			difference := next_item - item
			switch report_state {
			case STARTING:
				if (difference >= -3) && (difference <= 3) && !(difference == 0) {
					if difference > 0 {
						report_state = INCREASING
					} else {
						report_state = DECREASING
					}
				} else {
					is_safe = false
					break
				}
			case INCREASING:
				if difference <= 0 || difference > 3 {
					is_safe = false
					break
				}
			case DECREASING:
				difference = -difference
				if difference <= 0 || difference > 3 {
					is_safe = false
					break
				}
			}
		}
	}
	return is_safe
}

func parseInput(input string) (parsedList [][]string) {
	for _, line := range strings.Split(input, "\n") {
		items := strings.Split(line, " ")
		parsedList = append(parsedList, items)
	}

	return parsedList
}
