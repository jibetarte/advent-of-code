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

type Position struct {
	X int
	Y int
}

func main() {
	result_part_1 := part1()
	// result_part_2 := part2()
	fmt.Println(result_part_1)
	// fmt.Println(result_part_2)

}

func part1() int {
	var trailheads int
	parsed_input := parseInput(test)
	for i := range parsed_input {
		for j := range parsed_input[i] {
			if string(parsed_input[i][j]) == "0" {
				trailhead := &[]Position{}
				findTrailheads(i, j, i, j, parsed_input, 0, trailhead, true)
				trailheads += distinctTrailheads(*trailhead)
			}
		}
	}
	return trailheads
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
	return parsed_input
}

func findTrailheads(initial_x int, initial_y int, x int, y int, board [][]rune, previous_value int, trailheads_found *[]Position, first_time bool) {
	value, _ := strconv.Atoi(string(board[x][y]))
	if value == 9 && previous_value == 8 {
		*trailheads_found = append(*trailheads_found, Position{X: x, Y: y})
		return
	}
	if !first_time && (previous_value >= value || (value-previous_value) != 1 || (initial_x == x && initial_y == y)) {
		return
	}
	if x > 0 {
		findTrailheads(initial_x, initial_y, x-1, y, board, value, trailheads_found, false)
	}
	if y > 0 {
		findTrailheads(initial_x, initial_y, x, y-1, board, value, trailheads_found, false)
	}
	if x < len(board)-1 {
		findTrailheads(initial_x, initial_y, x+1, y, board, value, trailheads_found, false)
	}
	if y < len(board[0])-1 {
		findTrailheads(initial_x, initial_y, x, y+1, board, value, trailheads_found, false)
	}
}

func distinctTrailheads(trailheads []Position) (result int) {
	result = 0
	aux_trailheads := []Position{}
	for i := range trailheads {
		exists := false
		for j := range aux_trailheads {
			if trailheads[i].X == aux_trailheads[j].X && trailheads[i].Y == aux_trailheads[j].Y {
				exists = true
			}
		}
		if !exists {
			aux_trailheads = append(aux_trailheads, trailheads[i])
			result++
		}

	}
	return result
}
