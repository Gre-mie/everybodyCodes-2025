// https://everybody.codes/event/2025/quests/8

// had help from math human to calculate oposite pins

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const TOTAL_PINS int = 32

	sequence := getData("../notes/part1.txt")
	if len(sequence) < 2 {
		fmt.Println("ERROR: sequence to small")
		return
	}

	middleCrossed := 0
	half := (TOTAL_PINS / 2)
	for i:=0; i<len(sequence)-1; i++ {
		threadPin := sequence[i]
		nextPin := sequence[i+1]
		if threadPin < nextPin {
			if threadPin+half == nextPin {
				middleCrossed++
			}
		}
		if threadPin > nextPin {
			if threadPin-half == nextPin {
				middleCrossed++
			}
		}
	}
	fmt.Println("Results:", middleCrossed)
}

func getData(path string) []int {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: couldn't read from file\n\toriginal err: %v\n",)}

	sequence := []int{}
	for _, n := range strings.Split(string(data), ",") {
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("ERROR: couldn't convert string to int\n\toriginal err: %v\n", err)
		} else {sequence = append(sequence, num)}
	}
	return sequence
}
