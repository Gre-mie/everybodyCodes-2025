// https://everybody.codes/event/2025/quests/3

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("../notes/part1.txt")
	if err != nil {fmt.Printf("ERROR: couldn't read file\n\toriginal err: %v\n", err)}
	split_data := strings.Split(string(data), ",")
	
	content := make([]int, 0)
	for _, strnum := range split_data {
		n, err := strconv.Atoi(strnum)
		if err != nil {fmt.Printf("ERROR: couldn't convert string to number\n\toriginal err: %v\n", err)} else {
			content = append(content, n)
		}
	}
	slices.Sort(content)
	containers := slices.Compact(content)
	total := 0
	for _, size := range containers {
		total += size
	}

	fmt.Println("Result:", total)

}
