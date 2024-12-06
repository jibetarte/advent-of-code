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
	result := 0
	parsed_rules, parsed_updates := parseInput(test)
	for i := range parsed_updates {
		is_valid_update := true
		for j := range parsed_updates[i] {
			if j == 0 {
				continue
			}
			selected_item := parsed_updates[i][j]
			for k := 0; k < j; k++ {
				is_valid_item := search_invalid_rule(parsed_rules, selected_item, parsed_updates[i][k])
				if !is_valid_item {
					is_valid_update = false
					break
				}
			}
		}
		if is_valid_update {
			index := int(math.Floor(float64(len(parsed_updates[i]) / 2)))
			mid, _ := strconv.Atoi(parsed_updates[i][index])
			result += mid
		}
	}
	return result
}

func part2() int {
	result := 0
	parsed_rules, parsed_updates := parseInput(test)
	for i := range parsed_updates {
		is_valid_update := true
		for j := range parsed_updates[i] {
			if j == 0 {
				continue
			}
			selected_item := parsed_updates[i][j]
			for k := 0; k < j; k++ {
				is_valid_item := search_invalid_rule(parsed_rules, selected_item, parsed_updates[i][k])
				if !is_valid_item {
					is_valid_update = false
					break
				}
			}
		}
		if !is_valid_update {
			ordered_updates := order_update(parsed_rules, parsed_updates[i])

			index := int(math.Floor(float64(len(ordered_updates) / 2)))
			mid, _ := strconv.Atoi(ordered_updates[index])
			result += mid
		}
	}
	return result

}

func search_invalid_rule(parsed_rules [][]string, selected_item string, item_to_compare string) bool {
	is_valid_update := true
	for i := range parsed_rules {
		if parsed_rules[i][0] == selected_item && parsed_rules[i][1] == item_to_compare {
			return false
		}
	}
	return is_valid_update
}

func order_update(parsed_rules [][]string, parsed_update []string) []string {
	ordered_update := make([]string, len(parsed_update))
	copy(ordered_update, parsed_update)
	for i := 0; i < len(parsed_update)-1; i++ {
		for j := i + 1; j < len(parsed_update); j++ {
			if search_invalid_rule(parsed_rules, parsed_update[i], parsed_update[j]) {
				aux := parsed_update[i]
				parsed_update[i] = parsed_update[j]
				parsed_update[j] = aux
			}
		}
	}
	is_valid_update := true
	for i := range parsed_update {
		if i == 0 {
			continue
		}
		selected_item := parsed_update[i]
		for k := 0; k < i; k++ {
			is_valid_item := search_invalid_rule(parsed_rules, selected_item, parsed_update[k])
			if !is_valid_item {
				is_valid_update = false
				break
			}
		}
	}
	if !is_valid_update {
		return order_update(parsed_rules, parsed_update)
	}
	return parsed_update
}

func parseInput(input string) (input_rules [][]string, input_updates [][]string) {
	rows := strings.Split(input, "\n")
	is_rules := true
	for i := range rows {
		if rows[i] == "" {
			is_rules = false
			continue
		}
		if is_rules {
			input_rules = append(input_rules, strings.Split(rows[i], "|"))
		} else {
			input_updates = append(input_updates, strings.Split(rows[i], ","))
		}
	}
	return input_rules, input_updates
}
