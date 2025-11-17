package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	Y int
	X int
}

func main() {
	var board [][]string = getBoard("../notes/part1.txt")
	if len(board) < 1 {fmt.Printf("ERROR: board is empty"); return}

	turns := 4
	var drag_st point = getDragonPosition(board)
	board_height := len(board)
	board_width := len(board[0])

	markBoard(turns, board_height, board_width, board, drag_st, "X")
	printboard(board)
	sheepEaten := countMarks(board, "X")


	fmt.Println("Results: sheep eaten =", sheepEaten)
}

// structure data from file
func getBoard(path string) [][]string {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: can't read from file\n\toriginal err: %v\n", err)}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")
	results := make([][]string, 0, len(lines))
	for _, line := range lines {
		row := make([]string, 0, len(line))
		for _, char := range []rune(line) {
			row = append(row, string(char))
		}
		results = append(results, row)
	}
	return results
}

// get dragons starting position
func getDragonPosition(board [][]string) point {
	if len(board) < 1 {
		fmt.Printf("WARNING: empty board passed to, getDragonPosition")
		return point{}
	}
	for y, row := range board {
		for x, char := range row {
			if char == "D" {
				return point{
					Y: y,
					X: x,
				}
			}
		}
	}
	fmt.Printf("WARNING: dragon not found by, getDragonPosition")
	return point{}
}

// get a list of points possible at current position
func getPossibleMoves(height, width int, current point) []point {
	moves := [][]int{
		// Y : X
		{-2,-1}, {-2, 1},
		{-1,2}, {-1,-2},
		{2,-1}, {2,1},
		{1,-2}, {1,2},
	}
	possiblePoints := make([]point, 0, 8)
	for _, move := range moves {
		newpoint := point{
			Y: current.Y+move[0],
			X: current.X+move[1],
		}
		if newpoint.Y < 0 || newpoint.X < 0 {continue}
		if newpoint.Y >= height || newpoint.X >= width {continue}
		possiblePoints = append(possiblePoints, newpoint)
	}
	return possiblePoints
}

func markBoard(turns, board_height, board_width int, board [][]string, current point, mark string) {
	if turns <= 0 {return}
	var possible []point = getPossibleMoves(board_height, board_width, current)
	// mark the kills with X
	for _, move := range possible {
		char := board[move.Y][move.X]
		if char == "S" {
			board[move.Y][move.X] = mark
		}
		//recurse
		markBoard(turns-1, board_height, board_width, board, move, mark)
	}
}

func countMarks(board [][]string, mark string) int {
	kills := 0
	for _, row := range board {
		for _, char := range row {
			if char == mark {
				kills++
			}
		}
	}

	return kills
}

func printboard(board [][]string) {
	for _, row := range board {
		fmt.Println(row)
	}
}

