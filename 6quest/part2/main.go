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

func (t *typeset) getNumOfPairs() int {
	pairs := 0
	mentors := 0
	for _, letter := range t.indiv {
		switch letter {
		case string(t.chars[0]):
			mentors++
		case string(t.chars[1]):
			if mentors == 0 {continue}
			pairs += mentors
		}
	}
	return pairs
}

func main() {
	var data string = getStructuredData("../notes/part2.txt")
	var A *typeset = getSetOf(data, "A", "a")
	var B *typeset = getSetOf(data, "B", "b")
	var C *typeset = getSetOf(data, "C", "c")

	a_pairs := A.getNumOfPairs()
	b_pairs := B.getNumOfPairs()
	c_pairs := C.getNumOfPairs()

	fmt.Println("Result:", a_pairs+b_pairs+c_pairs)
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
