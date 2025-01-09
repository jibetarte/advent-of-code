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

type Coordinates struct {
	X int
	Y int
}

type Robot struct {
	Position Coordinates
	Velocity Coordinates
}

func main() {
	result_part_1 := part1()
	// result_part_2 := part2()
	fmt.Println(result_part_1)
	// fmt.Println(result_part_2)

}

var x_tiles = 101
var y_tiles = 103

func part1() int {
	var result int
	robots := parseInput(input)
	var moved_robots []Robot
	for i := 1; i <= 100; i++ {
		moved_robots = make([]Robot, 0)
		for i := range robots {
			position_x := robots[i].Position.X + robots[i].Velocity.X
			position_y := robots[i].Position.Y + robots[i].Velocity.Y
			if position_x < 0 {
				position_x = x_tiles + position_x
			}
			if position_y < 0 {
				position_y = y_tiles + position_y
			}
			if position_x > x_tiles {
				position_x = position_x - x_tiles
			}
			if position_y > y_tiles {
				position_y = position_y - y_tiles
			}
			moved_robot := Robot{
				Position: Coordinates{X: position_x, Y: position_y},
				Velocity: Coordinates{X: robots[i].Velocity.X, Y: robots[i].Velocity.Y},
			}
			moved_robots = append(moved_robots, moved_robot)
		}
		robots = moved_robots
	}
	cuadrant_1, cuadrant_2, cuadrant_3, cuadrant_4 := calculateCuadrants(robots)
	result = cuadrant_1 * cuadrant_2 * cuadrant_3 * cuadrant_4
	return result
}

func part2() int {
	res := 0

	return res
}

func parseInput(input string) (parsed_input []Robot) {
	rows := strings.Split(input, "\n")
	for i := range rows {
		values := strings.Split(rows[i], " ")
		position := strings.Split(values[0], "=")[1]
		position_x, _ := strconv.Atoi(strings.Split(position, ",")[0])
		position_y, _ := strconv.Atoi(strings.Split(position, ",")[1])
		velocities := strings.Split(values[1], "=")[1]
		velocity_x, _ := strconv.Atoi(strings.Split(velocities, ",")[0])
		velocity_y, _ := strconv.Atoi(strings.Split(velocities, ",")[1])
		parsed_input = append(parsed_input, Robot{Position: Coordinates{X: position_x, Y: position_y}, Velocity: Coordinates{X: velocity_x, Y: velocity_y}})
	}
	return
}

func calculateCuadrants(robots []Robot) (cuadrant_1 int, cuadrant_2 int, cuadrant_3 int, cuadrant_4 int) {
	for i := range robots {
		position_x := robots[i].Position.X
		position_y := robots[i].Position.Y
		switch {
		case position_x < x_tiles/2 && position_y < y_tiles/2:
			cuadrant_1++
		case position_x < x_tiles/2 && position_y > y_tiles/2:
			cuadrant_2++
		case position_x > x_tiles/2 && position_y < y_tiles/2:
			cuadrant_3++
		case position_x > x_tiles/2 && position_y > y_tiles/2:
			cuadrant_4++
		}
	}
	return
}
