// https://everybody.codes/event/2025/quests/5

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bone struct {
	last *bone // if nil is the root seg
	next *bone // if nil is the last seg
	// pointers used to check for nil
	left *int
	spine *int // compare by spine
	right *int
}

func createFishboneStructure(arr []int) (root *bone) {
	if len(arr) < 1 {return &bone{}}
	root = &bone{spine: &arr[0],}

	for i, n := range arr {
		if i == 0 {continue} // discount the first of the array
		current := root
		stop := false
		for !stop {
			//stop if a number is placed
			switch {
			case n < *current.spine:
				if current.left == nil {
					current.left = &arr[i]
					stop = true
				}
			case n > *current.spine:
				if current.right == nil {
					current.right = &arr[i]
					stop = true
				}
			}
			if stop {break}
			if current.next == nil {
				newbone := &bone{spine: &arr[i],}
				current.next = newbone
				current = newbone
				stop = true
			} else {
				current = current.next
			}
		}
	}
	//root.printBone()
	return root
}

func (b *bone) printBone() {
	fmt.Println("root")
	current := b
	stop := false
	for !stop {
		if current.next == nil {stop=true}
		arr := []string{}
		if current.left == nil {arr = append(arr, " ")} else {arr = append(arr, fmt.Sprintf("%v",*current.left))}
		if current.spine == nil {arr = append(arr, " ")} else {arr = append(arr, fmt.Sprintf("%v",*current.spine))}
		if current.right == nil {arr = append(arr, " ")} else {arr = append(arr, fmt.Sprintf("%v",*current.right))}

		fmt.Println(arr)
		current = current.next
	}
}

func (b *bone) sumSpine() int {
	sum := ""
	current := b
	for {
		if current == nil {break}
		sum += fmt.Sprintf("%v", *current.spine)
		current = current.next
	}
	if sum == "" {return 0}
	total, err := strconv.Atoi(sum)
	if err != nil {fmt.Printf("ERROR: couldn't convert string to int\n\toriginal err: %v\n", err)}
	return total
}

func main() {
	structures := getStructuredData("../notes/part2.txt")

	qualities := []int{}
	for _, arr := range structures {
		fishbone := createFishboneStructure(arr)
		sum := fishbone.sumSpine()
		qualities = append(qualities, sum)
	}

	min, max := getMinMax(qualities)
	fmt.Println("Result:", max - min) // result = the diference
}

func getMinMax(arr []int) (min int, max int) {
	min, max = arr[0], 0
	for _, n := range arr {
		if n < 0 {fmt.Printf("WARNING: fishbones contain negative numbers\n")}
		switch {
		case n < min:
			min = n
		case n > max:
			max = n
		}
	}
	return min, max
}

func getStructuredData(path string) map[int][]int {
	data, err := os.ReadFile(path)
	if err != nil {fmt.Printf("ERROR: couldn't read from file\n\toriginal err: %v\n", err)}
	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	results := map[int][]int{}
	var key int
	for _, line := range lines {
		sections := strings.Split(line, ":")
		k, err := strconv.Atoi(sections[0])
		if err != nil {fmt.Printf("ERROR: couldn't convert string to int\n\toriginal err: %v\n", err)}
		key = k
		
		split := strings.Split(sections[1], ",")
		intArr := []int{}
		for _, num := range split {
			n, err := strconv.Atoi(num)
			if err != nil {fmt.Printf("ERROR: couldn't convert string to int\n\toriginal err: %v\n", err)}
			intArr = append(intArr, n)
		}
		_, ok := results[key]
		if ok {fmt.Printf("ERROR: overwriting key %v in results array. \n\tfunction: getStructuredData\n", key)}
		results[key] = intArr
	}
	return results
}
