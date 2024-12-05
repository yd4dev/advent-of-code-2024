package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"unicode"

	util "github.com/yd4dev/advent-of-code-2024"
)

func main() {
	reader, file := util.OpenFile("input.txt")
	defer file.Close()
	var sum int64 = 0
	for {
		sum += parseMul(reader)
		_, err := reader.Peek(1)
		if err != nil {
			break
		}
	}
	fmt.Printf("Answer for Part 1: %v\n", sum)
}

func parseMul(reader *bufio.Reader) int64 {
	if !accept(reader, 'm') {
		return 0
	}

	if !accept(reader, 'u') {
		return 0
	}

	if !accept(reader, 'l') {
		return 0
	}

	if !accept(reader, '(') {
		return 0
	}

	mul1 := parseInt(reader)

	if !accept(reader, ',') {
		return 0
	}

	mul2 := parseInt(reader)

	if !accept(reader, ')') {
		return 0
	}
	return mul1 * mul2

}

func accept(reader *bufio.Reader, char rune) bool {
	rune, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	return rune == char
}

func parseInt(reader *bufio.Reader) int64 {
	num := make([]int64, 0)
	var sum int64 = 0

	rune, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	for unicode.IsDigit(rune) {
		temp, _ := strconv.ParseInt(string(rune), 10, 64)
		num = append(num, temp)
		rune, _, err = reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
	}
	reader.UnreadRune()

	for i := 0; i < len(num); i++ {
		sum += int64(math.Pow(10, float64(len(num)-i-1))) * num[i]
	}

	return sum

}
