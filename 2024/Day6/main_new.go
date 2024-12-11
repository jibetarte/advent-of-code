package main

// import (
// 	_ "embed"
// 	"fmt"
// 	"strings"
// )

// //go:embed input.txt
// var input string

// //go:embed test.txt
// var test string

// func main() {
// 	parsed_input := parseInput(test)

// 	// result_part_1 := part1(parsed_input)
// 	result_part_2 := part2(parsed_input)
// 	// fmt.Println(result_part_1)
// 	fmt.Println(result_part_2)
// }

// func part1(parsed_input [][]rune) int {
// 	result := 0
// 	dir := "N"
// 	pos_x, pos_y := findPlayer(parsed_input)
// 	parsed_input[pos_y][pos_x] = rune('X')
// 	finalBoard := advancePlayer(parsed_input, pos_x, pos_y, dir)
// 	printBoard(finalBoard)
// 	result = countX(finalBoard)
// 	return result
// }

// func part2(parsed_input [][]rune) int {
// 	result := 0
// 	dir := "N"
// 	pos_x, pos_y := findPlayer(parsed_input)
// 	parsed_input[pos_y][pos_x] = rune('X')
// 	finalBoard := advancePlayer(parsed_input, pos_x, pos_y, dir)

// 	printBoard(finalBoard)
// 	return result

// }

// func findPlayer(parsed_input [][]rune) (int, int) {
// 	for y := range parsed_input {
// 		for x := range parsed_input[y] {
// 			if string(parsed_input[y][x]) == "^" {
// 				return x, y
// 			}
// 		}
// 	}
// 	return 0, 0
// }

// func countX(parsed_input [][]rune) int {
// 	count := 0
// 	for y := range parsed_input {
// 		for x := range parsed_input[y] {
// 			if string(parsed_input[y][x]) == "X" {
// 				count += 1
// 			}
// 		}
// 	}
// 	return count
// }

// func printBoard(parsed_input [][]rune) {
// 	for y := range parsed_input {
// 		for x := range parsed_input[y] {
// 			fmt.Print(string(parsed_input[y][x]))
// 		}
// 		fmt.Println()
// 	}
// }

// func getNewDir(dir string) string {
// 	switch dir {
// 	case "N":
// 		return "E"
// 	case "E":
// 		return "S"
// 	case "S":
// 		return "W"
// 	case "W":
// 		return "N"
// 	}
// 	fmt.Println("Something went wrong, dir: ", dir)
// 	return ""
// }

// func parseInput(input string) (parsed_input [][]rune) {
// 	rows := strings.Split(input, "\n")
// 	for i := range rows {
// 		runes := []rune(rows[i])
// 		parsed_input = append(parsed_input, runes)
// 	}
// 	return parsed_input
// }

// func advancePlayerPart2(parsed_input [][]rune, pos_x int, pos_y int, dir string, count int) int {
// 	switch dir {
// 	case "N":
// 		if pos_y == 0 {
// 			return count
// 		}
// 		if string(parsed_input[pos_y-1][pos_x]) == "#" {
// 			copied_board := copy2DSlice(parsed_input)
// 			copied_board[pos_y-1][pos_x] = rune('#')
// 			is_loop := searchLoop(copied_board, pos_x, pos_y, dir, pos_x, pos_y, true)
// 			fmt.Println()

// 			if is_loop {
// 				count++
// 			}
// 			return advancePlayerPart2(parsed_input, pos_x, pos_y, getNewDir(dir), count)
// 		}

// 		return advancePlayerPart2(parsed_input, pos_x, pos_y-1, dir, count)
// 	case "S":
// 		if pos_y == len(parsed_input)-1 {
// 			return count
// 		}
// 		if string(parsed_input[pos_y+1][pos_x]) == "#" {
// 			copied_board := copy2DSlice(parsed_input)
// 			copied_board[pos_y+1][pos_x] = rune('#')
// 			is_loop := searchLoop(copied_board, pos_x, pos_y, dir, pos_x, pos_y, true)
// 			fmt.Println()

// 			if is_loop {
// 				count++
// 			}
// 			return advancePlayerPart2(parsed_input, pos_x, pos_y, getNewDir(dir), count)
// 		}

// 		return advancePlayerPart2(parsed_input, pos_x, pos_y+1, dir, count)
// 	case "E":
// 		if pos_x == len(parsed_input[pos_x])-1 {
// 			return count
// 		}
// 		if string(parsed_input[pos_y][pos_x+1]) == "#" {
// 			copied_board := copy2DSlice(parsed_input)
// 			copied_board[pos_y][pos_x+1] = rune('#')
// 			is_loop := searchLoop(copied_board, pos_x, pos_y, dir, pos_x, pos_y, true)
// 			fmt.Println()
// 			if is_loop {
// 				count++
// 			}
// 			return advancePlayerPart2(parsed_input, pos_x, pos_y, getNewDir(dir), count)
// 		}
// 		return advancePlayerPart2(parsed_input, pos_x+1, pos_y, dir, count)
// 	case "W":
// 		if pos_x == 0 {
// 			return count
// 		}
// 		if pos_x == len(parsed_input[pos_x]) {
// 			return count
// 		}
// 		if string(parsed_input[pos_y][pos_x-1]) == "#" {
// 			copied_board := copy2DSlice(parsed_input)
// 			copied_board[pos_y][pos_x-1] = rune('#')
// 			is_loop := searchLoop(copied_board, pos_x, pos_y, dir, pos_x, pos_y, true)
// 			fmt.Println()

// 			if is_loop {
// 				count++
// 			}
// 			return advancePlayerPart2(parsed_input, pos_x, pos_y, getNewDir(dir), count)
// 		}

// 		return advancePlayerPart2(parsed_input, pos_x-1, pos_y, dir, count)
// 	}
// 	return count
// }

// func copy2DSlice(original [][]rune) [][]rune {
// 	// Create a new slice with the same length as the original
// 	copied := make([][]rune, len(original))

// 	for i, row := range original {
// 		// Create a new slice for each row and copy its elements
// 		copied[i] = make([]rune, len(row))
// 		copy(copied[i], row)
// 	}

// 	return copied
// }

// func searchLoop(board [][]rune, pos_x int, pos_y int, dir string, initial_pos_x int, initial_pos_y int, is_first_time bool) bool {
// 	fmt.Println(pos_x, pos_y, dir)
// 	switch dir {
// 	case "N":
// 		if pos_y == 0 {
// 			return false
// 		}
// 		if initial_pos_x == pos_x && initial_pos_y == pos_y-1 && !is_first_time && string(board[pos_y-1][pos_x]) == "#" {
// 			return true
// 		}
// 		if is_first_time {
// 			is_first_time = false
// 		}
// 		if string(board[pos_y-1][pos_x]) == "#" {
// 			return searchLoop(board, pos_x, pos_y, getNewDir(dir), initial_pos_x, initial_pos_y, is_first_time)
// 		}
// 		return searchLoop(board, pos_x, pos_y-1, dir, initial_pos_x, initial_pos_y, is_first_time)
// 	case "S":
// 		if pos_y == len(board)-1 {
// 			return false
// 		}
// 		if initial_pos_x == pos_x && initial_pos_y == pos_y+1 && !is_first_time && string(board[pos_y+1][pos_x]) == "#" {
// 			return true
// 		}
// 		if is_first_time {
// 			is_first_time = false
// 		}
// 		if string(board[pos_y+1][pos_x]) == "#" {
// 			return searchLoop(board, pos_x, pos_y, getNewDir(dir), initial_pos_x, initial_pos_y, is_first_time)
// 		}
// 		return searchLoop(board, pos_x, pos_y+1, dir, initial_pos_x, initial_pos_y, is_first_time)
// 	case "E":
// 		if pos_x == len(board[pos_x])-1 {
// 			return false
// 		}
// 		if initial_pos_x == pos_x+1 && initial_pos_y == pos_y && !is_first_time && string(board[pos_y][pos_x+1]) == "#" {
// 			return true
// 		}
// 		if is_first_time {
// 			is_first_time = false
// 		}
// 		if string(board[pos_y][pos_x+1]) == "#" {
// 			return searchLoop(board, pos_x, pos_y, getNewDir(dir), initial_pos_x, initial_pos_y, is_first_time)
// 		}
// 		return searchLoop(board, pos_x+1, pos_y, dir, initial_pos_x, initial_pos_y, is_first_time)
// 	case "W":
// 		if pos_x == 0 {
// 			return false
// 		}
// 		if initial_pos_x == pos_x-1 && initial_pos_y == pos_y && !is_first_time && string(board[pos_y][pos_x-1]) == "#" {
// 			return true
// 		}
// 		if is_first_time {
// 			is_first_time = false
// 		}
// 		if string(board[pos_y][pos_x-1]) == "#" {
// 			return searchLoop(board, pos_x, pos_y, getNewDir(dir), initial_pos_x, initial_pos_y, is_first_time)
// 		}
// 		return searchLoop(board, pos_x-1, pos_y, dir, initial_pos_x, initial_pos_y, is_first_time)
// 	}
// 	fmt.Println("Something went wrong")
// 	return false
// }

// func advancePlayer(parsed_input [][]rune, pos_x int, pos_y int, dir string) [][]rune {
// 	switch dir {
// 	case "N":
// 		if pos_y == 0 {
// 			return parsed_input
// 		}
// 		if string(parsed_input[pos_y-1][pos_x]) == "#" {
// 			return advancePlayer(parsed_input, pos_x, pos_y, getNewDir(dir))
// 		}
// 		if string(parsed_input[pos_y-1][pos_x]) == "." {
// 			parsed_input[pos_y-1][pos_x] = rune('X')
// 		}
// 		return advancePlayer(parsed_input, pos_x, pos_y-1, dir)
// 	case "S":
// 		if pos_y == len(parsed_input)-1 {
// 			return parsed_input
// 		}
// 		if string(parsed_input[pos_y+1][pos_x]) == "#" {
// 			return advancePlayer(parsed_input, pos_x, pos_y, getNewDir(dir))
// 		}
// 		if string(parsed_input[pos_y+1][pos_x]) == "." {
// 			parsed_input[pos_y+1][pos_x] = rune('X')
// 		}
// 		return advancePlayer(parsed_input, pos_x, pos_y+1, dir)
// 	case "E":
// 		if pos_x == len(parsed_input[pos_x])-1 {
// 			return parsed_input
// 		}
// 		if string(parsed_input[pos_y][pos_x+1]) == "#" {
// 			return advancePlayer(parsed_input, pos_x, pos_y, getNewDir(dir))
// 		}
// 		if string(parsed_input[pos_y][pos_x+1]) == "." {
// 			parsed_input[pos_y][pos_x+1] = rune('X')
// 		}
// 		return advancePlayer(parsed_input, pos_x+1, pos_y, dir)
// 	case "W":
// 		if pos_x == 0 {
// 			return parsed_input
// 		}
// 		if string(parsed_input[pos_y][pos_x-1]) == "#" {
// 			return advancePlayer(parsed_input, pos_x, pos_y, getNewDir(dir))
// 		}
// 		if string(parsed_input[pos_y][pos_x-1]) == "." {
// 			parsed_input[pos_y][pos_x-1] = rune('X')
// 		}
// 		return advancePlayer(parsed_input, pos_x-1, pos_y, dir)
// 	}
// 	return parsed_input
// }
