package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.ReadFile("input.txt")

	split := strings.Split(strings.TrimSpace(string(file)), "\n\n")

	calories := make([]int, len(split))

	for i, j := range split {
		for _, j := range strings.Fields(j) {
			c, _ := strconv.Atoi(j)
			calories[i] += c
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	// Part 1
	fmt.Println("The Elf carrying the most Calories:", calories[0])

	// Part 2
	fmt.Println("Sum of the top three Elves carrying the most Calories:", calories[0]+calories[1]+calories[2])
}
