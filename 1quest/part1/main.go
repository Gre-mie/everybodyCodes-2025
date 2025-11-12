// https://everybody.codes/event/2025/quests/1

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type task struct {
	direction string
	steps int
}

func (t task) debug() {
	fmt.Printf("direction: %v, steps: %v\n", t.direction, t.steps)
}

func main() {
	names, instructions := getContentSections("../notes/part1.txt")
	
	position := 0
	for _, in := range instructions {
		switch in.direction {
		case "L":
			position -= in.steps
		case "R":
			position += in.steps
		default:
			fmt.Printf("ERROR: not L or R direction\n")
		}
		
		if position < 0 {
			position = 0
		}
		if position >= len(names) {
			position = len(names)-1
		}
	}

	fmt.Println("Result:", names[position])

}

func getContentSections(filepath string) (names []string, instructions []task) {
	file, err := os.Open(filepath)
	if err != nil {fmt.Printf("ERROR: file could not be oppened\noriginal err: %v\n", err)}

	contents, err := io.ReadAll(file)
	if err != nil {fmt.Printf("ERROR: could not read contents of file\noriginal err: %v\n", err)}
	file.Close()
	
	data := strings.Split(string(contents), "\n")
	names = strings.Split(data[0], ",")
	dat := strings.Split(data[len(data)-1], ",")

	// create array of instructions seperating parts and converting types
	for _, d := range dat {
		step, err := strconv.Atoi(string(d[1:]))
		if err != nil {fmt.Printf("ERROR: couldn't convert string to int\n original err: %v\n", err)}
		instructions = append(instructions, task{
			direction: string(d[0]),
			steps: step,
		})
	}

	return names, instructions
}

