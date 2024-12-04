package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(file)

	count := 0

	for {
		line, err := r.ReadString('\n')

		line = strings.Trim(line, " \n")

		numbers := make([]int64, 0)
		strsplit := strings.Split(line, " ")
		for _, i := range strsplit {
			num, _ := strconv.ParseInt(i, 10, 64)
			numbers = append(numbers, num)
		}

		increasing := true
		if numbers[1] < numbers[0] {
			increasing = false
		}

		valid := true

		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] < numbers[i+1] && increasing == false {
				valid = false
				break
			}
			if numbers[i] > numbers[i+1] && increasing == true {
				valid = false
				break
			}
			if math.Abs(float64(numbers[i]-numbers[i+1])) < 1 || math.Abs(float64(numbers[i]-numbers[i+1])) > 3 {
				valid = false
				break
			}
		}
		if valid {
			count++
		}
		if err != nil {
			break
		}
	}

	fmt.Println(count)

}
