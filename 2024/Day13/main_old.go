package main

// import (
// 	_ "embed"
// 	"fmt"
// 	"regexp"
// 	"strconv"
// 	"strings"
// )

// type Game struct {
// 	A     [2]int
// 	B     [2]int
// 	Prize [2]int
// }

// //go:embed input.txt
// var input string

// //go:embed test.txt
// var test string

// func main() {
// 	result_part_1 := part1()
// 	// result_part_2 := part2()
// 	fmt.Println(result_part_1)
// 	// fmt.Println(result_part_2)

// }

// func part1() int {
// 	var result int
// 	parsed_input := parseInput(test)
// 	for i := range parsed_input {
// 		game := calculateGame(parsed_input[i])
// 		fmt.Println(game)
// 		result += game
// 	}
// 	return result
// }

// func part2() int {
// 	res := 0

// 	return res
// }

// func parseInput(input string) (parsed_input []Game) {
// 	rows := strings.Split(input, "\n")
// 	game := Game{}
// 	for i := range rows {
// 		type_row := i % 4
// 		switch type_row {
// 		case 0:
// 			x, _ := regexp.Compile(`.X\+(\d+).`)
// 			x_match := x.FindStringSubmatch(rows[i])[1]
// 			y, _ := regexp.Compile(`.Y\+(\d+)`)
// 			y_match := y.FindStringSubmatch(rows[i])[1]
// 			game.A[0], _ = strconv.Atoi(x_match)
// 			game.A[1], _ = strconv.Atoi(y_match)
// 		case 1:
// 			x, _ := regexp.Compile(`.X\+(\d+).`)
// 			x_match := x.FindStringSubmatch(rows[i])[1]
// 			y, _ := regexp.Compile(`.Y\+(\d+)`)
// 			y_match := y.FindStringSubmatch(rows[i])[1]
// 			game.B[0], _ = strconv.Atoi(x_match)
// 			game.B[1], _ = strconv.Atoi(y_match)
// 		case 2:
// 			x, _ := regexp.Compile(`.X\=(\d+).`)
// 			x_match := x.FindStringSubmatch(rows[i])[1]
// 			y, _ := regexp.Compile(`.Y\=(\d+)`)
// 			y_match := y.FindStringSubmatch(rows[i])[1]
// 			game.Prize[0], _ = strconv.Atoi(x_match)
// 			game.Prize[1], _ = strconv.Atoi(y_match)
// 		case 3:
// 			parsed_input = append(parsed_input, game)
// 			game = Game{}
// 		}
// 	}
// 	parsed_input = append(parsed_input, game)
// 	return
// }

// func maximum(num ...int) (res int) {
// 	res = 0
// 	for _, j := range num {
// 		if j > res {
// 			res = j
// 		}
// 	}
// 	return
// }

// func calculateGame(game Game) (result int) {
// 	var cache = make(map[int]int)
// 	a_x_raz := game.Prize[0] / game.A[0]
// 	b_x_raz := game.Prize[0] / game.B[0]
// 	a_y_raz := game.Prize[1] / game.A[1]
// 	b_y_raz := game.Prize[1] / game.B[1]
// 	max_raz := maximum(a_x_raz, b_x_raz, a_y_raz, b_y_raz)
// 	result = findPrize(game, 0, [2]int{0, 0}, max_raz, cache)
// 	return
// }

// func findPrize(game Game, tokens int, current_prize [2]int, n int, cache map[int]int) int {
// 	if current_prize[0] == game.Prize[0] && current_prize[1] == game.Prize[1] {
// 		return tokens
// 	}
// 	if current_prize[0] > game.Prize[0] || current_prize[1] > game.Prize[1] {
// 		return -1
// 	}
// 	if n <= 1 {
// 		return -1
// 	}
// 	if v, ok := cache[n]; ok {
// 		return v
// 	}
// 	// Press A button
// 	a_aux_curr_prize := [2]int{current_prize[0] + game.A[0], current_prize[1] + game.A[1]}
// 	a_tokens := findPrize(game, tokens+3, a_aux_curr_prize, n-1, cache)
// 	// Press B button
// 	b_aux_curr_prize := [2]int{current_prize[0] + game.B[0], current_prize[1] + game.B[1]}
// 	b_tokens := findPrize(game, tokens+1, b_aux_curr_prize, n-1, cache)
// 	if a_tokens != -1 {
// 		cache[n] = a_tokens
// 		return a_tokens
// 	} else if b_tokens != -1 {
// 		cache[n] = b_tokens
// 		return b_tokens
// 	}
// 	return -1
// }
