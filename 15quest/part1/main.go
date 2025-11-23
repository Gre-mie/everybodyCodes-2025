// https://everybody.codes/event/2025/quests/15

package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

var (
	wall int = 1
	start int = 2 	// first postion should count as an empty space but never get changed
	end int = 3 	// last postion
)

type point struct {
	Y int
	X int
}

type steps struct {
	turn string
	step int
}

func main() {
	var data [][]steps = getStructuredData("../notes/testpart1.txt")
	commands := data[0]
	//commands = data[1]

	var grid *Grid = newGrid()
	grid.mapCave(commands) 

	grid.printGrid() //

	fmt.Println("Results:", grid.getShortestPath())

}

func getStructuredData(path string) [][]steps {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: couldn't read file\n\toriginal err: %v\n", err); return [][]steps{}}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	results := make([][]steps, 0, len(lines))
	for i, line := range lines {
		row := make([]steps, 0, len(lines[i]))
		commands := strings.Split(line, ",")
		for _, command := range commands {
			num, err := strconv.Atoi(command[1:])
			if err != nil {fmt.Printf("WARNING: couldnt' convert string to int, %v not added\n", command[1:])}
			row = append(row, steps{
				turn: string(command[0]),
				step: num,
			})
		}
		results = append(results, row)
	}
	return results
}
