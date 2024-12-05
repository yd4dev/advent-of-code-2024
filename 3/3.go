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

var mulEnabled = true

func main() {
	reader, file := util.OpenFile("input.txt")
	defer file.Close()

	fmt.Printf("Answer for Part 1: %v\n", loop(reader, false))

	reader, file = util.OpenFile("input.txt")
	defer file.Close()

	mulEnabled = true

	fmt.Printf("Answer for Part 2: %v\n", loop(reader, true))
}

func loop(reader *bufio.Reader, useMulEnabled bool) int64 {
	var sum int64 = 0
	for {
		bytes, _ := reader.Peek(1)
		if bytes[0] == 'd' {
			mulEnabled = parseDo(reader)
		} else {
			if mulEnabled || !useMulEnabled {
				sum += parseMul(reader)
			} else {
				parseMul(reader)
			}
			_, err := reader.Peek(1)
			if err != nil {
				break
			}
		}
	}
	return sum
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

func parseDo(reader *bufio.Reader) bool {
	if !accept(reader, 'd') {
		return mulEnabled
	}

	if !accept(reader, 'o') {
		return mulEnabled
	}

	rune, _, _ := reader.ReadRune()

	switch rune {
	case '(':
		if !accept(reader, ')') {
			reader.UnreadRune()
			return mulEnabled
		}
		return true
	case 'n':
		if !accept(reader, '\'') {
			reader.UnreadRune()
			return mulEnabled
		}
		if !accept(reader, 't') {
			reader.UnreadRune()
			return mulEnabled
		}
		if !accept(reader, '(') {
			reader.UnreadRune()
			return mulEnabled
		}
		if !accept(reader, ')') {
			reader.UnreadRune()
			return mulEnabled
		}
		return false
	default:
		return mulEnabled
	}
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
