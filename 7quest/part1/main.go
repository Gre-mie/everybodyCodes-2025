// https://everybody.codes/event/2025/quests/7

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// array, map
	names, rules := getDataSets("../notes/part1.txt")

	parsed := []string{}
	for _, name := range names {
		if parseName(name, rules) {
			parsed = append(parsed, name)
		}
	}

	fmt.Println(parsed)

}

func getDataSets(path string) (names []string, rules map[string]string) {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: cant read from file\n\toriginal err: %v\n", err)}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	names = strings.Split(lines[0], ",")
	rules = map[string]string{}
	for _, line := range lines[1:] {
		if line == "" {continue}
		pullchars := strings.Split(line, " > ")
		key := pullchars[0]
		rule := strings.ReplaceAll(pullchars[1], ",", "")
		_, ok := rules[key]
		if ok {fmt.Printf("WARNING: key %v in rules already exists, overwritten", key)}
		rules[key] = rule	
	}
	return names, rules
}

func parseName(name string, rules map[string]string) bool {
	if len(name) < 2 {
		fmt.Printf("WARNING: name '%v' is to short to compare to the rules\n", name)
		return true
	}
	n := []rune(name)
	for i:=0; i<len(n)-1; i++ {
		l := string(n[i])
		for key, rule := range rules {
			if l != key {continue}
			next := string(n[i+1])
			if !strings.Contains(rule, next) {return false}
		}
	}
	return true
}
