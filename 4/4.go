package main

import (
	"bufio"
	"fmt"
	"os"

	util "github.com/yd4dev/advent-of-code-2024"
)

var reader *bufio.Reader
var file *os.File

func main() {
	reader, file = util.OpenFile("input.txt")
	defer file.Close()
	list := parseFile()
	sum := checkHorizontal(list) + checkVertical(list) + checkDiagonal(list)
	fmt.Printf("Answer for Part 1: %v\n", sum)
}

func parseLine(line string) []rune {
	return []rune(line)
}

func parseFile() [][]rune {
	var list [][]rune
	for {
		line, err := reader.ReadString('\n')
		list = append(list, parseLine(line))
		if err != nil {
			break
		}
	}
	return list
}

func checkHorizontal(list [][]rune) int {
	count := 0
	for _, i := range list {
		for j := range i {
			if j > len(i)-5 {
				break
			}
			if (i[j] == 'X' && i[j+1] == 'M' && i[j+2] == 'A' && i[j+3] == 'S') || (i[j] == 'S' && i[j+1] == 'A' && i[j+2] == 'M' && i[j+3] == 'X') {
				count++
			}
		}
	}
	return count
}

func checkVertical(list [][]rune) int {
	count := 0
	for i := 0; i < len(list)-3; i++ {
		for j := range list[i] {
			if (list[i][j] == 'X' && list[i+1][j] == 'M' && list[i+2][j] == 'A' && list[i+3][j] == 'S') || (list[i][j] == 'S' && list[i+1][j] == 'A' && list[i+2][j] == 'M' && list[i+3][j] == 'X') {
				count++
			}
		}
	}
	return count
}

func checkDiagonal(list [][]rune) int {
	count := 0
	for i := range list {
		for j := 0; j < len(list[i]); j++ {
			if i < len(list)-3 {
				if j < len(list[i])-3 {
					if (list[i][j] == 'X' && list[i+1][j+1] == 'M' && list[i+2][j+2] == 'A' && list[i+3][j+3] == 'S') || (list[i][j] == 'S' && list[i+1][j+1] == 'A' && list[i+2][j+2] == 'M' && list[i+3][j+3] == 'X') {
						count++
					}
				}
				if j > 2 {

					if (list[i][j] == 'X' && list[i+1][j-1] == 'M' && list[i+2][j-2] == 'A' && list[i+3][j-3] == 'S') || (list[i][j] == 'S' && list[i+1][j-1] == 'A' && list[i+2][j-2] == 'M' && list[i+3][j-3] == 'X') {
						count++
					}
				}
			}
		}
	}
	return count
}
