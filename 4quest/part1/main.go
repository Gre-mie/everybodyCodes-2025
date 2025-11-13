// https://everybody.codes/event/2025/quests/4

// used Bing AI for equasion

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../notes/part1.txt")
	if err != nil {fmt.Printf("ERROR: couldn't read from file\n\roriginal err: %v\n", err)}
	dat := strings.ReplaceAll(string(data), "\r\n", "\n")
	cogs := []float32{}
	for _, strnum := range strings.Split(dat, "\n") {
		n, err := strconv.Atoi(strnum)
		if err != nil {fmt.Printf("ERROR: couldn't convert string to int\n\toriginal err: %v\n", err)}
		cogs = append(cogs, float32(n))
	}
	high := cogs[0]
	low := cogs[len(cogs)-1]
	res := (high/low)*2025
	fmt.Println("Result:", int(res))	
}
