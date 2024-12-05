package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	util "github.com/yd4dev/advent-of-code-2024"
)

func main() {
	r, file := util.OpenFile("input.txt")
	defer file.Close()

	a, b := parseLists(r)

	slices.Sort(a)
	slices.Sort(b)

	sum := calculateDifference(a, b)

	fmt.Printf("Answer for Part 1: %v\n", sum)

	similarity := calculateSimilarity(a, b)

	fmt.Printf("Answer for Part 1: %v\n", similarity)
}

func parseLists(reader *bufio.Reader) ([]int64, []int64) {
	a := make([]int64, 0)
	b := make([]int64, 0)

	for {
		number, err := reader.ReadString(' ')
		if err != nil {
			break
		}

		num, err := strconv.ParseInt(strings.TrimSpace(number), 10, 64)
		if err != nil {
			fmt.Println(err)
			break
		}

		a = append(a, num)

		number, _ = reader.ReadString('\n')

		num, err = strconv.ParseInt(strings.TrimSpace(number), 10, 64)
		if err != nil {
			fmt.Println(err)
			break
		}

		b = append(b, num)

	}
	return a, b
}

func calculateDifference(a []int64, b []int64) int64 {
	var sum int64 = 0

	for i := range a {
		abs := a[i] - b[i]
		if abs < 0 {
			abs = -abs
		}
		sum += abs
	}
	return sum
}

func calculateSimilarity(a []int64, b []int64) int64 {
	var similarity int64 = 0
	for i := range len(a) {
		for j := range len(b) {
			if a[i] == b[j] {
				similarity += a[i]
			}
		}
	}
	return similarity
}
