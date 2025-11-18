package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Group []int

func (g Group) move(from_index, to_index int) {
	g[from_index]--
	g[to_index]++
}

func (g Group) checksum() int {
	checksum := 0
	for i:=0; i<len(g); i++ {
			num := g[i]
			checksum += (i+1)*num
		}
	return checksum
}

func main() {
	var group Group = getGroup("../notes/part1.txt")
	if len(group) < 1 {fmt.Printf("WARNING: group array is empty\n")}

	total_rounds := 10
	rounds:=0
	swaps := true
	for rounds < total_rounds {
		if swaps {
			// first stage
			swaps = false
			for i:=0; i<len(group)-1; i++ {
				j:=i+1
				if group[i] > group[j] {
					group.move(i, j)
					swaps = true
				}
			}
		}
		if !swaps {
			// second stage
			for i:=0; i<len(group)-1; i++ {
				j:=i+1
				if group[i] < group[j] {
					group.move(j, i)
				}
			}
		}
		rounds++
	}
	fmt.Println("Result:", group.checksum())

}

func getGroup(path string) Group {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: couldn't read from file\n\toriginal err: %v\n", err)}
	content := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	var group Group = make(Group, 0, len(content))
	for _, item := range content {
		num, err := strconv.Atoi(item)
		if err == nil {group = append(group, num)} else {
			fmt.Printf("WARNING: could not convert string to int. %v not added to array\n", item)
		}
	}
	return group
}
