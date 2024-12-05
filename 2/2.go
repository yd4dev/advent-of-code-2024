package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	util "github.com/yd4dev/advent-of-code-2024"
)

func main() {

	reader, file := util.OpenFile("input.txt")

	defer file.Close()

	count := checkAllLines(reader)

	fmt.Printf("Answer for Part 1: %v", count)

}

func readAndParseLine(reader *bufio.Reader) ([]int64, error) {
	line, eof := reader.ReadString('\n')

	line = strings.Trim(line, "\r\n")

	numbers := make([]int64, 0)
	strsplit := strings.Split(line, " ")
	for _, i := range strsplit {
		num, _ := strconv.ParseInt(i, 10, 64)
		numbers = append(numbers, num)
	}

	return numbers, eof
}

func checkReport(report []int64) bool {
	increasing := true
	if report[1] < report[0] {
		increasing = false
	}

	valid := true

	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] && !increasing {
			valid = false
			break
		}
		if report[i] > report[i+1] && increasing {
			valid = false
			break
		}
		if math.Abs(float64(report[i]-report[i+1])) < 1 || math.Abs(float64(report[i]-report[i+1])) > 3 {
			valid = false
			break
		}
	}
	return valid
}

func checkAllLines(r *bufio.Reader) int {
	count := 0

	for {

		numbers, eof := readAndParseLine(r)

		valid := checkReport(numbers)

		if valid {
			count++
		}

		if eof != nil {
			break
		}
	}

	return count
}
