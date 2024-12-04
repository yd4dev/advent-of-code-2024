package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Print(err)
		return
	}

	r := bufio.NewReader(file)

	a := make([]int64, 0)
	b := make([]int64, 0)

	for {
		number, err := r.ReadString(' ')
		if err != nil {
			break
		}

		num, err := strconv.ParseInt(strings.TrimSpace(number), 10, 64)
		if err != nil {
			fmt.Println(err)
			break
		}

		a = append(a, num)

		number, _ = r.ReadString('\n')

		num, err = strconv.ParseInt(strings.TrimSpace(number), 10, 64)
		if err != nil {
			fmt.Println(err)
			break
		}

		b = append(b, num)

	}

	slices.Sort(a)
	slices.Sort(b)

	var sum float64 = 0

	for i := range a {
		sum += math.Abs(float64(a[i] - b[i]))
	}

	fmt.Println(int(sum))

	var similarity int64 = 0
	for i := range len(a) {
		for j := range len(b) {
			if a[i] == b[j] {
				similarity += a[i]
			}
		}
	}

	fmt.Println(similarity)

	defer file.Close()
}
