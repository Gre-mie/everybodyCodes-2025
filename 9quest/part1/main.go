package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var DNA map[int]string = getSequences("../notes/part1.txt")
	childKey := getChildKey(DNA)
	matches := []int{}
	for key, _ := range DNA {
		if key == childKey {continue}
		match := getMatches(DNA[childKey], DNA[key])
		matches = append(matches, match)
	}
	fmt.Println("Results:", matches[0]*matches[1])
}

func getSequences(path string) map[int]string {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: couldn't read from file\n\toriginal err: %v\n", err)}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")
	results := make(map[int]string, len(lines))
	for _, line := range lines {
		splitkey := strings.Split(line, ":")
		k, err := strconv.Atoi(splitkey[0])
		if err != nil {fmt.Printf("ERROR: couldn't convert string to int\n\toriginal err: %v\n", err)}
		key := k
		_, ok := results[key]
		if ok {fmt.Printf("WARNING: overwriting key %v", key)}
		results[key] = splitkey[1]
	}
	return results
}

func getChildKey(data map[int]string) int {
	if len(data) < 1 {
		fmt.Printf("WARNING: empty map given to, getChildKey\n")
		return -1
	}
	child := 1
	seg := 0
	for seg<len(data[child]) {
		if data[child][seg] == data[2][seg] || data[1][seg] == data[3][seg] {
			seg++
		} else {
			child = 2
			break
		}
	}
	seg = 0
	for seg<len(data[child]) && child==2 {
		if data[child][seg] == data[1][seg] || data[child][seg] == data[3][seg] {
			seg++
		} else {
			child = 3
			break
		}
	}
	seg = 0
	for seg<len(data[child]) && child==3 {
		if data[child][seg] == data[1][seg] || data[child][seg] == data[2][seg] {
			seg++
		} else {break}
	}
	return child
}

func getMatches(dna1, dna2 string) int {
	match := 0
	for i:=0; i<len(dna1); i++ {
		if dna1[i] == dna2[i] {match++}
	}
	return match
}
