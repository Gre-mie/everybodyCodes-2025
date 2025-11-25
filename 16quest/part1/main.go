// https://everybody.codes/event/2025/quests/16

package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

var (
	wall_places int = 90 // number of columns in the wall
)

func main() {
	commands := getStructuredData("../notes/part1.txt")
	blocks := getBlocks(commands, wall_places)

	fmt.Println("Result:", sum(blocks))
}

func getStructuredData(path string) []int {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: couldn't read from file\n\toriginal err: %v\n", err)}
	nums := []int{}
	for _, n := range strings.Split(string(data), ",") {
		num, err := strconv.Atoi(n)
		if err != nil {fmt.Printf("WARNING: couldn't convert string to int, skiped %v\n", n); continue}
		nums = append(nums, num)
	}
	return nums
}

func getBlocks(commands []int, columns int) []int {
	results := make([]int, columns)
	for _, command := range commands {
		for i, _ := range results {
			col := i+1
			if col == command{results[i]++; continue}
			if col < command {continue}
			if col % command == 0 {results[i]++}
		}
	}
	return results
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
