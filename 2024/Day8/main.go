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
	var antinodes []Position
	parsed_input := parseInput(input)
	var board_state [74][]Position
	for i := range parsed_input {
		for j := range parsed_input[i] {
			if string(parsed_input[i][j]) != "." {
				index := byte(parsed_input[i][j]) - 48
				board_state[index] = append(board_state[index], Position{i, j})
			}
		}
	}
	for k := range board_state {
		if len(board_state[k]) > 1 {
			for n := range board_state[k] {
				for m := range board_state[k][:n] {
					letter_antinodes := obtain_antinodes(board_state[k][n], board_state[k][m], len(parsed_input), len(parsed_input[0]))
					for l := range letter_antinodes {
						if !exists_antinode(letter_antinodes[l], antinodes) {
							antinodes = append(antinodes, letter_antinodes[l])
						}
					}
				}
			}
		}
	}
	return len(antinodes)
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

func obtain_antinodes(initial_node Position, final_node Position, x_max int, y_max int) []Position {
	offset_x := final_node.X - initial_node.X
	offset_y := final_node.Y - initial_node.Y
	initial_antinode := Position{initial_node.X - offset_x, initial_node.Y - offset_y}
	final_antinode := Position{final_node.X + offset_x, final_node.Y + offset_y}
	var antinodes []Position
	if initial_antinode.X >= 0 && initial_antinode.X < x_max && initial_antinode.Y >= 0 && initial_antinode.Y < y_max {
		antinodes = append(antinodes, initial_antinode)
	}
	if final_antinode.X >= 0 && final_antinode.X < x_max && final_antinode.Y >= 0 && final_antinode.Y < y_max {
		antinodes = append(antinodes, final_antinode)
	}
	return antinodes
}

func exists_antinode(antinode Position, antinodes []Position) bool {
	exists := false
	for i := range antinodes {
		if antinode.X == antinodes[i].X && antinode.Y == antinodes[i].Y {
			return true
		}
	}
	return exists
}
