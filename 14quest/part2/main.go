// https://everybody.codes/event/2025/quests/14

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var current_state [][]string = structureData("../notes/part2.txt")

	rounds := 2025
	var counts []int = make([]int, 0, rounds)
	for i:=0; i<rounds; i++ {
		current_state = nextState(current_state) // alters current state
		counts = append(counts, countTotalAlive(current_state))
	}
	fmt.Println("Results:", sumCounts(counts))
}

func structureData(path string) [][]string {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: couldn't read file\n\toriginal err: %v\n", err)}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	grid := make([][]string, 0, len(lines))
	for _, line := range lines {
		row := make([]string, 0)
		for _, char := range []rune(line) {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}
	return grid
}

func nextState(state [][]string) [][]string {
	if len(state) < 1 {fmt.Printf("WARNING: state is empty\n"); return [][]string{}}
	var new [][]string = copygrid(state)
	for i, row := range state {
		for j, _ := range row {
			surrounded_living := countAlive(i, j, state)
			if surrounded_living % 2 == 0 {
				if state[i][j] == "#" {new[i][j] = "."}
				if state[i][j] == "." {new[i][j] = "#"}
			}
		}
	}
	return new
}

func countAlive(i, j int, grid [][]string) int {
	alive := 0
	if inGrid(i-1, j-1, grid) {if grid[i-1][j-1] == "#" {alive++}}
	if inGrid(i-1, j+1, grid) {if grid[i-1][j+1] == "#" {alive++}}
	if inGrid(i+1, j-1, grid) {if grid[i+1][j-1] == "#" {alive++}}
	if inGrid(i+1, j+1, grid) {if grid[i+1][j+1] == "#" {alive++}}
	return alive
}

func inGrid(i, j int, grid [][]string) bool {
	if len(grid) < 1 {return false}
	if i < 0 || j < 0 || i>= len(grid) || j >= len(grid[i]) {return false}
	return true
}

func copygrid(state [][]string) [][]string {
	if len(state) < 1 {return [][]string{}}
	cp := make([][]string, 0, len(state))
	for _, row := range state {
		r := make([]string, 0, len(row))
		for _, char := range row {
			r = append(r, char)
		}
		cp = append(cp, r)
	}
	return cp
}

func printState(state [][]string) {
	for _, row := range state {
		fmt.Println(row)
	}
}

func countTotalAlive(current_state [][]string) int {
	if len(current_state) < 1 {return 0}
	total := 0
	for i, row := range current_state {
		for j, _ := range row {
			if current_state[i][j] == "#" {total += 1}
		}
	}
	return total
}

func sumCounts(counts []int) int {
	if len(counts) == 0 {return 0}
	return counts[0] + sumCounts(counts[1:])
}
