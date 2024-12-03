package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {
	parsedList1, parsedList2 := parseInput(input)
	// parsedList1, parsedList2 := parseInput(test)
	
	result_part_1 := part1(parsedList1, parsedList2)
	result_part_2 := part2(parsedList1, parsedList2)
	fmt.Print(result_part_1)
	fmt.Print(result_part_2)
}

func part1(parsedList1 []string, parsedList2 []string) int {
	sort.Strings(parsedList1)
	sort.Strings(parsedList2)
	result := 0

	for i := range parsedList1 {
		item1, _ := strconv.Atoi(parsedList1[i])
		item2, _ := strconv.Atoi(parsedList2[i])
		distance := item2 - item1
		if distance < 0 {
			distance = - distance
		}
		result += distance
	}
	return result
}

func part2(parsedList1 []string, parsedList2 []string) int {
	result := 0

	for i := range parsedList1 {
		item1, _ := strconv.Atoi(parsedList1[i])
		times := 0
		for j := range parsedList2 {
			item2, _ := strconv.Atoi(parsedList2[j])
			if item1 == item2 {
				times += 1
			}
		}
		score := item1 * times
		result += score
	}
	return result
}

func parseInput(input string) (parsedList1 []string, parsedList2 []string) {
	for _, line := range strings.Split(input, "\n") {
		items := strings.Split(line, "   ")
		parsedList1 = append(parsedList1, items[0])
		parsedList2 = append(parsedList2, items[1])
	}

	return parsedList1, parsedList2
}
