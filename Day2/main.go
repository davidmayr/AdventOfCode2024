package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	correct := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")
		numbers := make([]int, 0, len(split))

		increasing := false
		defined := false
		prev := 0
		failed := false

		for _, s := range split {
			number, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Error converting string %s to int\n", s)
				panic(err)
			}
			numbers = append(numbers, number)
		}

		for i, number := range numbers {
			if i > 0 && !defined {
				if number > prev {
					increasing = true
				}
				defined = true
			}

			var diff int

			if increasing {
				diff = number - prev
			} else {
				diff = prev - number
			}

			println(diff)
			if i > 0 {
				if increasing && number < prev {
					println("Failed 2")
					failed = true
					break
				} else if !increasing && number > prev {
					println("Failed 1")
					failed = true
					break
				} else if diff <= 0 || diff > 3 {
					println("Failed 3")

					failed = true
					break
				}
			}

			prev = number
		}
		println("___")
		if !failed {
			correct = correct + 1
		}
	}
	println(correct)
}
