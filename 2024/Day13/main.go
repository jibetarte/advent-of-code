package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	A     [2]int
	B     [2]int
	Prize [2]int
}

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
	var result int
	parsed_input := parseInput(test, 1)
	for i := range parsed_input {
		tokens := calculateGamePart1(parsed_input[i])
		if tokens > 0 {
			result += tokens
		}
	}
	return result
}

func part2() int {
	var result int
	parsed_input := parseInput(test, 10000000000)
	for i := range parsed_input {
		tokens := calculateGamePart2(parsed_input[i], 10000000000)
		if tokens > 0 {
			result += tokens
		}
	}
	return result
}

func parseInput(input string, unit int) (parsed_input []Game) {
	rows := strings.Split(input, "\n")
	game := Game{}
	for i := range rows {
		type_row := i % 4
		switch type_row {
		case 0:
			x, _ := regexp.Compile(`.X\+(\d+).`)
			x_match := x.FindStringSubmatch(rows[i])[1]
			y, _ := regexp.Compile(`.Y\+(\d+)`)
			y_match := y.FindStringSubmatch(rows[i])[1]
			game.A[0], _ = strconv.Atoi(x_match)
			game.A[1], _ = strconv.Atoi(y_match)
		case 1:
			x, _ := regexp.Compile(`.X\+(\d+).`)
			x_match := x.FindStringSubmatch(rows[i])[1]
			y, _ := regexp.Compile(`.Y\+(\d+)`)
			y_match := y.FindStringSubmatch(rows[i])[1]
			game.B[0], _ = strconv.Atoi(x_match)
			game.B[1], _ = strconv.Atoi(y_match)
		case 2:
			x, _ := regexp.Compile(`.X\=(\d+).`)
			x_match := x.FindStringSubmatch(rows[i])[1]
			y, _ := regexp.Compile(`.Y\=(\d+)`)
			y_match := y.FindStringSubmatch(rows[i])[1]
			prize_x, _ := strconv.Atoi(x_match)
			prize_y, _ := strconv.Atoi(y_match)
			game.Prize[0] = prize_x * unit
			game.Prize[1] = prize_y * unit
		case 3:
			parsed_input = append(parsed_input, game)
			game = Game{}
		}
	}
	parsed_input = append(parsed_input, game)
	return
}

func calculateGamePart1(game Game) (result int) {
	result = 0
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			x := i*game.A[0] + j*game.B[0]
			y := i*game.A[1] + j*game.B[1]
			if x == game.Prize[0] && y == game.Prize[1] {
				aux_result := (i * 3) + j
				if result == 0 || aux_result < result {
					result = aux_result
				}
			}
		}
	}
	return result
}

func calculateGamePart2Old(game Game, unit int) (result int) {
	result = 0
	for i := unit; i <= unit*10; i++ {
		for j := unit; j <= unit*10; j++ {
			x := i*game.A[0] + j*game.B[0]
			y := i*game.A[1] + j*game.B[1]
			if x == game.Prize[0] && y == game.Prize[1] {
				aux_result := (i * 3) + j
				if result == 0 || aux_result < result {
					result = aux_result
				}
			}
		}
	}
	return result
}

func calculateGamePart2(game Game, unit int) (result int) {
	result = 0
	go calculateGamePart2WThreads(game, unit, 1)
	go calculateGamePart2WThreads(game, unit, 2)
	go calculateGamePart2WThreads(game, unit, 3)
	go calculateGamePart2WThreads(game, unit, 4)
	go calculateGamePart2WThreads(game, unit, 5)
	go calculateGamePart2WThreads(game, unit, 6)
	go calculateGamePart2WThreads(game, unit, 7)
	go calculateGamePart2WThreads(game, unit, 8)
	go calculateGamePart2WThreads(game, unit, 9)
	go calculateGamePart2WThreads(game, unit, 10)
	return result
}

func calculateGamePart2WThreads(game Game, unit int, partition int) {
	result := 0
	fmt.Println(partition)
	for i := unit + (unit * partition); i <= unit+(unit*(partition+1)); i++ {
		for j := unit; j <= unit*10; j++ {
			x := i*game.A[0] + j*game.B[0]
			y := i*game.A[1] + j*game.B[1]
			if x == game.Prize[0] && y == game.Prize[1] {
				aux_result := (i * 3) + j
				if result == 0 || aux_result < result {
					result = aux_result
				}
			}
		}
	}
	fmt.Println(result)
}
