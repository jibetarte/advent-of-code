package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

type Plant struct {
	letter rune
	cost   int
}

func main() {
	result_part_1 := part1()
	// result_part_2 := part2()
	fmt.Println(result_part_1)
	// fmt.Println(result_part_2)

}

func part1() int {
	var result int
	parsed_input := parseInput(input)
	farm := initializeFarm()
	visited_plants := initializeVisitedPlants(len(parsed_input), len(parsed_input[0]))
	for i := range len(parsed_input) {
		for j := range parsed_input[i] {
			if !visited_plants[i][j] {
				area, perimeter := calculateArea(i, j, parsed_input[i][j], visited_plants, parsed_input)
				farm[parsed_input[i][j]-65].cost += area * perimeter
			}
		}
	}
	result = countCost(farm)
	return result
}

func part2() int {
	res := 0

	return res
}

func parseInput(input string) (parsed_input [][]rune) {
	rows := strings.Split(input, "\n")
	for i := range rows {
		runes := []rune(rows[i])
		parsed_input = append(parsed_input, runes)
	}
	return
}

func initializeFarm() (farm []Plant) {
	for i := 'A'; i <= 'Z'; i++ {
		farm = append(farm, Plant{letter: i, cost: 0})
	}
	return
}

func initializeVisitedPlants(x int, y int) (visitedFarm [][]bool) {
	for range x {
		var row []bool
		for range y {
			row = append(row, false)
		}
		visitedFarm = append(visitedFarm, row)
	}
	return
}

func calculateArea(pos_x int, pos_y int, letter rune, visited_plants [][]bool, parsed_input [][]rune) (area int, perimeter int) {
	s_area, n_area, e_area, w_area := 0, 0, 0, 0
	s_perimeter, n_perimeter, e_perimeter, w_perimeter := 0, 0, 0, 0
	if parsed_input[pos_x][pos_y] == letter {
		area = 1
		visited_plants[pos_x][pos_y] = true
	}
	if pos_x > 0 && parsed_input[pos_x-1][pos_y] == letter {
		if !visited_plants[pos_x-1][pos_y] {
			n_area, n_perimeter = calculateArea(pos_x-1, pos_y, letter, visited_plants, parsed_input)
		}
	} else {
		n_perimeter += 1
	}
	if pos_y > 0 && parsed_input[pos_x][pos_y-1] == letter {
		if !visited_plants[pos_x][pos_y-1] {
			e_area, e_perimeter = calculateArea(pos_x, pos_y-1, letter, visited_plants, parsed_input)
		}
	} else {
		e_perimeter += 1
	}
	if len(parsed_input)-1 > pos_x && parsed_input[pos_x+1][pos_y] == letter {
		if !visited_plants[pos_x+1][pos_y] {

			s_area, s_perimeter = calculateArea(pos_x+1, pos_y, letter, visited_plants, parsed_input)

		}
	} else {
		s_perimeter += 1
	}
	if len(parsed_input[0])-1 > pos_y && parsed_input[pos_x][pos_y+1] == letter {
		if !visited_plants[pos_x][pos_y+1] {

			w_area, w_perimeter = calculateArea(pos_x, pos_y+1, letter, visited_plants, parsed_input)

		}
	} else {
		w_perimeter += 1
	}

	area += s_area + n_area + e_area + w_area
	perimeter = s_perimeter + n_perimeter + e_perimeter + w_perimeter
	return
}

func countCost(farm []Plant) (cost int) {
	for i := range farm {
		cost += farm[i].cost
	}
	return
}
