package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input string

//go:embed test.txt
var test string

func main() {
	result_part_1 := part1()
	// result_part_2 := part2()
	fmt.Println(result_part_1)
	// fmt.Println(result_part_2)

}

func part1() int {
	parsed_input := parseInput(input)
	disk := parseDiskSpace(parsed_input)
	new_disk := compactDisk(disk)
	// printDisk(new_disk)
	checksum := getChecksum(new_disk)

	return checksum
}

func part2() int {
	res := 0

	return res
}

func parseInput(input string) (parsed_input []int) {
	runes := []rune(input)
	for i := range runes {
		num, _ := strconv.Atoi(string(runes[i]))
		parsed_input = append(parsed_input, num)
	}
	return parsed_input
}

func parseDiskSpace(input []int) (disk_space []*int) {
	is_free_space := false
	disk_index := 0
	count := 0
	for i := range input {
		for j := 0; j < input[i]; j++ {
			if is_free_space {
				disk_space = append(disk_space, nil)
			} else {
				current_count := count
				disk_space = append(disk_space, &current_count)
			}
			disk_index++
		}
		if !is_free_space {
			count++
		}
		is_free_space = !is_free_space
	}
	return disk_space
}

func compactDisk(disk []*int) (new_disk []*int) {
	i := 0
	j := len(disk) - 1
	for i < len(disk) && i <= j {
		if disk[i] == nil {
			new_disk = append(new_disk, disk[j])
			for disk[j-1] == nil {
				j--
			}
			j--
		} else {
			new_disk = append(new_disk, disk[i])
		}
		i++
	}
	return new_disk
}

func printDisk(disk []*int) {
	for i := range disk {
		if disk[i] == nil {
			fmt.Print(disk[i])
		} else {
			fmt.Print(*disk[i])
		}
	}
	fmt.Println()
}

func getChecksum(disk []*int) (checksum int) {
	for i := range disk {
		checksum += i * *disk[i]
	}
	return checksum
}
