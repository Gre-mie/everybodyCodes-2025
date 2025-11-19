package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type point struct {
	Y int
	X int
}

func main() {
	var barrels [][]int = getStructuredData("../notes/part2.txt")
	
	var lit [][]int = make([][]int, len(barrels))
	for i, row := range barrels {
		lit[i] = make([]int, len(row))
	}
	lightBarrels(barrels, lit, 0, 0)
	lightBarrels(barrels, lit, len(barrels)-1, len(barrels[0])-1)
	fmt.Println("Results:", countLitBarrels(lit))
}

func getStructuredData(path string) [][]int {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: couldn't read from file\n\toriginal err: %v\n", err); return [][]int{}}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	barrels := make([][]int, 0, len(lines))
	for _, line := range lines {
		row := make([]int, 0, len(line))
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {fmt.Printf("WARNING: couldn't convert string to int, %v not added\n", num)}
			row = append(row, num)
		}
		barrels = append(barrels, row)
	}
	return barrels
}

func printBarrels(barrels [][]int) {
	if len(barrels) < 1 {fmt.Printf("WARNING: barrels array is empty\n"); return}
	for _, row := range barrels {
		fmt.Printf("[")
		for _, char := range row {
			switch len(string(char)) {
			case 0:
				fmt.Printf("   ")
			case 1:
				fmt.Printf(" %v ", char)
			default:
				fmt.Printf("%v ", char)
			}
		}
		fmt.Printf("]\n")
	}
}

func lightBarrels(barrels, lit [][]int, startY, startX int) {
	if len(barrels) < 1 {fmt.Printf("WARNING: barrels array is empty\n"); return}

	lit[startY][startX] = -1 // fire started here
	light(point{Y:startY, X:startX}, barrels, lit, []point{
		point{Y:-1, X:0}, // up
		point{Y:1, X:0},  // down
		point{Y:0, X:-1}, // left
		point{Y:0, X:1},  // right
	})
}

func light(current point, barrels, lit [][]int, moves []point) {
	if current.Y < 0 || current.Y >= len(barrels) || current.X < 0 {return} // out of grid
	if current.X >= len(barrels[current.Y]) {return} // out of grid

	new_fire := []point{}
	// mark fires in lit and add their position to new_fire
	for _, move := range moves {
		new := point{Y: current.Y+move.Y, X: current.X+move.X}
		if new.Y < 0 || new.Y >= len(barrels) || new.X < 0 || new.X >= len(barrels[0]) {continue}
		if lit[new.Y][new.X] < 0 {continue} // ignore visited
		if barrels[current.Y][current.X] >= barrels[new.Y][new.X] {
			lit[new.Y][new.X] = -1
			new_fire = append(new_fire, new)
		} else {
			lit[new.Y][new.X]++
		}
	}
	// recurse for newly started fires
	for _, position := range new_fire {
		light(position, barrels, lit, moves)
	}
}

func countLitBarrels(barrels [][]int) int {
	if len(barrels) < 1 {fmt.Printf("WARNING: barrels array is empty\n"); return 0}
	lit := 0
	for _, row := range barrels {
		for _, barrel := range row {
			if barrel < 0 {lit++}
		}
	}
	return lit
}
