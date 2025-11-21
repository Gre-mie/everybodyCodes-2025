// https://everybody.codes/event/2025/quests/13

package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	var dial []int = getStructuredData("../notes/part1.txt")

	move_dial := 2025
	dial_index := move_dial % len(dial)

	if len(dial) <= dial_index {fmt.Printf("ERROR: index %v out of array bounds, arr len: %v\n", dial_index, len(dial)); return}
	fmt.Println("Results:", dial[dial_index])

}

func getStructuredData(path string) []int {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: couldn't read from file\n\toriginal err: %v\n", err); return []int{}}
	content := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	left := []int{1}
	right := []int{}
	for i, item := range content {
		num, err := strconv.Atoi(item)
		if err != nil {fmt.Printf("WARNING: couldn't convert string to int, %v not added\n", item)}
		
		if i % 2 == 0 { // even: left
			left = append(left, num)
		} else { // odd: right
			right = append(right, num)
		}
	}
	// add right elements to left array last in first out
	for i:=len(right)-1; i >= 0; i-- {
		left = append(left, right[i])
	}
	return left
}
