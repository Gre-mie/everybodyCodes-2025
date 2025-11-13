// https://everybody.codes/event/2025/quests/2

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sample struct {
	content []int
}

func (s sample) debug() {
	fmt.Printf("content: %v\n", s.content)
}

func (s sample) add(other sample) sample {
	result := make([]int, 2)
	for i:=0; i<2; i++ {
		result[i] = s.content[i] + other.content[i]
	}
	return sample{content: result,}
}

func (s sample) mul(other sample) sample {
	x1 := s.content[0]
	y1 := s.content[1]
	x2 := other.content[0]
	y2 := other.content[1]
	return sample{
		content: []int{(x1*x2)-(y1*y2), (x1*y2)+(y1*x2)},
	}
}

func (s sample) div(other sample) sample {
	x1:=s.content[0]
	y1:=s.content[1]
	x2:=other.content[0]
	y2:=other.content[1]
	var res1 int
	var res2 int
	if x1 == 0 || x2 == 0 {res1 = 0} else {res1 =x1/x2}
	if y1 == 0 || y2 == 0 {res2 = 0} else {res2 =y1/y2}
	return sample{
		content: []int{res1, res2},
	}
}

func (s sample) calculate(data sample) sample {
	m := s.mul(s)
	d := m.div(sample{
		content: []int{10,10},
	})
	a := d.add(data)
	return a
}

func (s sample) format() string {
	return fmt.Sprintf("[%v,%v]", s.content[0], s.content[1])
}


func main() {
	content := getData("../notes/part1.txt")

	start := sample{
		content: []int{0,0},
	}

	iterations := 3
	for i:=0; i<iterations; i++ {
		start = start.calculate(content)
	}
	println("result:", start.format())
}

func getData(path string) sample {
	content, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: couldn't read from file\noriginal err: %v\n", err)}
	normalised := strings.ReplaceAll(string(content), "\r\n", "\n")

	// structure data
	lines := strings.Split(normalised, "\n")

	pullNames := strings.SplitN(lines[0], "=", 2)
	pullNums := strings.Split(strings.Trim(pullNames[1], "[]"), ",")

	nums := make([]int, 0, 2)
	for _, num := range pullNums {
		n, err := strconv.Atoi(num)
		if err != nil {fmt.Printf("ERROR: couldn't convert string to number\noriginal err: %v", err)}
		nums = append(nums, n)
	}
	sample := sample{
		content: nums,
	}
	
	return sample
}
