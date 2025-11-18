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
	var group Group = getGroup("../notes/part2.txt")
	if len(group) < 1 {fmt.Printf("WARNING: group array is empty\n")}

	rounds:=0
	swapsP1 := true
	swapsP2 := true
	for swapsP2 {
		if swapsP1 {
			// first stage
			swapsP1 = false
			for i:=0; i<len(group)-1; i++ {
				j:=i+1
				if group[i] > group[j] {
					group.move(i, j)
					swapsP1 = true
				}
			}
		}
		if !swapsP1 {
			// second stage
			swapsP2 = false
			for i:=0; i<len(group)-1; i++ {
				j:=i+1
				if group[i] < group[j] {
					group.move(j, i)
					swapsP2 = true
				}
			}
		}
		if swapsP2 {rounds++}
	}
	fmt.Println("Result:", rounds)

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
