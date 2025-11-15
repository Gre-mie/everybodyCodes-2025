// https://everybody.codes/event/2025/quests/6

package main

import (
	"fmt"
	"os"
)

type typeset struct {
	chars string
	mentors int
	novices int
	indiv []string
}

func main() {
	var data string = getStructuredData("../notes/part1.txt")
	var A *typeset = getSetOf(data, "A", "a")

	pairs := 0
	mentors := 0
	for _, letter := range A.indiv {
		switch letter {
		case "A":
			mentors++
		case "a":
			if mentors == 0 {continue}
			pairs += mentors
		}
	}

	fmt.Println("Result:", pairs)
}

func getStructuredData(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: couldn't read data\n\toriginal err: %v", err)}
	return string(data)
}

func getSetOf(data, upper, lower string) *typeset {
	set := typeset {
		chars: upper+lower,
		mentors: 0,
		novices: 0,
		indiv: []string{},
	}

	for _, letter := range []rune(data) {
		l := string(letter)
		switch l {
		case upper:
			set.indiv = append(set.indiv, l)
			set.mentors++
		case lower:
			set.indiv = append(set.indiv, l)
			set.novices++
		}
	}
	return &set
}
